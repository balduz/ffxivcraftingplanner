package crafttree

import (
	"log"
	"sort"
)

type ObtainType int

const (
	CRAFT        ObtainType = iota
	ENEMY        ObtainType = iota
	NPC          ObtainType = iota
	SPECIAL_SHOP ObtainType = iota
	DUNGEON      ObtainType = iota
	TRIAL        ObtainType = iota
)

type CurrencyType int

const (
	GIL CurrencyType = iota
)

type craftingList struct {
	ingredients map[int]ingredient
}

type byTier []*PreCraft

func (a byTier) Len() int           { return len(a) }
func (a byTier) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byTier) Less(i, j int) bool { return a[i].Tier > a[j].Tier }

func GetCraftingListFor(ids []int) *CraftingList {
	var trees []*treeNode
	for _, id := range ids {
		tree, err := buildTreeNode(id, 1)
		if err != nil {
			log.Printf("GetCraftingList(_): error getting tree for id %d, error %s", id, err)
		}
		trees = append(trees, tree)
	}

	cl := craftingList{
		ingredients: make(map[int]ingredient),
	}
	for _, tree := range trees {
		cl.visitTree(tree, 1, 1)
	}

	c := []*Crystal{}
	pc := []*PreCraft{}
	g := []*GatheringIngredient{}
	o := []*OtherIngredient{}

	for _, ing := range cl.ingredients {
		switch ing.(type) {
		case *Crystal:
			c = append(c, ing.(*Crystal))
		case *PreCraft:
			pc = append(pc, ing.(*PreCraft))
		case *GatheringIngredient:
			g = append(g, ing.(*GatheringIngredient))
		case *OtherIngredient:
			o = append(o, ing.(*OtherIngredient))
		}
	}

	sort.Sort(byTier(pc))

	return &CraftingList{
		Crystals:  c,
		PreCrafts: pc,
		Gathering: g,
		Other:     o,
	}
}

func (cl *craftingList) visitTree(n *treeNode, quantity, tier int) {
	newQuantity := quantity * n.requiredQuantity
	if val, ok := cl.ingredients[n.id]; ok {
		val.increaseQuantity(newQuantity)
		val.setTier(tier)
	} else {
		i := NewIngredient(n, tier)
		i.setValue(n)
		cl.ingredients[n.id] = i
	}

	for _, child := range n.children {
		cl.visitTree(child, newQuantity, tier+1)
	}
}
