package gamedata

type Crystal struct {
	ID   int
	Name string
	Icon string
}

var crystalsMap map[int]*Crystal

func GetCrystal(id int) *Crystal {
	if val, ok := crystalsMap[id]; ok {
		return val
	}
	return nil
}

func initializeCrystalsMap() {
	crystalsMap = map[int]*Crystal{
		2: {
			ID:   2,
			Icon: "https://secure.xivdb.com/img/game_local/2/2.jpg",
			Name: "Fire Shard",
		},
		3: {
			ID:   3,
			Icon: "https://secure.xivdb.com/img/game_local/3/3.jpg",
			Name: "Ice Shard",
		},
		4: {
			ID:   4,
			Icon: "https://secure.xivdb.com/img/game_local/4/4.jpg",
			Name: "Wind Shard",
		},
		5: {
			ID:   5,
			Icon: "https://secure.xivdb.com/img/game_local/5/5.jpg",
			Name: "Earth Shard",
		},
		6: {
			ID:   6,
			Icon: "https://secure.xivdb.com/img/game_local/6/6.jpg",
			Name: "Lightning Shard",
		},
		7: {
			ID:   7,
			Icon: "https://secure.xivdb.com/img/game_local/7/7.jpg",
			Name: "Water Shard",
		},
		8: {
			ID:   8,
			Icon: "https://secure.xivdb.com/img/game_local/8/8.jpg",
			Name: "Fire Crystal",
		},
		9: {
			ID:   9,
			Icon: "https://secure.xivdb.com/img/game_local/9/9.jpg",
			Name: "Ice Crystal",
		},
		10: {
			ID:   10,
			Icon: "https://secure.xivdb.com/img/game_local/1/10.jpg",
			Name: "Wind Crystal",
		},
		11: {
			Icon: "https://secure.xivdb.com/img/game_local/1/11.jpg",
			Name: "Earth Crystal",
		},
		12: {
			ID:   12,
			Icon: "https://secure.xivdb.com/img/game_local/1/12.jpg",
			Name: "Lightning Crystal",
		},
		13: {
			ID:   13,
			Icon: "https://secure.xivdb.com/img/game_local/1/13.jpg",
			Name: "Water Crystal",
		},
		14: {
			ID:   14,
			Icon: "https://secure.xivdb.com/img/game_local/1/14.jpg",
			Name: "Fire Cluster",
		},
		15: {
			ID:   15,
			Icon: "https://secure.xivdb.com/img/game_local/1/15.jpg",
			Name: "Ice Cluster",
		},
		16: {
			ID:   16,
			Icon: "https://secure.xivdb.com/img/game_local/1/16.jpg",
			Name: "Wind Cluster",
		},
		17: {
			ID:   17,
			Icon: "https://secure.xivdb.com/img/game_local/1/17.jpg",
			Name: "Earth Cluster",
		},
		18: {
			ID:   18,
			Icon: "https://secure.xivdb.com/img/game_local/1/18.jpg",
			Name: "Lightning Cluster",
		},
		19: {
			ID:   19,
			Icon: "https://secure.xivdb.com/img/game_local/1/19.jpg",
			Name: "Water Cluster",
		},
	}
}
