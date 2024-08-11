package model

type Pool struct {
	PoolType        int64
	GachaCount      int64
	LoseCount       int64
	GuaranteesCount int64
	Rank5Count      int64
	Rank4Count      int64
	Rank3Count      int64
	StoredCount     int64
	RecordList      []DisplayRecord
}
