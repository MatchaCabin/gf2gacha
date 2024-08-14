package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	logFile, err := os.OpenFile("gf2gacha.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{})
}
