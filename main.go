package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/michellesanver/omnomhub-go/user"
	"github.com/michellesanver/omnomhub-go/recipe"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user", user.SaveUser)
	http.HandleFunc("/recipe", recipe.PostRecipe)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
