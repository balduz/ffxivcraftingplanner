package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"ffxivcraftingplanner/data"
)

func setUpLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	setUpLogger()

	mux := http.NewServeMux()

	mux.Handle("/web/assets/", http.StripPrefix("/web/assets/", http.FileServer(http.Dir("web/assets"))))
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/recipe/", recipeHandler)

	http.ListenAndServe(getPort(), mux)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("name").ParseFiles("web/html/base.html", "web/html/search-results.tmpl", "web/html/ingredient-row.tmpl")
	t.ExecuteTemplate(w, "base", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	searchValue := r.FormValue("search")
	log.Printf("Searching for Recipes with: %s\n", searchValue)
	results, err := data.SearchRecipes(searchValue)

	if err != nil {
		// TODO(cbalduz): show some error.
	}

	t, err := template.New("webpage").ParseFiles("web/html/search-results.tmpl")
	if err != nil {
		log.Fatalf("Error in searchHandler: %s", err)
	}

	err = t.ExecuteTemplate(w, "search-results", results)
	if err != nil {
		log.Fatalf("Error in searchHandler: %s", err)
	}
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/recipe/"):]
	fmt.Printf("Getting Recipe for Id: %s\n", id)
	err := data.BuildCraftingList(w, id)
	if err != nil {
		log.Fatalf("Error in searchHandler: %s", err)
	}

}
