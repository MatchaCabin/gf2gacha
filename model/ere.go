package model

type Ere struct {
	Result  [][]interface{} `json:"result"`
	Time    int64           `json:"time"`
	TypeMap [][]string      `json:"typeMap"`
	Account string          `json:"account"`
}
