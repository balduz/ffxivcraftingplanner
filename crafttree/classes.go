package crafttree

type class struct {
	name string
	icon string
	abbr string
}

var classesMap map[string]*class

func getClass(name string) *class {
	if val, ok := classesMap[name]; ok {
		return val
	}
	return nil
}

func initializeClassesMap() {
	classesMap = map[string]*class{
		"Carpenter": &class{
			name: "Carpenter",
			icon: "https://secure.xivdb.com/img/classes/set2/carpenter.png",
			abbr: "CRP",
		},
		"Armorer": &class{
			name: "Armorer",
			icon: "https://secure.xivdb.com/img/classes/set2/armorer.png",
			abbr: "ARM",
		},
		"Blacksmith": &class{
			name: "Blacksmith",
			icon: "https://secure.xivdb.com/img/classes/set2/blacksmith.png",
			abbr: "BSM",
		},
		"Goldsmith": &class{
			name: "Goldsmith",
			icon: "https://secure.xivdb.com/img/classes/set2/goldsmith.png",
			abbr: "GSM",
		},
		"Leatherworker": &class{
			name: "Leatherworker",
			icon: "https://secure.xivdb.com/img/classes/set2/leatherworker.png",
			abbr: "LTW",
		},
		"Weaver": &class{
			name: "Weaver",
			icon: "https://secure.xivdb.com/img/classes/set2/weaver.png",
			abbr: "WVR",
		},
		"Alchemist": &class{
			name: "Alchemist",
			icon: "https://secure.xivdb.com/img/classes/set2/alchemist.png",
			abbr: "ALC",
		},
		"Culinarian": &class{
			name: "Culinarian",
			icon: "https://secure.xivdb.com/img/classes/set2/culinarian.png",
			abbr: "CUL",
		},
		"Botanist": &class{
			name: "Botanist",
			icon: "https://secure.xivdb.com/img/classes/set2/botanist.png",
			abbr: "BTN",
		},
		"Miner": &class{
			name: "Miner",
			icon: "https://secure.xivdb.com/img/classes/set2/miner.png",
			abbr: "MIN",
		},
		"Fisher": &class{
			name: "Fisher",
			icon: "https://secure.xivdb.com/img/classes/set2/fisher.png",
			abbr: "FSH",
		},
	}
}
