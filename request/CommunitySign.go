package request

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type CommunitySignData struct {
	GetExp       int    `json:"get_exp"`
	GetItemCount int    `json:"get_item_count"`
	GetItemName  string `json:"get_item_name"`
	GetItemUrl   string `json:"get_item_url"`
	GetScore     int    `json:"get_score"`
}

func CommunitySign(webToken string) (CommunitySignData, error) {
	apiUrl := `https://gf2-bbs-api.sunborngame.com/community/task/sign_in`

	dataBytes, err := CommunityPost(apiUrl, nil, webToken)
	if err != nil {
		return CommunitySignData{}, errors.WithStack(err)
	}

	var data CommunitySignData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return CommunitySignData{}, errors.WithStack(err)
	}

	return data, nil
}
