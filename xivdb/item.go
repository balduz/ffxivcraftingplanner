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
