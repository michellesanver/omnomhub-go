package user

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func PostUser(rw http.ResponseWriter, req *http.Request) {
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

	CreateUser(u.Name)
}

func CreateUser(username string) int64 {
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

