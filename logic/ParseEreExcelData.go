package logic

import (
	"bytes"
	"gf2gacha/model"
	"gf2gacha/preload"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func ParseEreExcelData(ereExcelData []byte) ([]model.LocalRecord, error) {
	reader, err := excelize.OpenReader(bytes.NewReader(ereExcelData))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer reader.Close()

	var recordList []model.LocalRecord
	for _, sheetName := range reader.GetSheetList() {
		var poolType int64
		switch sheetName {
		case "常规采购":
			poolType = 1
		case "定向采购":
			poolType = 3
		case "军备提升":
			poolType = 4
		default:
			logrus.Warnf("未知Sheet名称:%s", sheetName)
			continue
		}
		rows, err := reader.GetRows(sheetName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		var fieldMapping map[string]int
		for rowIndex, row := range rows {
			//先获取标题栏索引值，以防止有乱序的情况
			if rowIndex == 0 {
				for colIndex, cell := range row {
					switch cell {
					case "时间":
						fieldMapping["GachaTimestamp"] = colIndex
					case "备注":
						fieldMapping["PoolId"] = colIndex
					case "类别":
						fieldMapping["ItemType"] = colIndex
					case "名称":
						fieldMapping["ItemName"] = colIndex
					}
				}
				continue
			}

			var poolId, itemId, gachaTimestamp int64
			var itemType, itemName string
			poolIdIndex := fieldMapping["PoolId"]
			if poolIdIndex >= len(row) {
				return nil, errors.Errorf("%s第%d行:备注(Index:%d)越界", sheetName, rowIndex, poolIdIndex)
			}
			poolId, err = strconv.ParseInt(row[poolIdIndex], 10, 64)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			gachaTimestampIndex := fieldMapping["GachaTimestamp"]
			if gachaTimestampIndex >= len(row) {
				return nil, errors.Errorf("%s第%d行:时间(Index:%d)越界", sheetName, rowIndex, gachaTimestampIndex)
			}
			gachaTimestamp, err = strconv.ParseInt(row[gachaTimestampIndex], 10, 64)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			itemTypeIndex := fieldMapping["ItemType"]
			if itemTypeIndex >= len(row) {
				return nil, errors.Errorf("%s第%d行:类别(Index:%d)越界", sheetName, rowIndex, itemTypeIndex)
			}
			itemType = row[itemTypeIndex]

			itemNameIndex := fieldMapping["ItemName"]
			if itemNameIndex >= len(row) {
				return nil, errors.Errorf("%s第%d行:名称(Index:%d)越界", sheetName, rowIndex, itemNameIndex)
			}
			itemName = row[itemNameIndex]

			switch itemType {
			case "角色":
				itemId = preload.DollNameMapping[itemName]
			case "武器":
				itemId = preload.WeaponNameMapping[itemName]
			default:
				return nil, errors.Errorf("%s第%d行:未知类别(Index:%d):%s", sheetName, rowIndex, itemTypeIndex, itemType)
			}
			recordList = append(recordList, model.LocalRecord{
				PoolType:       poolType,
				PoolId:         poolId,
				ItemId:         itemId,
				GachaTimestamp: gachaTimestamp,
			})
		}
	}

	return recordList, nil
}
