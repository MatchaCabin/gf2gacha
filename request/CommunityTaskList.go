package request

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type CommunityTaskListData struct {
	DailyTask []struct {
		TaskName         string `json:"task_name"`
		TaskContext      string `json:"task_context"`
		CompleteCount    int64  `json:"complete_count"`
		MaxCompleteCount int64  `json:"max_complete_count"`
	} `json:"daily_task"`
	MoreTask []struct {
		TaskName         string `json:"task_name"`
		TaskContext      string `json:"task_context"`
		CompleteCount    int64  `json:"complete_count"`
		MaxCompleteCount int64  `json:"max_complete_count"`
	} `json:"more_task"`
}

func CommunityTaskList(webToken string) (CommunityTaskListData, error) {
	apiUrl := `https://gf2-bbs-api.sunborngame.com/community/task/get_current_task_list`

	dataBytes, err := CommunityGet(apiUrl, nil, webToken)
	if err != nil {
		return CommunityTaskListData{}, errors.WithStack(err)
	}

	var data CommunityTaskListData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return CommunityTaskListData{}, errors.WithStack(err)
	}

	return data, nil
}
