package logic

import (
	"gf2gacha/model"
	"github.com/pkg/errors"
)

func SaveLocalRecord(uid string, recordList []model.LocalRecord) (int64, error) {
	if uid == "" {
		return 0, errors.New("uid为空")
	}

	session := model.Engine.NewSession()
	defer session.Close()

	err := session.Table(uid).Sync(new(model.LocalRecord))
	if err != nil {
		return 0, errors.WithStack(err)
	}

	n, err := session.Table(uid).Insert(&recordList)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return n, nil
}
