package recipe

import (
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type Ingredient struct {
	Name string `json:"name"` // E.G. 'Onions'
	State string `json:"state"` // E.G. 'chopped'
	Unit string `json:"unit"` // E.G. 'grams'
	Amount int64 `json:"amount"` // E.G. 20
}

type Recipe struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Directions string `json:"directions"`
	ImageUrl string `json:"imageUrl"`
	CookingTime string `json:"cookingTime"`
	PreparationTime string `json:"preparationTime"`
	Servings int64 `json:"servings"`
	Ingredients[] Ingredient `json:"ingredients"`
}

func PostRecipe(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	var r Recipe
	err = json.Unmarshal(body, &r)

	if err != nil {
		panic(err)
	}

	CreateRecipe(r)
}

func CreateRecipe(r Recipe) int64 {
	// todo: Implement this to create recipe
	return 0
}
