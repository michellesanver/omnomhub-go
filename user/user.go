package user

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/twinj/uuid"
	"encoding/base64"
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/errors"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/messages"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string `json:"email"`
	DisplayName string `json:"display_name"`
	Password string `json:"password"`
}

type appError struct {
	Error   error
	Message string
	Code    int
}

func SaveUser(rw http.ResponseWriter, req *http.Request) {
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

	if req.Method == "PUT" {
		if u.Id != nil {
			http.Error(rw, "You can not enter an ID when creating a user. You probably want to use POST to update a user with an existing ID.", 400)
		}

		u.Id = uuid.NewV4();

		err = CreateUser(u, rw)

		if err != nil {

		}
	}

	if req.Method == "POST" {
		// @TODO: Update user
	}
}

func CreateUser(u User, rw http.ResponseWriter) (error) {
	driver := bolt.NewDriver()
	// @TODO: Put this in a config
	conn, err := driver.OpenNeo("bolt://api.omnomhub.dev:7687")
	defer conn.Close()

	if err != nil {
		return err;
	}

	stmt, err := conn.PrepareNeo("CREATE (node:User {id:{id}, display_name:{display_name}, email:{email}, password:{password}})")

	if err != nil {
		return err;
	}

	hashedPassword, err := hashPassword(u.Password)

	if err != nil {
		return err;
	}

	result, err := stmt.ExecNeo(map[string]interface{}{
		"id": u.Id.String(),
		"display_name": u.DisplayName,
		"email": u.Email,
		"password": hashedPassword,
	})

	if err != nil {
		log.Info(err.Error())

		switch v := err.(type) {
		case *errors.Error:
			err = handleDbError(v.Inner(), rw)
			return 0, err
		default:
			http.Error(rw, err.Error(), 500)
			return 0, err
		}
	}

	numResult, err := result.RowsAffected()

	if numResult == 0 {
		return error("No user was created.")
	}

	if err != nil {
		return err;
	}

	fmt.Printf("Created the user: %s\n", u.DisplayName)

	// Closing the statement will also close the rows
	stmt.Close()

	return err
}

func handleDbError(err error, rw http.ResponseWriter) (error) {

	switch v := err.(type) {
	case messages.FailureMessage:
		c := v.Metadata["code"].(string)
		m := v.Metadata["message"].(string)

		if c == "Neo.ClientError.Schema.ConstraintValidationFailed" {
			http.Error(rw, m, 409)
		}
	default:
		http.Error(rw, err.Error(), 500)
	}

	return err;
}

func hashPassword(p string) (string, error) {
	size := 32
	rb := make([]byte, size)
	_, err := rand.Read(rb)

	generatedSalt := base64.URLEncoding.EncodeToString(rb)

	saltAndPassword := append([]byte(p), []byte(generatedSalt)...)
	hashedPassword, err := bcrypt.GenerateFromPassword(saltAndPassword, 10)

	return string(hashedPassword), err
}

func checkServerError(rw http.ResponseWriter, err error, eMsg string) {
	if err != nil {
		http.Error(rw, eMsg, 500)
		panic(err)
	}
}
