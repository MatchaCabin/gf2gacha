package logic

import (
	"fmt"
	"gf2gacha/request"
	"gf2gacha/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

func HandleCommunityTasks() (messageList []string, err error) {
	logInfo, err := util.GetLogInfo()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	webToken, err := request.CommunityLogin(logInfo.AccessToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userInfo, err := request.CommunityUserInfo(webToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	messageList = append(messageList, fmt.Sprintf("当前用户: %s (UID:%d)", userInfo.Name, userInfo.Uid))

	taskListData, err := request.CommunityTaskList(webToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, dailyTask := range taskListData.DailyTask {
		if dailyTask.CompleteCount < dailyTask.MaxCompleteCount {
			switch dailyTask.TaskName {
			case "浏览帖子":
				viewMessageList, err := handleCommunityTaskView(webToken, dailyTask.MaxCompleteCount-dailyTask.CompleteCount)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				messageList = append(messageList, viewMessageList...)
			case "点赞帖子":
				likeMessageList, err := handleCommunityTaskLike(webToken, dailyTask.MaxCompleteCount-dailyTask.CompleteCount)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				messageList = append(messageList, likeMessageList...)
			case "分享帖子":
				shareMessageList, err := handleCommunityTaskShare(webToken, dailyTask.MaxCompleteCount-dailyTask.CompleteCount)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				messageList = append(messageList, shareMessageList...)
			default:
				logrus.Errorf("未知的社区任务%s", dailyTask.TaskName)
			}
		}
	}

	messageList = append(messageList, "社区任务和兑换完成")
	return messageList, nil
}

func handleCommunityTaskView(webToken string, times int64) (messageList []string, err error) {
	var count int64
	topicListData, err := request.CommunityTopicList(webToken, 0)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, topic := range topicListData.List {
		_, err = request.CommunityTopicView(webToken, topic.TopicId)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		messageList = append(messageList, fmt.Sprintf("浏览官方板块主题『%s』", topic.Title))

		count++
		if count == times {
			break
		}
	}

	return messageList, nil
}

func handleCommunityTaskLike(webToken string, times int64) (messageList []string, err error) {
	var count int64
	topicListData, err := request.CommunityTopicList(webToken, 0)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, topic := range topicListData.List {
		if !topic.IsLike {
			//未点赞的直接点赞
			err = request.CommunityTopicLike(webToken, topic.TopicId)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			messageList = append(messageList, fmt.Sprintf("点赞官方板块主题『%s』", topic.Title))
		} else {
			//已点赞的取消点赞再点赞
			err = request.CommunityTopicLike(webToken, topic.TopicId)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			time.Sleep(50 * time.Millisecond)
			err = request.CommunityTopicLike(webToken, topic.TopicId)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			messageList = append(messageList, fmt.Sprintf("取消并再次点赞官方板块主题『%s』", topic.Title))
		}

		count++
		if count == times {
			break
		}
	}

	return messageList, nil
}

func handleCommunityTaskShare(webToken string, times int64) (messageList []string, err error) {
	var count int64
	topicListData, err := request.CommunityTopicList(webToken, 0)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, topic := range topicListData.List {
		err = request.CommunityTopicShare(webToken, topic.TopicId)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		messageList = append(messageList, fmt.Sprintf("转发官方板块主题『%s』", topic.Title))

		count++
		if count == times {
			break
		}
	}

	return messageList, nil
}
