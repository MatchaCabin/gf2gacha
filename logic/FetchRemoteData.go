package logic

import (
	"encoding/json"
	"gf2gacha/model"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func FetchRemoteData(gachaUrl, accessToken, next string, poolType int64) (data model.ResponseData, err error) {
	reqBody := url.Values{}
	if next != "" {
		reqBody.Set("next", next)
	}
	reqBody.Set("type_id", strconv.FormatInt(poolType, 10))

	request, err := http.NewRequest("POST", gachaUrl, strings.NewReader(reqBody.Encode()))
	if err != nil {
		return model.ResponseData{}, errors.WithStack(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", accessToken)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return model.ResponseData{}, errors.WithStack(err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.ResponseData{}, errors.WithStack(err)
	}
	var respBody model.ResponseBody
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		return model.ResponseData{}, errors.WithStack(err)
	}
	if respBody.Code != 0 {
		return model.ResponseData{}, errors.Errorf("%s(Code %d)", respBody.Message, respBody.Code)
	}

	err = json.Unmarshal(respBody.Data, &data)
	if err != nil {
		return model.ResponseData{}, errors.WithStack(err)
	}
	return data, nil
}
