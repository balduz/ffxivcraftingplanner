package crafttree

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

type ingredient interface {
	setValue(*treeNode)
	increaseQuantity(int)
}

func (ing *Ingredient) increaseQuantity(q int) {
	ing.Quantity += q
}

func (c *Crystal) setValue(n *treeNode) {
	// TODO: is there anything to be done with crystals?
}

func (c *PreCraft) setValue(n *treeNode) {
	var cs []*CrafterClass
	for _, c := range n.craftingValue.classes {
		cs = append(cs, &CrafterClass{
			Level: n.craftingValue.recipeLevel,
			Name:  c.class.Name,
			Icon:  c.class.Icon,
		})
	}
	c.Classes = cs

	if n.otherValue != nil {
		c.Obtain = getObtainMethods(n.otherValue)
	}
}

func (c *GatheringIngredient) setValue(n *treeNode) {
	c.Type = &GatheringType{
		Icon: n.gatheringValue.gatherType.Icon,
		Name: n.gatheringValue.gatherType.Name,
	}

	if n.otherValue != nil {
		c.Obtain = getObtainMethods(n.otherValue)
	}
}

func (c *OtherIngredient) setValue(n *treeNode) {
	if n.otherValue != nil {
		c.Obtain = getObtainMethods(n.otherValue)
	}
}

func NewIngredient(n *treeNode) ingredient {
	ing := &Ingredient{
		ID:       n.id,
		Icon:     n.icon,
		Name:     n.name,
		Quantity: n.requiredQuantity,
	}

	if n.isCrystal {
		return &Crystal{ing}
	}
	if n.craftingValue != nil {
		return &PreCraft{ing, nil, nil}
	}
	if n.gatheringValue != nil {
		return &GatheringIngredient{ing, nil, nil}
	}
	return &OtherIngredient{ing, nil}
}
