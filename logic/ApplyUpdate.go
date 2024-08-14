package logic

import (
	"context"
	"gf2gacha/util"
	"github.com/google/go-github/v63/github"
	"github.com/inconshreveable/go-update"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/exec"
)

func ApplyUpdate() error {
	client := github.NewClient(http.DefaultClient)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "MatchaCabin", "gf2gacha")
	if err != nil {
		return errors.WithStack(err)
	}

	if util.GetVersion() == release.GetTagName() {
		return nil
	}

	var link string
	for _, asset := range release.Assets {
		if asset.GetName() == "gf2gacha.exe" {
			link = asset.GetBrowserDownloadURL()
			break
		}
	}

	if link != "" {
		proxyLink := `https://mirror.ghproxy.com/` + link
		logrus.Infof("代理链接:%s", proxyLink)
		//优先尝试用国内代理下载
		respProxy, err := http.Get(proxyLink)
		if os.IsTimeout(err) {
			logrus.Infof("源链接:%s", link)
			resp, err := http.Get(link)
			if err != nil {
				return errors.WithStack(err)
			}
			defer resp.Body.Close()
			logrus.Infof("下载成功")
			err = update.Apply(respProxy.Body, update.Options{})
			if err != nil {
				return errors.WithStack(err)
			}
			logrus.Infof("更新成功")

			err = restart()
			if err != nil {
				return errors.WithStack(err)
			}
			logrus.Infof("重启成功")
			os.Exit(1)
		} else if err != nil {
			return errors.WithStack(err)
		}
		defer respProxy.Body.Close()
		logrus.Infof("使用代理下载成功")

		err = update.Apply(respProxy.Body, update.Options{})
		if err != nil {
			return errors.WithStack(err)
		}
		logrus.Infof("更新成功")

		err = restart()
		if err != nil {
			return errors.WithStack(err)
		}
		logrus.Infof("重启成功")
		os.Exit(1)
	}

	return nil
}

func restart() error {
	execPath, err := os.Executable()
	if err != nil {
		return err
	}
	cmd := exec.Command(execPath)
	return cmd.Start()
}
