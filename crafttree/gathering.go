package crafttree

type gatherType struct {
	name string
	icon string
}

var gatheringTypesMap map[string]*gatherType

func getGatheringType(name string) *gatherType {
	if val, ok := gatheringTypesMap[name]; ok {
		return val
	}
	return nil
}

func initializeGatheringTypesMap() {
	gatheringTypesMap = map[string]*gatherType{
		"Harvesting": &gatherType{
			name: "Harvesting",
			icon: "/web/assets/icons/harvesting.png",
		},
		"Logging": &gatherType{
			name: "Logging",
			icon: "/web/assets/icons/logging.png",
		},
		"Quarrying": &gatherType{
			name: "Quarrying",
			icon: "/web/assets/icons/quarrying.png",
		},
		"Mining": &gatherType{
			name: "Mining",
			icon: "/web/assets/icons/mining.png",
		},
		"Fishing": &gatherType{
			name: "Fishing",
			icon: "/web/assets/icons/logging.png",
		},
	}
}
