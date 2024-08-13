package request

import (
	"github.com/pkg/errors"
)

func CommunityExchange(webToken string, exchangeId int64) error {
	apiUrl := `https://gf2-bbs-api.sunborngame.com/community/item/exchange`

	params := map[string]interface{}{
		"exchange_id": exchangeId,
	}

	_, err := CommunityPost(apiUrl, params, webToken)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
