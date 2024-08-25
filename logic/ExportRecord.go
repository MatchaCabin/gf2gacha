package logic

import (
	"encoding/json"
	"fmt"
	"gf2gacha/logger"
	"gf2gacha/model"
	"gf2gacha/preload"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"time"
)

func ExportRawJson(uid, saveDir string) error {
	rawMap := make(map[int64][]model.RemoteRecord)
	for _, poolTypeUnit := range preload.PoolTypeMap {
		localRecordList, err := GetLocalRecord(uid, poolTypeUnit.Id, 0)
		if err != nil {
			return errors.WithStack(err)
		}
		for _, record := range localRecordList {
			rawMap[poolTypeUnit.Id] = append(rawMap[poolTypeUnit.Id], model.RemoteRecord{
				PoolId:         record.PoolId,
				ItemId:         record.ItemId,
				GachaTimestamp: record.GachaTimestamp,
			})
		}
	}
	rawJsonBytes, err := json.MarshalIndent(&rawMap, "", "\t")
	if err != nil {
		return errors.WithStack(err)
	}

	err = os.WriteFile(filepath.Join(saveDir, fmt.Sprintf("gf2gacha-raw-%s_%s.json", uid, time.Now().Format("20060102150405"))), rawJsonBytes, 0755)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ExportMccExcel(uid, saveDir string) error {
	excel := excelize.NewFile()
	defer excel.Close()

	styleTitle, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Family: "Microsoft YaHei",
			Size:   14,
			Bold:   true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	styleRank3, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Family: "Microsoft YaHei",
			Size:   14,
			Color:  "#5B9BD5",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	styleRank4, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Family: "Microsoft YaHei",
			Size:   14,
			Color:  "#CC00CC",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	styleRank5, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Family: "Microsoft YaHei",
			Size:   14,
			Color:  "#FFC000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	for _, poolTypeUnit := range preload.PoolTypeMap {
		localRecordList, err := GetLocalRecord(uid, poolTypeUnit.Id, 0)
		if err != nil {
			return errors.WithStack(err)
		}
		if len(localRecordList) > 0 {
			var sheetName string
			switch poolTypeUnit.Id {
			case 1:
				sheetName = "常规采购"
			case 3:
				sheetName = "定向采购"
			case 4:
				sheetName = "军备提升"
			case 5:
				sheetName = "初始采购"
			case 8:
				sheetName = "神秘箱"
			default:
				logger.Logger.Warnf("未定义的卡池 poolType:%d poolName:%s", poolTypeUnit.Id, poolTypeUnit.Name)
				continue
			}
			_, err = excel.NewSheet(sheetName)
			if err != nil {
				return errors.WithStack(err)
			}

			err = excel.SetSheetRow(sheetName, "A1", &[]string{"卡池Id", "抽卡时间", "道具Id", "道具名称"})
			if err != nil {
				return errors.WithStack(err)
			}

			err = excel.SetRowStyle(sheetName, 1, 1, styleTitle)
			if err != nil {
				return errors.WithStack(err)
			}
			err = excel.SetColWidth(sheetName, "A", "A", 10)
			if err != nil {
				return errors.WithStack(err)
			}
			err = excel.SetColWidth(sheetName, "B", "B", 36)
			if err != nil {
				return errors.WithStack(err)
			}
			err = excel.SetColWidth(sheetName, "C", "C", 12)
			if err != nil {
				return errors.WithStack(err)
			}
			err = excel.SetColWidth(sheetName, "D", "D", 42)
			if err != nil {
				return errors.WithStack(err)
			}

			for i, record := range localRecordList {
				rowIndex := i + 2
				cellName := fmt.Sprintf("A%d", rowIndex)
				poolId := record.PoolId
				gachaTime := time.Unix(record.GachaTimestamp, 0).Format("2006年01月02日 15:04:05")
				itemId := record.ItemId
				itemName := preload.LangMap[preload.ItemMap[itemId].Name.Id]
				err = excel.SetSheetRow(sheetName, cellName, &[]interface{}{poolId, gachaTime, itemId, itemName})
				if err != nil {
					return errors.WithStack(err)
				}

				itemRank := preload.ItemRankMap[poolId][itemId]
				switch itemRank {
				case 3:
					err = excel.SetRowStyle(sheetName, rowIndex, rowIndex, styleRank3)
					if err != nil {
						return errors.WithStack(err)
					}
				case 4:
					err = excel.SetRowStyle(sheetName, rowIndex, rowIndex, styleRank4)
					if err != nil {
						return errors.WithStack(err)
					}
				case 5:
					err = excel.SetRowStyle(sheetName, rowIndex, rowIndex, styleRank5)
					if err != nil {
						return errors.WithStack(err)
					}
				default:
					logger.Logger.Warnf("未知道具等级%d %+v", itemRank, record)
				}
			}
		}
	}

	err = excel.DeleteSheet("Sheet1")
	if err != nil {
		return errors.WithStack(err)
	}

	err = excel.SaveAs(filepath.Join(saveDir, fmt.Sprintf("gf2gacha-excel-%s_%s.xlsx", uid, time.Now().Format("20060102150405"))))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
