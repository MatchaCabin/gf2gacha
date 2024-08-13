package request

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type CommunityUserInfoData struct {
	User struct {
		AuthLock           int64  `json:"auth_lock"`
		AuthType           int64  `json:"auth_type"`
		Avatar             string `json:"avatar"`
		Exp                int64  `json:"exp"`
		Fans               int64  `json:"fans"`
		Favors             int64  `json:"favors"`
		Follows            int64  `json:"follows"`
		GameCommanderLevel int64  `json:"game_commander_level"`
		GameNickName       string `json:"game_nick_name"`
		GameUid            int64  `json:"game_uid"`
		IpLocation         string `json:"ip_location"`
		IsAdmin            bool   `json:"is_admin"`
		IsAuthor           bool   `json:"is_author"`
		IsFollow           bool   `json:"is_follow"`
		Level              int64  `json:"level"`
		Likes              int64  `json:"likes"`
		NextLvExp          int64  `json:"next_lv_exp"`
		NickName           string `json:"nick_name"`
		Score              int64  `json:"score"`
		Signature          string `json:"signature"`
		Uid                int64  `json:"uid"`
	} `json:"user"`
}

func CommunityUserInfo(webToken string) (CommunityUserInfoData, error) {
	apiUrl := `https://gf2-bbs-api.sunborngame.com/community/member/info`
	params := map[string]interface{}{
		"uid": 0,
	}

	dataBytes, err := CommunityPost(apiUrl, params, webToken)
	if err != nil {
		return CommunityUserInfoData{}, errors.WithStack(err)
	}

	var data CommunityUserInfoData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return CommunityUserInfoData{}, errors.WithStack(err)
	}

	return data, nil
}
