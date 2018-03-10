package crafttree

import (
	"fmt"
	"log"

	"ffxivcraftingplanner/xivdb"
)

type location struct {
	mapName    string
	regionName string
	position   *position
}

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

func getLocations(points []*xivdb.MapPoint) []*location {
	var locations []*location
	for _, p := range points {
		locations = append(locations, &location{
			mapName:    p.PlacenameName,
			regionName: p.RegionName,
			position: &position{
				x: p.Position.Position.X,
				y: p.Position.Position.Y,
			},
		})
	}

	return locations
}

type mapDataGetter interface {
	GetMapData() (*xivdb.MapData, error)
}

func getObtainData(name string, m mapDataGetter) *obtainData {
	var locations []*location
	data, err := m.GetMapData()
	if err == nil {
		name = fmt.Sprintf("%s (%s)", name, data.Maps[0].PlacenameName)
		locations = getLocations(data.Points)
	}

	return &obtainData{
		name:      name,
		locations: locations,
	}
}

func getEnemyObtainMethod(i *xivdb.Item) (*obtainMethod, error) {
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

		obtainDataArray = append(obtainDataArray, getObtainData(e.Name, e))
	}
	o.data = obtainDataArray
	return o, nil
}

func getShopsObtainMethod(i *xivdb.Item) (*obtainMethod, error) {
	o := &obtainMethod{
		obtainType: NPC,
	}
	var obtainDataArray []*obtainData
	for _, shop := range i.Shops {
		if npc, err := shop.GetNpc(); err == nil {
			obtainDataArray = append(obtainDataArray, getObtainData(npc.Name, npc))
		}
	}
	o.data = obtainDataArray
	return o, nil
}
