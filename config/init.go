package config

import (
	"fmt"
	"gf2gacha/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			logger.Logger.Warnf("配置文件不存在，将自动创建默认配置文件")
			err = viper.SafeWriteConfig()
			if err != nil {
				logger.Logger.Panic(err)
			}
		} else {
			logger.Logger.Panic(err)
		}
	}
}

func GetGameDataDir() string {
	return viper.GetString("GameDataDir")
}

func SetGameDataDir(gameDataDir string) error {
	viper.Set("GameDataDir", gameDataDir)
	return viper.WriteConfig()
}

func IsSetGameDataDir() bool {
	return viper.IsSet("GameDataDir")
}

func GetWebToken(uid string) string {
	return viper.GetString(fmt.Sprintf("%s.webToken", uid))
}

func SetWebToken(uid, token string) error {
	viper.Set(fmt.Sprintf("%s.webToken", uid), token)
	return viper.WriteConfig()
}

func IsSetExchangeList() bool {
	return viper.IsSet("exchangeList")
}

func GetExchangeList() []int64 {
	intList := viper.GetIntSlice("exchangeList")
	var int64List []int64
	for i := range intList {
		int64List = append(int64List, int64(intList[i]))
	}

	return int64List
}

func SetExchangeList(exchangeList []int64) error {
	viper.Set("exchangeList", exchangeList)
	return viper.WriteConfig()
}

func GetFont() string {
	return viper.GetString("font")
}

func SetFont(newFont string) error {
	viper.Set("font", newFont)
	return viper.WriteConfig()
}

func GetLayout() int64 {
	return viper.GetInt64("layout")
}

func SetLayout(layoutType int64) error {
	viper.Set("layout", layoutType)
	return viper.WriteConfig()
}

func GetCapturePort() int {
	return viper.GetInt("capturePort")
}

func SetCapturePort(port int) error {
	viper.Set("capturePort", port)
	return viper.WriteConfig()
}
