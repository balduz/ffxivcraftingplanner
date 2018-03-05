package crafttree

type crystal struct {
	id   int
	name string
	icon string
}

var crystalsMap map[int]*crystal

func getCrystal(id int) *crystal {
	if val, ok := crystalsMap[id]; ok {
		return val
	}
	return nil
}

func initializeCrystalsMap() {
	crystalsMap = map[int]*crystal{
		2: &crystal{
			id:   2,
			icon: "https://secure.xivdb.com/img/game_local/2/2.jpg",
			name: "Fire Shard",
		},
		3: {
			id:   3,
			icon: "https://secure.xivdb.com/img/game_local/3/3.jpg",
			name: "Ice Shard",
		},
		4: {
			id:   4,
			icon: "https://secure.xivdb.com/img/game_local/4/4.jpg",
			name: "Wind Shard",
		},
		5: {
			id:   5,
			icon: "https://secure.xivdb.com/img/game_local/5/5.jpg",
			name: "Earth Shard",
		},
		6: {
			id:   6,
			icon: "https://secure.xivdb.com/img/game_local/6/6.jpg",
			name: "Lightning Shard",
		},
		7: {
			id:   7,
			icon: "https://secure.xivdb.com/img/game_local/7/7.jpg",
			name: "Water Shard",
		},
		8: {
			id:   8,
			icon: "https://secure.xivdb.com/img/game_local/8/8.jpg",
			name: "Fire Crystal",
		},
		9: {
			id:   9,
			icon: "https://secure.xivdb.com/img/game_local/9/9.jpg",
			name: "Ice Crystal",
		},
		10: {
			id:   10,
			icon: "https://secure.xivdb.com/img/game_local/1/10.jpg",
			name: "Wind Crystal",
		},
		11: {
			icon: "https://secure.xivdb.com/img/game_local/1/11.jpg",
			name: "Earth Crystal",
		},
		12: {
			id:   12,
			icon: "https://secure.xivdb.com/img/game_local/1/12.jpg",
			name: "Lightning Crystal",
		},
		13: {
			id:   13,
			icon: "https://secure.xivdb.com/img/game_local/1/13.jpg",
			name: "Water Crystal",
		},
		14: {
			id:   14,
			icon: "https://secure.xivdb.com/img/game_local/1/14.jpg",
			name: "Fire Cluster",
		},
		15: {
			id:   15,
			icon: "https://secure.xivdb.com/img/game_local/1/15.jpg",
			name: "Ice Cluster",
		},
		16: {
			id:   16,
			icon: "https://secure.xivdb.com/img/game_local/1/16.jpg",
			name: "Wind Cluster",
		},
		17: {
			id:   17,
			icon: "https://secure.xivdb.com/img/game_local/1/17.jpg",
			name: "Earth Cluster",
		},
		18: {
			id:   18,
			icon: "https://secure.xivdb.com/img/game_local/1/18.jpg",
			name: "Lightning Cluster",
		},
		19: {
			id:   19,
			icon: "https://secure.xivdb.com/img/game_local/1/19.jpg",
			name: "Water Cluster",
		},
	}
}
