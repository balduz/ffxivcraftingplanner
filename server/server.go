package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"ffxivcraftingplanner/crafttree"
	"ffxivcraftingplanner/helper"
	"ffxivcraftingplanner/xivdb"
)

const webHTMLDirPath = "web/html"

func getPath(s string) string {
	return fmt.Sprintf("%s/%s", webHTMLDirPath, s)
}

func getPaths(s []string) []string {
	paths := make([]string, len(s))
	for i, p := range s {
		paths[i] = getPath(p)
	}
	return paths
}

type server struct {
	mux *http.ServeMux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s", r.URL)
	s.mux.ServeHTTP(w, r)
}

func StartServer(port string) error {
	mux := http.NewServeMux()

	mux.Handle("/web/assets/", http.StripPrefix("/web/assets/", http.FileServer(http.Dir("web/assets"))))
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/craftinglist", craftingListHandler)
	mux.HandleFunc("/recipe/", recipeHandler)

	return http.ListenAndServe(port, &server{mux})
}

type templateData struct {
	files []string
	name  string
	data  interface{}
}

type templateDataFunc func(r *http.Request, s *Session) (templateData, error)

func handle(w http.ResponseWriter, r *http.Request, f templateDataFunc) {
	s := GetSession(r)
	if s == nil {
		s = AddSessionCookie(w)
	}
	td, err := f(r, s)
	if err != nil {
		log.Fatalf("error with template %s. %s", td.name, err)
	}
	t, err := template.New("").Funcs(helper.FuncMap()).ParseFiles(getPaths(td.files)...)
	if err != nil {
		log.Fatalf("error with template %s. %s", td.name, err)
	}
	err = t.ExecuteTemplate(w, td.name, td.data)
	if err != nil {
		log.Fatalf("error with template %s. %s", td.name, err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	f := func(r *http.Request, s *Session) (templateData, error) {
		return templateData{
			files: []string{"base.html", "search/search-results.tmpl"},
			name:  "base",
			data:  nil,
		}, nil
	}
	handle(w, r, f)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	f := func(r *http.Request, s *Session) (templateData, error) {
		searchValue := r.FormValue("search")
		results, err := searchRecipes(searchValue)
		return templateData{
			files: []string{"search/search-results.tmpl"},
			name:  "search",
			data:  results,
		}, err
	}
	handle(w, r, f)
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	f := func(r *http.Request, s *Session) (templateData, error) {
		strID := r.URL.Path[len("/recipe/"):]
		id, _ := strconv.Atoi(strID)
		recipe, err := xivdb.GetRecipe(id)

		s.CraftingList = append(s.CraftingList, recipe.Item.ID)

		cl := crafttree.GetCraftingListFor(s.CraftingList)
		return templateData{
			files: []string{"craftlist/craftinglist.tmpl", "craftlist/other-obtain-methods.tmpl"},
			name:  "craftinglist",
			data:  cl,
		}, err
	}
	handle(w, r, f)
}

func craftingListHandler(w http.ResponseWriter, r *http.Request) {
	f := func(r *http.Request, s *Session) (templateData, error) {
		cl := crafttree.GetCraftingListFor(s.CraftingList)
		return templateData{
			files: []string{"craftlist/craftinglist.tmpl", "craftlist/other-obtain-methods.tmpl"},
			name:  "craftinglist",
			data:  cl,
		}, nil
	}
	handle(w, r, f)
}

type Result struct {
	Icon     string
	ItemName string
	Location string
	Quantity int
	ID       int
}

func searchRecipes(s string) ([]Result, error) {
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
