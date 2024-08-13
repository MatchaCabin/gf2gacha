package util

import (
	"gf2gacha/model"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"regexp"
)

func GetLogInfo() (logInfo model.LogInfo, err error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return model.LogInfo{}, errors.WithStack(err)
	}

	logPath := filepath.Join(userHome, "/AppData/LocalLow/SunBorn/少女前线2：追放/Player.log")
	logData, err := os.ReadFile(logPath)
	if err != nil {
		return model.LogInfo{}, errors.WithStack(err)
	}

	regexpGamePath, err := regexp.Compile(`\[Subsystems] Discovering subsystems at path (.+)/UnitySubsystems`)
	if err != nil {
		return model.LogInfo{}, errors.WithStack(err)
	}
	resultGamePath := regexpGamePath.FindSubmatch(logData)
	if len(resultGamePath) == 2 {
		logInfo.TablePath = filepath.Join(string(resultGamePath[1]), "LocalCache/Data/Table")
	} else {
		return model.LogInfo{}, errors.New("未在日志中找到游戏路径")
	}

	regexpUserInfo, err := regexp.Compile(`"access_token":"(.+?)".+"uid":(\d+)`)
	if err != nil {
		return model.LogInfo{}, errors.WithStack(err)
	}
	resultUserInfo := regexpUserInfo.FindSubmatch(logData)
	if len(resultUserInfo) == 3 {
		logInfo.AccessToken = string(resultUserInfo[1])
		logInfo.Uid = string(resultUserInfo[2])
	} else {
		return model.LogInfo{}, errors.New("未在日志中找到AccessToken或Uid")
	}

	regexpGachaUrl, err := regexp.Compile(`"gacha_record_url":"(.*?)"`)
	if err != nil {
		return model.LogInfo{}, errors.WithStack(err)
	}
	resultGachaUrl := regexpGachaUrl.FindSubmatch(logData)
	if len(resultGachaUrl) == 2 {
		logInfo.GachaUrl = string(resultGachaUrl[1])
	} else {
		return model.LogInfo{}, errors.New("未在日志中找到抽卡链接")
	}

	return logInfo, nil
}
