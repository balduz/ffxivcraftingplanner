package gamedata

type GatherType struct {
	Name string
	Icon string
}

var gatheringTypesMap map[string]*GatherType

func GetGatheringType(name string) *GatherType {
	if val, ok := gatheringTypesMap[name]; ok {
		return val
	}
	return nil
}

func initializeGatheringTypesMap() {
	gatheringTypesMap = map[string]*GatherType{
		"Harvesting": {
			Name: "Harvesting",
			Icon: "/web/assets/icons/harvesting.png",
		},
		"Logging": {
			Name: "Logging",
			Icon: "/web/assets/icons/logging.png",
		},
		"Quarrying": {
			Name: "Quarrying",
			Icon: "/web/assets/icons/quarrying.png",
		},
		"Mining": {
			Name: "Mining",
			Icon: "/web/assets/icons/mining.png",
		},
		"Fishing": {
			Name: "Fishing",
			Icon: "/web/assets/icons/logging.png",
		},
	}
}
