package util

import (
	"gf2gacha/model"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"regexp"
)

func GetGameInfo() (gameInfo model.Info, err error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return model.Info{}, errors.WithStack(err)
	}

	logPath := filepath.Join(userHome, "/AppData/LocalLow/SunBorn/少女前线2：追放/Player.log")
	logData, err := os.ReadFile(logPath)
	if err != nil {
		return model.Info{}, errors.WithStack(err)
	}

	regexpGamePath, err := regexp.Compile(`\[Subsystems] Discovering subsystems at path (.+)/UnitySubsystems`)
	if err != nil {
		return model.Info{}, errors.WithStack(err)
	}
	resultGamePath := regexpGamePath.FindSubmatch(logData)
	if len(resultGamePath) == 2 {
		gameInfo.TablePath = filepath.Join(string(resultGamePath[1]), "LocalCache/Data/Table")
	} else {
		return model.Info{}, errors.New("未在日志中找到游戏路径")
	}

	regexpUserInfo, err := regexp.Compile(`"access_token":"(.+?)".+"uid":(\d+)`)
	if err != nil {
		return model.Info{}, errors.WithStack(err)
	}
	resultUserInfo := regexpUserInfo.FindSubmatch(logData)
	if len(resultUserInfo) == 3 {
		gameInfo.AccessToken = string(resultUserInfo[1])
		gameInfo.Uid = string(resultUserInfo[2])
	} else {
		return model.Info{}, errors.New("未在日志中找到AccessToken或Uid")
	}

	regexpGachaUrl, err := regexp.Compile(`"gacha_record_url":"(.*?)"`)
	if err != nil {
		return model.Info{}, errors.WithStack(err)
	}
	resultGachaUrl := regexpGachaUrl.FindSubmatch(logData)
	if len(resultGachaUrl) == 2 {
		gameInfo.GachaUrl = string(resultGachaUrl[1])
	} else {
		return model.Info{}, errors.New("未在日志中找到抽卡链接")
	}

	return gameInfo, nil
}
