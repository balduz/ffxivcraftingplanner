package crafttree

import "ffxivcraftingplanner/gamedata"

type precraft struct {
	id          int
	icon        string
	quantity    int
	yield       int
	ingredients []*precraft
	classes     []*gamedata.Class
}

type gatheringIngredient struct {
	id       int
	icon     string
	quantity int
	class    string
	kind     string
	location string
	mapName  string
	level    int
}

type otherIngredient struct {
	id            int
	icon          string
	quantity      string
	obtainMethods []*obtainMethod
}

type craftPhase map[int]precraft
type crafts []*craftPhase

type CraftingList struct {
	Crystals  []*Crystal
	PreCrafts []*PreCraft
	Gathering []*GatheringIngredient
	Other     []*OtherIngredient
}
