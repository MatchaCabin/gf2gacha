package logic

import (
	"encoding/json"
	"fmt"
	"gf2gacha/model"
	"gf2gacha/preload"
	"github.com/pkg/errors"
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
