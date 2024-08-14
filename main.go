package main

import (
	"embed"
	"fmt"
	"gf2gacha/logger"
	_ "gf2gacha/logger"
	"gf2gacha/util"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     fmt.Sprintf("少女前线2:追放 抽卡导出分析工具%v Powered by MccWiki抹茶", util.GetVersion()),
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
		logger.Logger.Error(err)
	}
}
