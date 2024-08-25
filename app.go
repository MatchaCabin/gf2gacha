package main

import (
	"context"
	"gf2gacha/config"
	"gf2gacha/logger"
	"gf2gacha/logic"
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/pkg/errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetLogInfo() model.LogInfo {
	info, err := util.GetLogInfo()
	if err != nil {
		logger.Logger.Error(err)
		return model.LogInfo{}
	}
	return info
}

func (a *App) GetUserList() []string {
	userList, err := logic.GetUserList()
	if err != nil {
		logger.Logger.Error(err)
		return nil
	}
	return userList
}

func (a *App) GetPoolInfo(uid string, poolType int64) model.Pool {
	pool, err := logic.GetPoolInfo(uid, poolType)
	if err != nil {
		logger.Logger.Error(err)
		return model.Pool{}
	}
	return pool
}

func (a *App) UpdatePoolInfo(isFull bool) ([]string, error) {
	messageList, err := logic.UpdatePoolInfo(isFull)
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return messageList, nil
}

func (a *App) MergeEreRecord(uid, typ string) (message string, err error) {
	if uid == "" {
		return "", errors.New("UID为空,请至少更新一次数据再进行合并")
	}

	var fileOption runtime.OpenDialogOptions
	switch strings.ToLower(typ) {
	case "json":
		fileOption.Title = "请选择Exilium Recruit Export的Json文件"
		fileOption.Filters = []runtime.FileFilter{{DisplayName: "EreJsonData", Pattern: "*.json"}}
	case "excel":
		fileOption.Title = "请选择Exilium Recruit Export的Excel文件"
		fileOption.Filters = []runtime.FileFilter{{DisplayName: "EreExcelData", Pattern: "*.xlsx"}}
	default:
		return "", errors.Errorf("unknown type: %s", typ)
	}
	erePath, err := runtime.OpenFileDialog(a.ctx, fileOption)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	if erePath == "" {
		return "", errors.New("用户取消导入")
	}

	err = logic.MergeEreRecord(uid, erePath, typ)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	message = "合并成功"
	return
}

func (a *App) ImportRawJson(uid string, isReverse bool) (message string, err error) {
	if uid == "" {
		return "", errors.New("UID为空,请至少更新一次数据再进行导出")
	}

	fileOption := runtime.OpenDialogOptions{
		Title:   "请选择RawJson文件",
		Filters: []runtime.FileFilter{{DisplayName: "RawJsonData", Pattern: "*.json"}},
	}
	importFilePath, err := runtime.OpenFileDialog(a.ctx, fileOption)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	if importFilePath == "" {
		return "", errors.New("用户取消导入")
	}

	err = logic.ImportRawJson(uid, importFilePath, isReverse)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	message = "合并成功"
	return
}

func (a *App) ExportRawJson(uid string) (message string, err error) {
	if uid == "" {
		return "", errors.New("UID为空,请至少更新一次数据再进行导出")
	}

	fileOption := runtime.OpenDialogOptions{
		Title: "选择RawJson保存目录",
	}
	saveDir, err := runtime.OpenDirectoryDialog(a.ctx, fileOption)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	if saveDir == "" {
		return "", errors.New("用户取消导出")
	}

	err = logic.ExportRawJson(uid, saveDir)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	message = "导出成功"
	return
}

func (a *App) ImportMccExcel(uid string) (message string, err error) {
	if uid == "" {
		return "", errors.New("UID为空,请至少更新一次数据再进行导出")
	}

	fileOption := runtime.OpenDialogOptions{
		Title:   "请选择MccExcel文件",
		Filters: []runtime.FileFilter{{DisplayName: "MccExcel", Pattern: "*.xlsx"}},
	}
	importFilePath, err := runtime.OpenFileDialog(a.ctx, fileOption)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	if importFilePath == "" {
		return "", errors.New("用户取消导入")
	}

	err = logic.ImportMccExcel(uid, importFilePath)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	message = "MccExcel合并成功"
	return
}

func (a *App) ExportMccExcel(uid string) (message string, err error) {
	if uid == "" {
		return "", errors.New("UID为空,请至少更新一次数据再进行导出")
	}

	fileOption := runtime.OpenDialogOptions{
		Title: "选择MccExcel保存目录",
	}
	saveDir, err := runtime.OpenDirectoryDialog(a.ctx, fileOption)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	if saveDir == "" {
		return "", errors.New("用户取消导出")
	}

	err = logic.ExportMccExcel(uid, saveDir)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	message = "MccExcel导出成功"
	return
}

func (a *App) HandleCommunityTasks() (messageList []string, err error) {
	messageList, err = logic.HandleCommunityTasks()
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	return
}

func (a *App) CheckUpdate() (string, error) {
	version, err := logic.CheckUpdate()
	if err != nil {
		logger.Logger.Error(err)
		return "", err
	}
	return version, nil
}

func (a *App) ApplyUpdate() (string, error) {
	err := logic.ApplyUpdate()
	if err != nil {
		logger.Logger.Error(err)
		return "", err
	}
	return "", nil
}

func (a *App) GetCommunityExchangeList() ([]model.CommunityExchangeList, error) {
	list, err := logic.GetCommunityExchangeList()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return list, nil
}

func (a *App) GetSettingExchangeList() ([]int64, error) {
	if !config.IsSetExchangeList() {
		exchangeList, err := logic.GetCommunityExchangeList()
		if err != nil {
			logger.Logger.Error(err)
			return nil, err
		}
		var idList []int64
		for _, item := range exchangeList {
			idList = append(idList, item.Id)
		}
		err = config.SetExchangeList(idList)
		if err != nil {
			logger.Logger.Error(err)
			return nil, err
		}
	}

	return config.GetExchangeList(), nil
}

func (a *App) SaveSettingExchangeList(exchangeList []int64) error {
	return config.SetExchangeList(exchangeList)
}
