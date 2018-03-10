package xivdb

import (
	"encoding/json"
	"log"
)

// GetGatheringNodes returns a slice with all the gathering nodes available
// for the Item.
func (i *Item) GetGatheringNodes() ([]*Node, error) {
	var nodesMap map[string]Node
	err := json.Unmarshal(i.Gathering[0].Nodes, &nodesMap)
	if err != nil {
		log.Printf("GetNodesForItem(_): id %d returned error %s", i.ID, err)
		return nil, err
	}
	var nodes []*Node
	for _, v := range nodesMap {
		nodes = append(nodes, &v)
	}

	return nodes, nil
}

// GetNpc returns the NPC for the given shop.
func (s *Shop) GetNpc() (*Npc, error) {
	var npc Npc
	err := json.Unmarshal(s.Npc, &npc)
	if err != nil {
		log.Printf("GetNpc(_): %s", err)
		return nil, err
	}

	return &npc, nil
}

func (n *Npc) GetMapData() (*MapData, error) {
	return getMapData(n.MapData)
}

func (e *Enemy) GetMapData() (*MapData, error) {
	return getMapData(e.MapData)
}

func getMapData(raw json.RawMessage) (*MapData, error) {
	var data *MapData
	err := json.Unmarshal(raw, &data)
	if err != nil {
		log.Printf("npc.GetMapData(): %s", err)
		return nil, err
	}

	return data, nil
}
