package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger *logrus.Logger

func init() {
	logFile, err := os.Create("gf2gacha.log")
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(logFile, os.Stdout)
	Logger = logrus.New()
	Logger.SetOutput(mw)
	Logger.SetFormatter(&logrus.TextFormatter{})
}
