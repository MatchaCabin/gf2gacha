package model

import "encoding/json"

type ResponseBody struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type ResponseData struct {
	RecordList []RemoteRecord `json:"list"`
	Next       string         `json:"next"`
}
