package xivdb

import "encoding/json"

type Item struct {
	CraftingRecipe []*Recipe `json:"craftable"`
	CategoryName   string    `json:"category_name"`
	Enemies        []*Enemy
	Gathering      []*Gathering
	Help           string
	Icon           string
	ID             int
	Instances      []*Instance
	KindName       string `json:"kind_name"`
	Leves          []*Leve
	Name           string
	PriceHigh      int `json:"price_high"`
	PriceLow       int `json:"price_low"`
	PriceMid       int `json:"price_mid"`
	Quests         []*Quest
	//Shops          []*shop
	SpecialShops []*SpecialShop `json:"special_shops_obtain"`
}

type Enemy struct {
	Icon    string
	ID      int
	MapData *MapData `json:"map_data"`
	//MapPrimary *mapPrimary `json:"map_primary"`
	Name string
}

type Gathering struct {
	Kind  string `json:"type_name"`
	Stars int
	Nodes json.RawMessage
}

type Node struct {
	ID    int
	Level int
	Name  string
}

type Recipe struct {
	ClassName             string `json:"class_name"`
	CraftQuantity         int    `json:"craft_quantity"`
	Difficulty            int
	Icon                  string
	ID                    int
	IsSpecialization      int `json:"is_specialization_required"`
	Item                  *Item
	ItemName              string `json:"item_name"`
	Masterbook            *Masterbook
	Name                  string
	Quality               int
	RequiredControl       int `json:"required_control"`
	RequiredCraftsmanship int `json:"required_craftsmanship"`
	RequiredLevel         int `json:"level_view"`
	Stars                 int
	Tree                  []*Ingredient
}

type Ingredient struct {
	Quantity int
	ID       int
}

type Masterbook struct {
	Help string
	Icon string
	Name string
}

type Quest struct {
	ID   int
	Name string
}

type Instance struct {
	Name        string
	ContentName string `json:"content_name"`
	ID          int
}

type Leve struct {
	Name string
	ID   int
}

type Shop struct {
	Name string
	Npc  *Npc
}

type Npc struct {
	Name       string
	MapData    *MapData    `json:"map_data"`
	MapPrimary *MapPrimary `json:"map_primary"`
}

type SpecialShop struct {
	Name  string
	Items *xvidbSpecialShopTradeData
}

type xvidbSpecialShopTradeData struct {
	Cost1         int   `json:"cost_count_1"`
	Cost2         int   `json:"cost_count_2"`
	Cost3         int   `json:"cost_count_3"`
	Item1         *Item `json:"cost_item_1"`
	Item2         *Item `json:"cost_item_2"`
	Item3         *Item `json:"cost_item_3"`
	Receive1      int   `json:"receive_count_1"`
	Receive2      int   `json:"receive_count_2"`
	ReceivedItem1 *Item `json:"receive_item_1"`
	ReceivedItem2 *Item `json:"receive_item_2"`
}

type MapPrimary struct {
	Placename *Placename
	Position  *Position
}

type Placename struct {
	ID   int
	Name string
	Maps *PlacenameMap
}

type PlacenameMap struct {
	PlacenameID int `json:"id"`
	RegionID    int `json:"region"`
}

type Position struct {
	X float32
	Y float32
}

type MapData struct {
	Maps   []*MapDataMap
	Points []*MapPoint
}

type MapDataMap struct {
	PlacenameID   int    `json:"placename_id"`
	PlacenameName string `json:"placename_name"`
	RegionID      int    `json:"region_id"`
	RegionName    string `json:"region_name"`
}

type MapPoint struct {
	Position      *xivdAppPosition `json:"app_position"`
	PlacenameID   int              `json:"placename_id"`
	PlacenameName string           `json:"placename_name"`
	RegionID      int              `json:"region_id"`
	RegionName    string           `json:"region_name"`
}

type xivdAppPosition struct {
	Position *Position `json:"ingame"`
}

type RecipesSearch struct {
	Recipes RecipeResults
}

type RecipeResults struct {
	Results []RecipeResult
}

type RecipeResult struct {
	ClassName  string `json:"class_name"`
	Icon       string
	ID         int
	ItemName   string `json:"item_name"`
	Masterbook string
	Name       string
	Stars      int
	StarsHTML  string `json:"stars_html"`
}
