package crafttree

import (
	"ffxivcraftingplanner/xivdb"
	"fmt"
	"log"
)

type obtainData struct {
	name      string
	locations []*location
	currency  string
}

type obtainMethod struct {
	data       []*obtainData
	obtainType ObtainType
}

type position struct {
	x float32
	y float32
}

type location struct {
	mapName    string
	regionName string
	position   *position
}

type craftingClass struct {
	class *class
	level int
}

type craftingNode struct {
	yield       int
	recipeLevel int
	classes     []*craftingClass
}
type gatheringData struct {
	gatherType *gatherType
	kind       string
	nodes      []*gatheringNode
}
type gatheringNode struct {
	level   int
	mapID   int
	mapName string
}
type otherNode struct {
	obtainMethods []*obtainMethod
}

type treeNode struct {
	id   int
	name string
	icon string

	requiredQuantity int
	children         []*treeNode

	craftingValue  *craftingNode
	gatheringValue *gatheringData
	otherValue     *otherNode
	isCrystal      bool
}

func buildGatheringData(i *xivdb.Item) (*gatheringData, error) {
	if len(i.Gathering) == 0 {
		return nil, nil
	}

	// TODO: remove this, but to find out what has multiple gathering structs.
	if len(i.Gathering) > 1 {
		fmt.Printf("GATHERING: Item with id %d has %d gathering.", i.ID, len(i.Gathering))
	}

	itemNodes, _ := i.GetGatheringNodes()
	var nodes []*gatheringNode
	for _, v := range itemNodes {
		nodes = append(nodes, &gatheringNode{
			level:   v.Level,
			mapID:   v.ID,
			mapName: v.Name,
		})
	}

	gt := getGatheringType(i.Gathering[0].Kind)
	if gt == nil {
		log.Fatalf("could not get gathering type for %d, type: %s", i.ID, i.Gathering[0].Kind)
	}
	return &gatheringData{
		gatherType: gt,
		kind:       i.Gathering[0].Kind,
		nodes:      nodes,
	}, nil
}

func buildOtherNode(i *xivdb.Item) *otherNode {
	var obtainMethods []*obtainMethod

	if len(i.Enemies) > 0 {
		o := &obtainMethod{
			obtainType: ENEMY,
		}
		var obtainDataArray []*obtainData
		for _, enemy := range i.Enemies {
			e, err := xivdb.GetEnemy(enemy.ID)
			if err != nil {
				log.Printf("could not retrieve enemy with id %d, error %s", enemy.ID, err)
				continue
			}
			var locations []*location
			for _, p := range e.MapData.Points {
				locations = append(locations, &location{
					mapName:    p.PlacenameName,
					regionName: p.RegionName,
					position: &position{
						x: p.Position.Position.X,
						y: p.Position.Position.Y,
					},
				})
			}

			od := &obtainData{
				name:      enemy.Name,
				locations: locations,
			}
			obtainDataArray = append(obtainDataArray, od)
		}
		o.data = obtainDataArray
		obtainMethods = append(obtainMethods, o)
	}
	return &otherNode{
		obtainMethods: obtainMethods,
	}
}

func buildCraftingNode(i *xivdb.Item) *craftingNode {
	craftNode := &craftingNode{
		yield: i.CraftingRecipe[0].CraftQuantity,
	}

	var classes []*craftingClass
	for _, recipe := range i.CraftingRecipe {
		if c := getClass(recipe.ClassName); c != nil {
			classes = append(classes, &craftingClass{
				level: recipe.RequiredLevel,
				class: c,
			})
		} else {
			log.Fatalf("Could not find class for: %s", recipe.ClassName)
		}
	}
	craftNode.classes = classes
	return craftNode
}

func buildTreeNode(id, quantity int) (*treeNode, error) {
	if c := getCrystal(id); c != nil {
		return &treeNode{
			id:               id,
			name:             c.name,
			icon:             c.icon,
			requiredQuantity: quantity,
			isCrystal:        true,
		}, nil
	}

	data, _ := xivdb.GetItem(id)
	node := &treeNode{
		id:               id,
		name:             data.Name,
		icon:             data.Icon,
		requiredQuantity: quantity,
	}

	gNode, err := buildGatheringData(data)
	if err != nil {
		log.Fatalf("could not get gathering data for %d, error: %s", id, err)
	}
	node.gatheringValue = gNode

	node.otherValue = buildOtherNode(data)

	if data.CraftingRecipe != nil && len(data.CraftingRecipe) > 0 {
		node.craftingValue = buildCraftingNode(data)

		var children []*treeNode
		for _, ingr := range data.CraftingRecipe[0].Tree {
			tn, err := buildTreeNode(ingr.ID, ingr.Quantity)
			if err != nil {
				log.Printf("error in buildtreeNode: %s", err)
			}
			children = append(children, tn)
		}

		node.children = children
	}

	return node, nil
}
