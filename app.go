package main

import (
	"context"
	"fmt"
	"gf2gacha/config"
	"gf2gacha/logger"
	"gf2gacha/logic"
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/elazarl/goproxy"
	"github.com/pkg/errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

// App struct
type App struct {
	ctx          context.Context
	captureMutex sync.Mutex
	tcpPort      int
	tcpListener  *net.Listener
	MitmServer   *http.Server
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

func (a *App) GetCurrentVersion() string {
	return util.GetVersion()
}

func (a *App) GetLatestVersion() (string, error) {
	version, err := logic.GetLatestVersion()
	if err != nil {
		logger.Logger.Error(err)
		return "", err
	}
	return version, nil
}

func (a *App) UpdateTo(version string) (string, error) {
	err := logic.UpdateTo(version)
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

func (a *App) GetSettingFont() (string, error) {
	return config.GetFont(), nil
}

func (a *App) SaveSettingFont(newFont string) error {
	return config.SetFont(newFont)
}

func (a *App) GetSettingLayout() (int64, error) {
	return config.GetLayout(), nil
}

func (a *App) SaveSettingLayout(layoutType int64) error {
	return config.SetLayout(layoutType)
}

func (a *App) GetSettingCapturePort() int {
	return config.GetCapturePort()
}

func (a *App) SaveSettingCapturePort(port int) error {
	return config.SetCapturePort(port)
}

func (a *App) CaptureStart() error {
	_, err := os.Stat("ca.crt")
	if err != nil {
		logger.Logger.Errorf("CA证书不存在:%v", err)
		err = util.GenCA()
		if err != nil {
			logger.Logger.Errorf("生成CA证书出错:%v", err)
			return err
		}
	}

	if !util.IsTrustedCA() {
		err = util.InstallCA()
		if err != nil {
			logger.Logger.Errorf("安装CA证书出错:%v", err)
			return err
		}
	}

	if a.tcpListener == nil {
		configPort := config.GetCapturePort()
		if configPort > 0 && configPort <= 65535 {
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configPort))
			if err == nil {
				a.captureMutex.Lock()
				a.tcpPort = configPort
				a.tcpListener = &listener
				a.captureMutex.Unlock()
				logger.Logger.Infof("配置的端口%d可用\n", configPort)
			} else {
				logger.Logger.Warnf("无法监听配置的端口 %d: %v，将尝试自动选择端口。", configPort, err)
			}
		}

		// 如果通过配置未能成功创建监听器，则遍历端口范围自动选择
		if a.tcpListener == nil {
			for port := 60000; port < 60010; port++ {
				listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
				if err != nil {
					logger.Logger.Error(err)
					continue
				}

				a.captureMutex.Lock()
				a.tcpPort = port
				a.tcpListener = &listener
				a.captureMutex.Unlock()
				logger.Logger.Infof("端口%d可用\n", port)
				break
			}
		}
	}

	cert, err := util.ParseCA()
	if err != nil {
		logger.Logger.Errorf("读取CA证书出错:%v", err)
		return err
	}

	customCaMitm := &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(cert)}
	var customAlwaysMitm goproxy.FuncHttpsHandler = func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		return customCaMitm, host
	}

	proxy := goproxy.NewProxyHttpServer()
	cond := goproxy.DstHostIs("gf2-gacha-record.sunborngame.com")

	proxy.OnRequest(cond).HandleConnect(customAlwaysMitm)
	proxy.OnRequest(cond).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if req.URL.Path == "/list" {
			gachaUrl := req.URL.String()
			accessToken := req.Header.Get("Authorization")

			userInfo, err := logic.GetUserInfoFromBBS(accessToken)
			if err != nil {
				logger.Logger.Errorf("未从社区获取到用户信息:%v", err)
				return req, nil
			}

			err = logic.AppendLog(accessToken, userInfo.User.GameUid, gachaUrl)
			if err != nil {
				logger.Logger.Errorf("追加游戏日志失败:%v", err)
				return req, nil
			}

			logger.Logger.Infof("uid: %d\n", userInfo.User.GameUid)
			logger.Logger.Infof("gachaUrl: %s\n", gachaUrl)
			logger.Logger.Infof("accessToken: %s\n", accessToken)

			runtime.EventsEmit(a.ctx, "captureSuccess")
		}

		return req, nil
	})

	//启动监听
	if a.MitmServer == nil {
		a.captureMutex.Lock()
		a.MitmServer = &http.Server{Handler: proxy}

		go func(mitmServer *http.Server, tcpListener *net.Listener) {
			logger.Logger.Info("捕获协程启动")
			err := mitmServer.Serve(*tcpListener)
			if err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					logger.Logger.Info("捕获正常停止")
					return
				} else {
					logger.Logger.Errorf("捕获异常停止: %v", err)
				}
				a.captureMutex.Lock()
				a.tcpListener = nil
				a.MitmServer = nil
				a.captureMutex.Unlock()
			}
		}(a.MitmServer, a.tcpListener)

		err = util.EnableSysProxy(a.tcpPort)
		if err != nil {
			logger.Logger.Errorf("设置代理信息出错:%v", err)
			return err
		}

		a.captureMutex.Unlock()
	}

	return nil
}

func (a *App) CaptureClose() error {
	a.captureMutex.Lock()
	defer a.captureMutex.Unlock()

	if a.MitmServer != nil {
		err := a.MitmServer.Shutdown(context.Background())
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Error(err)
			return err
		}
		a.tcpListener = nil
		a.MitmServer = nil
	}

	err := util.DisableSysProxy()
	if err != nil {
		logger.Logger.Errorf("还原代理信息出错:%v", err)
		return err
	}

	return nil
}
