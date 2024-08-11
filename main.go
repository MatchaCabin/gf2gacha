package main

import (
	"embed"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"io"
	"os"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	logFile, err := os.Create("gf2gacha.log")
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "少女前线2:追放 抽卡导出分析工具 Powered by MccWiki抹茶",
		Width:     1280,
		Height:    720,
		MinWidth:  1280,
		MinHeight: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
