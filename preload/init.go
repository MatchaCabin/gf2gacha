package preload

import (
	"gf2gacha/pb"
	"gf2gacha/util"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	LangMap     = make(map[int64]string)
	ItemMap     = make(map[int64]*pb.ItemDataUnit)
	UpItemMap   = make(map[int64]int64)
	ItemRankMap = make(map[int64]map[int64]int64)

	DollNameMapping   = make(map[string]int64)
	WeaponNameMapping = make(map[string]int64)
)

func init() {
	info, err := util.GetLogInfo()
	if err != nil {
		logrus.Panic(err)
	}

	var langCNData pb.LangPackageTableCnData
	err = util.GetTableData(info.TablePath, &langCNData)
	if err != nil {
		logrus.Panic(err)
	}
	for i, unit := range langCNData.Units {
		LangMap[unit.Id] = langCNData.Units[i].Content
	}

	var itemData pb.ItemData
	err = util.GetTableData(info.TablePath, &itemData)
	if err != nil {
		logrus.Panic(err)
	}
	for i, item := range itemData.Units {
		ItemMap[item.Id] = itemData.Units[i]
		switch item.Type {
		case 10:
			DollNameMapping[LangMap[item.Name.Id]] = item.Id
		case 20:
			WeaponNameMapping[LangMap[item.Name.Id]] = item.Id
		}
	}

	var gachaData pb.GachaData
	err = util.GetTableData(info.TablePath, &gachaData)
	if err != nil {
		logrus.Panic(err)
	}
	for _, unit := range gachaData.Units {
		//等级为5在卡池中不一定算5星，比如神秘箱中的导体5/6，因此需要从卡池道具列表中提取等级
		ItemRankMap[unit.Id] = make(map[int64]int64)
		//提取人形,散爆这里居然有中文冒号你敢信
		rateDesGun := strings.ReplaceAll(unit.RateDesGun, "：", ":")
		rateDollStringGroupList := strings.Split(rateDesGun, ",")
		for _, rateDollStringGroup := range rateDollStringGroupList {
			var rank int64
			var dollStringGroup string
			switch {
			case strings.HasPrefix(rateDollStringGroup, "5:"):
				rank = 5
				dollStringGroup = strings.TrimPrefix(rateDollStringGroup, "5:")
			case strings.HasPrefix(rateDollStringGroup, "4:"):
				rank = 4
				dollStringGroup = strings.TrimPrefix(rateDollStringGroup, "4:")
			case strings.HasPrefix(rateDollStringGroup, "3:"):
				rank = 3
				dollStringGroup = strings.TrimPrefix(rateDollStringGroup, "3:")
			default:
				if rateDollStringGroup != "" {
					logrus.Errorf("未知的人形掉落列表:%s", rateDollStringGroup)
				}
				continue
			}
			dollStringList := strings.Split(dollStringGroup, ":")
			for _, dollString := range dollStringList {
				dollId, err := strconv.ParseInt(dollString, 10, 64)
				if err != nil {
					logrus.Panic(err)
				}
				ItemRankMap[unit.Id][dollId] = rank
			}
		}
		//提取武器,散爆这里居然有中文冒号你敢信
		rateDesWeapon := strings.ReplaceAll(unit.RateDesWeapon, "：", ":")
		rateWeaponStringGroupList := strings.Split(rateDesWeapon, ",")
		for _, rateWeaponStringGroup := range rateWeaponStringGroupList {
			var rank int64
			var weaponStringGroup string
			switch {
			case strings.HasPrefix(rateWeaponStringGroup, "5:"):
				rank = 5
				weaponStringGroup = strings.TrimPrefix(rateWeaponStringGroup, "5:")
			case strings.HasPrefix(rateWeaponStringGroup, "4:"):
				rank = 4
				weaponStringGroup = strings.TrimPrefix(rateWeaponStringGroup, "4:")
			case strings.HasPrefix(rateWeaponStringGroup, "3:"):
				rank = 3
				weaponStringGroup = strings.TrimPrefix(rateWeaponStringGroup, "3:")
			default:
				if rateWeaponStringGroup != "" {
					logrus.Errorf("未知的道具等级列表:%s", rateWeaponStringGroup)
				}
				continue
			}
			weaponStringList := strings.Split(weaponStringGroup, ":")
			for _, weaponString := range weaponStringList {
				weaponId, err := strconv.ParseInt(weaponString, 10, 64)
				if err != nil {
					logrus.Panic(err)
				}
				ItemRankMap[unit.Id][weaponId] = rank
			}
		}

		//限定池需要提取upItem来分辨是否歪
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
						logrus.Panic(err)
					}
					UpItemMap[unit.Id] = upItemId
					break
				}
			}
		}
	}
}
