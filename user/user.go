package user

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/twinj/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string `json:"name"`

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

	// Check if ID exists, if not...

	// Create a uuid and create user.
	u.Id = uuid.NewV4();
	CreateUser(u)

	// If it does, update user.
}

func CreateUser(u User) int64 {
	driver := bolt.NewDriver()
	// @TODO: Put this in a config
	conn, err := driver.OpenNeo("bolt://api.omnomhub.dev:7687")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	stmt, err := conn.PrepareNeo("CREATE (node:User {username:{username}, id:{id}})")
	if err != nil {
		panic(err)
	}

	result, err := stmt.ExecNeo(map[string]interface{}{
		"username": u.Username,
		"id": u.Id.String(),
	})

	if err != nil {
		panic(err)
	}

	numResult, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created the user: %s\n", u.Username)

	// Closing the statement will also close the rows
	stmt.Close()

	return numResult
}
