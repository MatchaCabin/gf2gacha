package logic

import (
	"context"
	"gf2gacha/logger"
	"gf2gacha/util"
	"github.com/google/go-github/v63/github"
	"github.com/pkg/errors"
	"net/http"
)

func CheckUpdate() (string, error) {
	client := github.NewClient(http.DefaultClient)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "MatchaCabin", "gf2gacha")
	if err != nil {
		return "", errors.WithStack(err)
	}
	logger.Logger.Infof("Latest release: %s", *release.TagName)

	if util.GetVersion() != release.GetTagName() {
		return release.GetTagName(), nil
	}

	return "", nil
}
