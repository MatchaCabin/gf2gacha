package request

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

type CommunityTopicListData struct {
	LastTid int64 `json:"last_tid"`
	List    []struct {
		CategoryId   int64    `json:"category_id"`
		CategoryName string   `json:"category_name"`
		CommentNum   int64    `json:"comment_num"`
		Content      string   `json:"content"`
		CreateTime   string   `json:"create_time"`
		IsAdmin      bool     `json:"is_admin"`
		IsAuthor     bool     `json:"is_author"`
		IsFavor      bool     `json:"is_favor"`
		IsLike       bool     `json:"is_like"`
		LikeNum      int64    `json:"like_num"`
		PicList      []string `json:"pic_list"`
		Title        string   `json:"title"`
		TopicId      int64    `json:"topic_id"`
		UserAvatar   string   `json:"user_avatar"`
		UserId       int64    `json:"user_id"`
		UserLevel    int64    `json:"user_level"`
		UserNickName string   `json:"user_nick_name"`
	} `json:"list"`
	NextPage bool  `json:"next_page"`
	PubTime  int64 `json:"pub_time"`
	Total    int64 `json:"total"`
}

func CommunityTopicList(webToken string, userId int64) (CommunityTopicListData, error) {
	apiUrl := `https://gf2-bbs-api.sunborngame.com/community/topic/list`

	values := url.Values{
		"last_tid":   []string{"0"},
		"pub_time":   []string{"0"},
		"reply_time": []string{"0"},
		"hot_value":  []string{"0"},
	}
	if userId > 0 {
		values.Set("sort_type", "1")
		values.Set("category_id", "100")
		values.Set("query_type", "4")
		values.Set("user_id", strconv.FormatInt(userId, 10))
	} else {
		values.Set("sort_type", "2")
		values.Set("category_id", "5")
		values.Set("query_type", "1")
	}

	dataBytes, err := CommunityGet(apiUrl, values, webToken)
	if err != nil {
		return CommunityTopicListData{}, errors.WithStack(err)
	}

	var data CommunityTopicListData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return CommunityTopicListData{}, errors.WithStack(err)
	}

	return data, nil
}
