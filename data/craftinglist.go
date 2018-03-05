package data

import (
	"html/template"
	"net/http"
	"strconv"

	"ffxivcraftingplanner/crafttree"
	"ffxivcraftingplanner/helper"
	"ffxivcraftingplanner/xivdb"
)

func BuildCraftingList(w http.ResponseWriter, s string) error {
	id, _ := strconv.Atoi(s)
	r, err := xivdb.GetRecipe(id)
	if err != nil {
		return err
	}

	cl := crafttree.GetCraftingListFor([]int{r.Item.ID})

	t, err := template.New("webpage").Funcs(helper.FuncMap()).ParseFiles("web/html/precrafts-table.tmpl")
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, "precrafts-table", cl)
}

type Result struct {
	Icon     string
	ItemName string
	Location string
	Quantity int
	ID       int
}

func SearchRecipes(s string) ([]Result, error) {
	data, _ := xivdb.SearchRecipes(s)

	var results []Result
	for _, recipeResult := range data.Recipes.Results {
		results = append(results, Result{
			Icon:     recipeResult.Icon,
			ItemName: recipeResult.Name,
			Quantity: 1,
			Location: "Some Location",
			ID:       recipeResult.ID,
		})
	}

	return results, nil
}
