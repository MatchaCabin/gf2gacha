package logic

import (
	"context"
	"gf2gacha/util"
	"github.com/google/go-github/v63/github"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CheckUpdate() (string, error) {
	client := github.NewClient(http.DefaultClient)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "MatchaCabin", "gf2gacha")
	if err != nil {
		return "", errors.WithStack(err)
	}
	logrus.Infof("Latest release: %s", *release.TagName)

	if util.GetVersion() != release.GetTagName() {
		return release.GetTagName(), nil
	}

	return "", nil
}
