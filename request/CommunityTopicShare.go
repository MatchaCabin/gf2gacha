package request

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

func CommunityTopicShare(webToken string, topicId int64) error {
	apiUrl := fmt.Sprintf(`https://gf2-bbs-api.sunborngame.com/community/topic/share/%s`, topicId)

	values := url.Values{
		"id": []string{strconv.FormatInt(topicId, 10)},
	}

	_, err := CommunityGet(apiUrl, values, webToken)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
