package logic

import (
	"gf2gacha/model"
	"github.com/pkg/errors"
)

func SaveLocalRecord(uid string, recordList []model.LocalRecord) error {
	if uid == "" {
		return errors.New("uid为空")
	}

	session := model.Engine.NewSession()
	defer session.Close()

	err := session.Table(uid).Sync(new(model.LocalRecord))
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = session.Table(uid).Insert(&recordList)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
