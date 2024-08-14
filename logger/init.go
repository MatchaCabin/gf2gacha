package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	logFile, err := os.Create("gf2gacha.log")
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{})
}
