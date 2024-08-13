package request

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

type CommunityTopicViewData struct {
	BadNum          int64    `json:"bad_num"`
	CategoryId      int64    `json:"category_id"`
	CategoryName    string   `json:"category_name"`
	CommentNum      int64    `json:"comment_num"`
	Content         string   `json:"content"`
	CreateTime      string   `json:"create_time"`
	IpLocation      string   `json:"ip_location"`
	IsAdmin         bool     `json:"is_admin"`
	IsAuthor        bool     `json:"is_author"`
	IsBad           bool     `json:"is_bad"`
	IsFavor         bool     `json:"is_favor"`
	IsFollow        bool     `json:"is_follow"`
	IsLike          bool     `json:"is_like"`
	LikeNum         int64    `json:"like_num"`
	LikeUserAvatars []string `json:"like_user_avatars"`
	PicList         []string `json:"pic_list"`
	Title           string   `json:"title"`
	TopicId         int64    `json:"topic_id"`
	UpdateTime      string   `json:"update_time"`
	UserAvatar      string   `json:"user_avatar"`
	UserId          int64    `json:"user_id"`
	UserLevel       int64    `json:"user_level"`
	UserNickName    string   `json:"user_nick_name"`
	ViewNum         int64    `json:"view_num"`
}

func CommunityTopicView(webToken string, topicId int64) (CommunityTopicViewData, error) {
	apiUrl := fmt.Sprintf(`https://gf2-bbs-api.sunborngame.com/community/topic/%s`, topicId)

	values := url.Values{
		"id": []string{strconv.FormatInt(topicId, 10)},
	}

	dataBytes, err := CommunityGet(apiUrl, values, webToken)
	if err != nil {
		return CommunityTopicViewData{}, errors.WithStack(err)
	}

	var data CommunityTopicViewData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return CommunityTopicViewData{}, errors.WithStack(err)
	}

	return data, nil
}
