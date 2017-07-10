package main

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"fmt"
	"io/ioutil"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user", postUser)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func postUser(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	var u User
	err = json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}

	createUser(u.Name)
}

func createUser(username string) int64 {
	driver := bolt.NewDriver()
	// @TODO: Put this in a config
	conn, err := driver.OpenNeo("bolt://api.omnomhub.dev:7687")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	stmt, err := conn.PrepareNeo("CREATE (node:User {name:{name}})")
	if err != nil {
		panic(err)
	}

	result, err := stmt.ExecNeo(map[string]interface{}{"name": username})
	if err != nil {
		panic(err)
	}

	numResult, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created the user: %s\n", username)

	// Closing the statement will also close the rows
	stmt.Close()

	return numResult
}
