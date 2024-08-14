package logic

import (
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func MergeEreRecord(uid, erePath, typ string) error {
	ereFileData, err := os.ReadFile(erePath)
	if err != nil {
		return errors.WithStack(err)
	}

	var ereRecordList []model.LocalRecord

	switch strings.ToLower(typ) {
	case "json":
		ereRecordList, err = ParseEreJsonData(ereFileData)
		if err != nil {
			return errors.WithStack(err)
		}
	case "excel":
		ereRecordList, err = ParseEreExcelData(ereFileData)
		if err != nil {
			return errors.WithStack(err)
		}
	default:
		return errors.Errorf("unknown ere type: %s", typ)
	}

	//解析成功后先做备份
	err = util.BackupDB()
	if err != nil {
		return errors.WithStack(err)
	}

	var erePool1RecordList, erePool3RecordList, erePool4RecordList, erePool5RecordList []model.LocalRecord
	//需要每个池子分开处理
	for i, record := range ereRecordList {
		switch record.PoolType {
		case 1:
			erePool1RecordList = append(erePool1RecordList, ereRecordList[i])
		case 3:
			erePool3RecordList = append(erePool3RecordList, ereRecordList[i])
		case 4:
			erePool4RecordList = append(erePool4RecordList, ereRecordList[i])
		case 5:
			erePool5RecordList = append(erePool5RecordList, ereRecordList[i])
		}
	}

	err = mergeEreJsonRecord(uid, 1, erePool1RecordList)
	if err != nil {
		return errors.WithStack(err)
	}
	err = mergeEreJsonRecord(uid, 3, erePool3RecordList)
	if err != nil {
		return errors.WithStack(err)
	}
	err = mergeEreJsonRecord(uid, 4, erePool4RecordList)
	if err != nil {
		return errors.WithStack(err)
	}
	err = mergeEreJsonRecord(uid, 5, erePool5RecordList)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func mergeEreJsonRecord(uid string, poolType int64, ereRecordList []model.LocalRecord) error {
	if len(ereRecordList) == 0 {
		return nil
	}

	localRecordList, err := GetLocalRecord(uid, poolType)
	if err != nil {
		return errors.WithStack(err)
	}

	//如果本地无记录，直接导入
	if len(localRecordList) == 0 {
		err = SaveLocalRecord(uid, ereRecordList)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	firstLocalRecord := localRecordList[0]
	firstLocalRecordSequence := 0
	//判断本地记录是十连中的第几条
	if len(localRecordList) > 1 {
		for i, localRecord := range localRecordList[1:] {
			if localRecord.GachaTimestamp != firstLocalRecord.GachaTimestamp {
				if i > 0 {
					firstLocalRecordSequence = 9 - i
				}
				break
			}
		}
	}
	logrus.Infof("本地第一条:%+v, seq:%d", firstLocalRecord, firstLocalRecordSequence)

	var mergeLocalRecordList []model.LocalRecord

	var previousEreReocrdTimestamp int64
	ereRecordSequence := 0
	for i, ereRecord := range ereRecordList {
		if ereRecord.GachaTimestamp == previousEreReocrdTimestamp {
			ereRecordSequence++
		} else {
			ereRecordSequence = 0
		}

		//抽卡时间一致，道具一致，十连中的序号一致，则判断为同一条
		if ereRecord.GachaTimestamp == firstLocalRecord.GachaTimestamp && ereRecord.ItemId == firstLocalRecord.ItemId && ereRecordSequence == firstLocalRecordSequence {
			logrus.Infof("ERE最后一条:%+v, seq:%d", ereRecord, ereRecordSequence)
			break
		}
		mergeLocalRecordList = append(mergeLocalRecordList, ereRecordList[i])
		previousEreReocrdTimestamp = ereRecord.GachaTimestamp
	}

	//主要为了消除Id
	for _, record := range localRecordList {
		mergeLocalRecordList = append(mergeLocalRecordList, model.LocalRecord{
			PoolType:       record.PoolType,
			PoolId:         record.PoolId,
			ItemId:         record.ItemId,
			GachaTimestamp: record.GachaTimestamp,
		})
	}

	err = RemoveLocalRecord(uid, poolType)
	if err != nil {
		return errors.WithStack(err)
	}

	err = SaveLocalRecord(uid, mergeLocalRecordList)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
