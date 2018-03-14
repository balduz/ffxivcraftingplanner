package xivdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const (
	apiEnemyURL         = "https://api.xivdb.com/enemy/%d"
	apiItemURL          = "https://api.xivdb.com/item/%d"
	apiRecipeURL        = "https://api.xivdb.com/recipe/%d"
	apiSearchRecipesURL = "http://api.xivdb.com/search?string=%s&one=recipes"
)

type jsonToXivdbFunc func(*http.Response) (interface{}, error)

var itemsCache sync.Map
var enemiesCache sync.Map
var recipesCache sync.Map

// getFromXivdb retrieves from the XIVDB API the interface{} from the given url,
// parsing the JSON with the given parsing function.
func getFromXivdb(url string, parseFunc jsonToXivdbFunc) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("searchFromXivdb(_): returned error %v", err)
		return nil, err
	}

	result, err := parseFunc(resp)
	if err != nil {
		log.Printf("searchFromXivdb(_): could not decode json, error %v", err)
		return nil, err
	}

	return result, nil
}

// getFromXivdb retrieves from the Xivdb API the interface{} with the given id,
// trying to catch it from cache first.
func getFromXivdbOrCache(c sync.Map, url string, id int, parseFunc jsonToXivdbFunc) (interface{}, error) {
	if data, ok := c.Load(id); ok {
		return data, nil
	}

	result, err := getFromXivdb(url, parseFunc)
	if err != nil {
		return nil, err
	}

	c.Store(id, result)
	return result, nil
}

// GetItem returns an Item given the ID, retrieved from XIVDB.
func GetItem(id int) (*Item, error) {
	url := fmt.Sprintf(apiItemURL, id)
	res, err := getFromXivdbOrCache(itemsCache, url, id, func(r *http.Response) (interface{}, error) {
		var data Item
		err := json.NewDecoder(r.Body).Decode(&data)
		return &data, err
	})
	if err != nil {
		log.Printf("GetItem(%d): getting from Xivdb returned %s", id, err)
		return nil, err
	}

	return res.(*Item), nil
}

// GetEnemy returns an Enemy given the ID, retrieved from XIVDB.
func GetEnemy(id int) (*Enemy, error) {
	url := fmt.Sprintf(apiEnemyURL, id)
	res, err := getFromXivdbOrCache(enemiesCache, url, id, func(r *http.Response) (interface{}, error) {
		var data Enemy
		err := json.NewDecoder(r.Body).Decode(&data)
		return &data, err
	})
	if err != nil {
		log.Printf("GetEnemy(%d): getting from Xivdb returned %s", id, err)
		return nil, err
	}

	return res.(*Enemy), nil
}

// GetRecipe returns a Recipe given the ID, retrieved from XIVDB.
func GetRecipe(id int) (*Recipe, error) {
	url := fmt.Sprintf(apiRecipeURL, id)
	res, err := getFromXivdbOrCache(recipesCache, url, id, func(r *http.Response) (interface{}, error) {
		var data Recipe
		err := json.NewDecoder(r.Body).Decode(&data)
		return &data, err
	})
	if err != nil {
		log.Printf("GetRecipe(%d): getting from Xivdb returned %s", id, err)
		return nil, err
	}

	return res.(*Recipe), nil
}

// SearchRecipes returns all the recipes whose name match the given string.
func SearchRecipes(s string) (*RecipesSearch, error) {
	url := fmt.Sprintf(apiSearchRecipesURL, s)
	res, err := getFromXivdb(url, func(r *http.Response) (interface{}, error) {
		var data RecipesSearch
		err := json.NewDecoder(r.Body).Decode(&data)
		return &data, err
	})
	if err != nil {
		log.Printf("SearchRecipes(%s): getting from Xivdb returned %s", s, err)
		return nil, err
	}

	return res.(*RecipesSearch), nil
}
