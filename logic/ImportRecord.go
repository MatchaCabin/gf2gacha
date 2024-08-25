package logic

import (
	"encoding/json"
	"gf2gacha/logger"
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"os"
	"slices"
	"strconv"
	"time"
)

func ImportRawJson(uid, rawJsonPath string, isReverse bool) error {
	rawJsonBytes, err := os.ReadFile(rawJsonPath)
	if err != nil {
		return errors.WithStack(err)
	}

	rawMap := make(map[int64][]model.RemoteRecord)
	err = json.Unmarshal(rawJsonBytes, &rawMap)
	if err != nil {
		return errors.WithStack(err)
	}

	//解析成功后先做备份
	err = util.BackupDB()
	if err != nil {
		return errors.WithStack(err)
	}

	for poolType, rawRecordList := range rawMap {
		var recordList []model.LocalRecord
		if isReverse {
			slices.Reverse(rawRecordList)
		}
		for _, record := range rawRecordList {
			recordList = append(recordList, model.LocalRecord{
				PoolType:       poolType,
				PoolId:         record.PoolId,
				ItemId:         record.ItemId,
				GachaTimestamp: record.GachaTimestamp,
			})
		}

		localRecordList, err := GetLocalRecord(uid, poolType, 0)
		if err != nil {
			return errors.WithStack(err)
		}

		newRecordList := MergeRecord(localRecordList, recordList)
		if len(newRecordList) == 0 {
			return nil
		}

		err = RemoveLocalRecord(uid, poolType)
		if err != nil {
			return errors.WithStack(err)
		}

		err = SaveLocalRecord(uid, newRecordList)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func ImportMccExcel(uid, excelPath string) error {
	excel, err := excelize.OpenFile(excelPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer excel.Close()

	poolMap := make(map[int64][]model.LocalRecord)
	for _, sheetName := range excel.GetSheetList() {
		var poolType int64
		switch sheetName {
		case "常规采购":
			poolType = 1
		case "定向采购":
			poolType = 3
		case "军备提升":
			poolType = 4
		case "初始采购":
			poolType = 5
		case "神秘箱":
			poolType = 8
		default:
			logger.Logger.Warnf("未知Sheet名称:%s", sheetName)
			continue
		}
		rows, err := excel.GetRows(sheetName)
		if err != nil {
			errors.WithStack(err)
		}

		fieldMapping := make(map[string]int)
		for rowIndex, row := range rows {
			//先获取标题栏索引值，以防止有乱序的情况
			if rowIndex == 0 {
				for colIndex, cell := range row {
					switch cell {
					case "卡池Id":
						fieldMapping["PoolId"] = colIndex
					case "抽卡时间":
						fieldMapping["GachaTime"] = colIndex
					case "道具Id":
						fieldMapping["ItemId"] = colIndex
					}
				}
				continue
			}

			var poolId, itemId, gachaTimestamp int64
			poolIdIndex := fieldMapping["PoolId"]
			if poolIdIndex >= len(row) {
				return errors.Errorf("%s第%d行:卡池Id(Index:%d)越界或不存在", sheetName, rowIndex, poolIdIndex)
			}
			poolId, err = strconv.ParseInt(row[poolIdIndex], 10, 64)
			if err != nil {
				return errors.WithStack(err)
			}

			gachaTimeIndex := fieldMapping["GachaTime"]
			if gachaTimeIndex >= len(row) {
				return errors.Errorf("%s第%d行:抽卡时间(Index:%d)越界或不存在", sheetName, rowIndex, gachaTimeIndex)
			}

			t, err := time.ParseInLocation("2006年01月02日 15:04:05", row[gachaTimeIndex], time.Local)
			if err != nil {
				return errors.WithStack(err)
			}
			gachaTimestamp = t.Unix()

			itemIdIndex := fieldMapping["ItemId"]
			if itemIdIndex >= len(row) {
				return errors.Errorf("%s第%d行:道具Id(Index:%d)越界或不存在", sheetName, rowIndex, itemIdIndex)
			}
			itemId, err = strconv.ParseInt(row[itemIdIndex], 10, 64)
			if err != nil {
				return errors.WithStack(err)
			}

			poolMap[poolType] = append(poolMap[poolType], model.LocalRecord{
				PoolType:       poolType,
				PoolId:         poolId,
				ItemId:         itemId,
				GachaTimestamp: gachaTimestamp,
			})
		}
	}

	if len(poolMap) == 0 {
		return errors.New("Excel内无可解析数据")
	}

	//解析成功后先做备份
	err = util.BackupDB()
	if err != nil {
		return errors.WithStack(err)
	}

	for poolType, recordList := range poolMap {
		localRecordList, err := GetLocalRecord(uid, poolType, 0)
		if err != nil {
			return errors.WithStack(err)
		}

		newRecordList := MergeRecord(localRecordList, recordList)
		if len(newRecordList) == 0 {
			return nil
		}

		err = RemoveLocalRecord(uid, poolType)
		if err != nil {
			return errors.WithStack(err)
		}

		err = SaveLocalRecord(uid, newRecordList)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
