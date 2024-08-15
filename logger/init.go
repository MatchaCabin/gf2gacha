package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var Logger *zap.SugaredLogger

func init() {
	logFile, err := os.Create("gf2gacha.log")
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(logFile, os.Stdout)
	mwSyncer := zapcore.AddSync(mw)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, mwSyncer, zapcore.InfoLevel)
	Logger = zap.New(core).Sugar()
}
