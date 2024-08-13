package request

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type CommunityExchangeListData struct {
	List []struct {
		ExchangeId       int64  `json:"exchange_id"`
		ItemName         string `json:"item_name"`
		ItemCount        int64  `json:"item_count"`
		ItemPic          string `json:"item_pic"`
		ItemContext      string `json:"item_context"`
		UseScore         int64  `json:"use_score"`
		ExchangeCount    int64  `json:"exchange_count"`
		MaxExchangeCount int64  `json:"max_exchange_count"`
		Cycle            string `json:"cycle"`
	} `json:"list"`
	Total int64 `json:"total"`
}

func CommunityExchangeList(webToken string) (CommunityExchangeListData, error) {
	apiUrl := `https://gf2-bbs-api.sunborngame.com/community/item/exchange_list`

	dataBytes, err := CommunityGet(apiUrl, nil, webToken)
	if err != nil {
		return CommunityExchangeListData{}, errors.WithStack(err)
	}

	var data CommunityExchangeListData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return CommunityExchangeListData{}, errors.WithStack(err)
	}

	return data, nil
}
