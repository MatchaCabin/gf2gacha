package logic

import (
	"gf2gacha/model"
	"github.com/pkg/errors"
)

func RemoveLocalRecord(uid string, poolType int64) (err error) {
	session := model.Engine.NewSession()
	defer session.Close()

	_, err = session.Table(uid).Delete(&model.LocalRecord{PoolType: poolType})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
