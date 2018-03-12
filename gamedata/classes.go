package gamedata

type Class struct {
	Name string
	Icon string
	Abbr string
}

var classesMap map[string]*Class

func GetClass(name string) *Class {
	if val, ok := classesMap[name]; ok {
		return val
	}
	return nil
}

func initializeClassesMap() {
	classesMap = map[string]*Class{
		"Carpenter": {
			Name: "Carpenter",
			Icon: "https://secure.xivdb.com/img/classes/set2/carpenter.png",
			Abbr: "CRP",
		},
		"Armorer": {
			Name: "Armorer",
			Icon: "https://secure.xivdb.com/img/classes/set2/armorer.png",
			Abbr: "ARM",
		},
		"Blacksmith": {
			Name: "Blacksmith",
			Icon: "https://secure.xivdb.com/img/classes/set2/blacksmith.png",
			Abbr: "BSM",
		},
		"Goldsmith": {
			Name: "Goldsmith",
			Icon: "https://secure.xivdb.com/img/classes/set2/goldsmith.png",
			Abbr: "GSM",
		},
		"Leatherworker": {
			Name: "Leatherworker",
			Icon: "https://secure.xivdb.com/img/classes/set2/leatherworker.png",
			Abbr: "LTW",
		},
		"Weaver": {
			Name: "Weaver",
			Icon: "https://secure.xivdb.com/img/classes/set2/weaver.png",
			Abbr: "WVR",
		},
		"Alchemist": {
			Name: "Alchemist",
			Icon: "https://secure.xivdb.com/img/classes/set2/alchemist.png",
			Abbr: "ALC",
		},
		"Culinarian": {
			Name: "Culinarian",
			Icon: "https://secure.xivdb.com/img/classes/set2/culinarian.png",
			Abbr: "CUL",
		},
		"Botanist": {
			Name: "Botanist",
			Icon: "https://secure.xivdb.com/img/classes/set2/botanist.png",
			Abbr: "BTN",
		},
		"Miner": {
			Name: "Miner",
			Icon: "https://secure.xivdb.com/img/classes/set2/miner.png",
			Abbr: "MIN",
		},
		"Fisher": {
			Name: "Fisher",
			Icon: "https://secure.xivdb.com/img/classes/set2/fisher.png",
			Abbr: "FSH",
		},
	}
}
