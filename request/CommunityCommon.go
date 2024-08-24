package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

type CommonResponse struct {
	Code    int             `json:"Code"`
	Message string          `json:"Message"`
	Data    json.RawMessage `json:"data"`
}

func (c CommonResponse) Error() string {
	return fmt.Sprintf("%s(Code:%d)", c.Message, c.Code)
}

func CommunityGet(apiUrl string, values url.Values, webToken string) (dataBytes json.RawMessage, err error) {
	apiUrl += "?" + values.Encode()
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", "https://gf2-bbs.sunborngame.com")
	request.Header.Set("Referer", "https://gf2-bbs.sunborngame.com")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0")
	if webToken != "" {
		request.Header.Set("Authorization", webToken)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var respBody CommonResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if respBody.Code != 0 {
		return nil, errors.WithStack(respBody)
	}

	return respBody.Data, nil
}

func CommunityPost(apiUrl string, params map[string]interface{}, webToken string) (dataBytes json.RawMessage, err error) {
	requestBodyBytes, err := json.Marshal(&params)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", "https://gf2-bbs.sunborngame.com")
	request.Header.Set("Referer", "https://gf2-bbs.sunborngame.com")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0")
	if webToken != "" {
		request.Header.Set("Authorization", webToken)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var respBody CommonResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if respBody.Code != 0 {
		return nil, errors.WithStack(respBody)
	}

	return respBody.Data, nil
}
