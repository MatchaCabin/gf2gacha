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

	var dl string
	for _, asset := range release.Assets {
		if asset.GetName() == "gf2gacha.exe" {
			dl = asset.GetBrowserDownloadURL()
			break
		}
	}

	if dl != "" {
		logrus.Infof("下载链接:%s", dl)
		resp, err := http.Get(dl)
		if os.IsTimeout(err) {
			//如果超时，尝试用国内代理下载
			respProxy, err := http.Get(`https://mirror.ghproxy.com/` + dl)
			if err != nil {
				return errors.WithStack(err)
			}
			defer respProxy.Body.Close()

		} else if err != nil {
			return errors.WithStack(err)
		}
		defer resp.Body.Close()

		err = update.Apply(resp.Body, update.Options{})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
