package model

type LogInfo struct {
	TablePath   string `json:"tablePath"`
	AccessToken string `json:"accessToken"`
	Uid         string `json:"uid"`
	GachaUrl    string `json:"gachaUrl"`
}
