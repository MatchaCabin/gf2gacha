package logic

import (
	"gf2gacha/model"
	"github.com/pkg/errors"
)

func GetLocalRecord(uid string, poolType int64) (recordList []model.LocalRecord, err error) {
	session := model.Engine.NewSession()
	defer session.Close()

	err = session.Table(uid).Sync(new(model.LocalRecord))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = session.Table(uid).Find(&recordList, &model.LocalRecord{PoolType: poolType})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return recordList, nil
}
