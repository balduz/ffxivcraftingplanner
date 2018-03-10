package helper

import (
	"html/template"

	"ffxivcraftingplanner/crafttree"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"ObtainIcon": func(o crafttree.ObtainType) string {
			switch o {
			case crafttree.ENEMY:
				return "/web/assets/icons/enemy_drop.png"
			case crafttree.NPC:
				return "/web/assets/icons/gil.png"
			}
			return "a"
		},
	}
}
