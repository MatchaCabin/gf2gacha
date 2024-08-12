package model

import (
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", "./gf2gacha.db")
	if err != nil {
		logrus.Panic(err)
	}
}
