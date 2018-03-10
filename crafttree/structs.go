package crafttree

type precraft struct {
	id          int
	icon        string
	quantity    int
	yield       int
	ingredients []*precraft
	classes     []*class
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

type Ingredient struct {
	ID       int
	Icon     string
	Name     string
	Quantity int
}

type ObtainMethodData struct {
	Name      string
	Locations []*location
}

type ObtainMethod struct {
	Data []*ObtainMethodData
	Type ObtainType
	// TODO: add for shops!
}

type Crystal struct {
	*Ingredient
}

type CrafterClass struct {
	Name  string
	Level int
	Icon  string
}

type PreCraft struct {
	*Ingredient
	Classes []*CrafterClass
	Obtain  []*ObtainMethod
}

type GatheringType struct {
	Name string
	Icon string
}

type GatheringIngredient struct {
	*Ingredient
	Type   *GatheringType
	Obtain []*ObtainMethod
}

type OtherIngredient struct {
	*Ingredient
	Obtain []*ObtainMethod
}

type CraftingList struct {
	Crystals  []*Crystal
	PreCrafts []*PreCraft
	Gathering []*GatheringIngredient
	Other     []*OtherIngredient
}
