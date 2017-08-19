package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/michellesanver/omnomhub-go/user"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user", user.PostUser)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
