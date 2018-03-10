package crafttree

import (
	"log"
)

type FinalIngredient interface {
	SetValue(*treeNode)
	IncreaseQuantity(int)
}

func (c *Crystal) IncreaseQuantity(q int) { c.Quantity += q }
func (c *Crystal) SetValue(n *treeNode)   {}

func (c *PreCraft) IncreaseQuantity(q int) { c.Quantity += q }
func (c *PreCraft) SetValue(n *treeNode) {
	var cs []*CrafterClass
	for _, c := range n.craftingValue.classes {
		cs = append(cs, &CrafterClass{
			Level: n.craftingValue.recipeLevel,
			Name:  c.class.name,
			Icon:  c.class.icon,
		})
	}
	log.Printf("setting value for precraft with ID %d, classes: %v", n.id, cs)
	c.Classes = cs

	if n.otherValue != nil {
		var oms []*ObtainMethod
		for _, o := range n.otherValue.obtainMethods {

			var od []*ObtainMethodData
			for _, data := range o.data {
				od = append(od, &ObtainMethodData{
					Name:      data.name,
					Locations: data.locations,
				})
			}

			oms = append(oms, &ObtainMethod{
				Data: od,
				Type: o.obtainType,
			})
		}

		c.Obtain = oms
	}
}

func (c *GatheringIngredient) IncreaseQuantity(q int) { c.Quantity += q }
func (c *GatheringIngredient) SetValue(n *treeNode) {
	c.Type = &GatheringType{
		Icon: n.gatheringValue.gatherType.icon,
		Name: n.gatheringValue.gatherType.name,
	}

	if n.otherValue != nil {
		var oms []*ObtainMethod
		for _, o := range n.otherValue.obtainMethods {

			var od []*ObtainMethodData
			for _, data := range o.data {
				od = append(od, &ObtainMethodData{
					Name:      data.name,
					Locations: data.locations,
				})
			}

			oms = append(oms, &ObtainMethod{
				Data: od,
				Type: o.obtainType,
			})
		}

		c.Obtain = oms
	}
}

func (c *OtherIngredient) IncreaseQuantity(q int) { c.Quantity += q }
func (c *OtherIngredient) SetValue(n *treeNode) {
	var oms []*ObtainMethod
	for _, o := range n.otherValue.obtainMethods {

		var od []*ObtainMethodData
		for _, data := range o.data {
			od = append(od, &ObtainMethodData{
				Name:      data.name,
				Locations: data.locations,
			})
		}

		oms = append(oms, &ObtainMethod{
			Data: od,
			Type: o.obtainType,
		})
	}

	c.Obtain = oms
}

func NewIngredient(n *treeNode) FinalIngredient {
	ing := &Ingredient{
		ID:       n.id,
		Icon:     n.icon,
		Name:     n.name,
		Quantity: n.requiredQuantity,
	}

	// TODO: implement crystals!!!
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
