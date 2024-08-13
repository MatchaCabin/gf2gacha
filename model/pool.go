package model

type Pool struct {
	PoolType        int64           `json:"poolType"`
	GachaCount      int64           `json:"gachaCount"`
	LoseCount       int64           `json:"loseCount"`
	GuaranteesCount int64           `json:"guaranteesCount"`
	Rank5Count      int64           `json:"rank5Count"`
	Rank4Count      int64           `json:"rank4Count"`
	Rank3Count      int64           `json:"rank3Count"`
	StoredCount     int64           `json:"storedCount"`
	RecordList      []DisplayRecord `json:"recordList"`
}
