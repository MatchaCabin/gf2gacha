package logic

import (
	"gf2gacha/model"
	"github.com/pkg/errors"
)

func GetUserList() ([]string, error) {
	metas, err := model.Engine.DBMetas()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var uidList []string
	for _, meta := range metas {
		uidList = append(uidList, meta.Name)
	}

	return uidList, nil
}
