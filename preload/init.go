package preload

import (
	"gf2gacha/pb"
	"gf2gacha/util"
	"strconv"
	"strings"
)

var (
	LangMap   = make(map[int64]string)
	ItemMap   = make(map[int64]*pb.ItemDataUnit)
	UpItemMap = make(map[int64]int64)
)

func init() {
	info, err := util.GetGameInfo()
	if err != nil {
		panic(err)
	}

	var langCNData pb.LangPackageTableCnData
	err = util.GetTableData(info.TablePath, &langCNData)
	if err != nil {
		panic(err)
	}
	for i, unit := range langCNData.Units {
		LangMap[unit.Id] = langCNData.Units[i].Content
	}

	var itemData pb.ItemData
	err = util.GetTableData(info.TablePath, &itemData)
	if err != nil {
		panic(err)
	}
	for i, unit := range itemData.Units {
		ItemMap[unit.Id] = itemData.Units[i]
	}

	var gachaData pb.GachaData
	err = util.GetTableData(info.TablePath, &gachaData)
	if err != nil {
		panic(err)
	}
	for _, unit := range gachaData.Units {
		if unit.Type == 3 || unit.Type == 4 {
			var upItemGroupString string
			if unit.GunUpItem != "" {
				upItemGroupString = unit.GunUpItem
			} else {
				upItemGroupString = unit.WeaponUpItem
			}
			upItemGroupList := strings.Split(upItemGroupString, ",")
			for _, upItemGroup := range upItemGroupList {
				if strings.HasPrefix(upItemGroup, "5:") {
					upItemString := strings.TrimPrefix(upItemGroup, "5:")
					upItemId, err := strconv.ParseInt(upItemString, 10, 64)
					if err != nil {
						panic(err)
					}
					UpItemMap[unit.Id] = upItemId
					break
				}
			}
		}
	}
}
