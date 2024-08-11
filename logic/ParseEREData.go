package logic

import (
	"encoding/json"
	"gf2gacha/model"
	"github.com/pkg/errors"
	"strconv"
)

func ParseEreData(ereJsonData []byte) ([]model.LocalRecord, error) {
	var ereData model.Ere
	err := json.Unmarshal(ereJsonData, &ereData)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var recordList []model.LocalRecord
	for _, result := range ereData.Result {
		poolType, err := strconv.ParseInt(result[0].(string), 10, 64)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, ereUnit := range result[1].([]interface{}) {
			ereUnitArray := ereUnit.([]interface{})
			recordList = append(recordList, model.LocalRecord{
				PoolType:       poolType,
				PoolId:         int64(ereUnitArray[4].(float64)),
				ItemId:         int64(ereUnitArray[0].(float64)),
				GachaTimestamp: int64(ereUnitArray[5].(float64)),
			})
		}
	}

	return recordList, nil
}
