package users

import (
	"context"
	"database/sql"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

type User struct {
	UserId         int    `json:"user_id"`
	Username       string `json:"username"`
	HashedPassword string `json:"password"`
	Email          string `json:"email"`
	FirstName      string `json:"first name"`
	LastName       string `json:"last name"`
	Party          string `json:"party"`
}

type Users struct {
	Users []User
}

type LoginCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var lc LoginCreds
	err := decoder.Decode(&lc)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/test")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "SELECT TRUE " +
		"FROM Users " +
		"WHERE username=? AND hashed_password=?"

	var exists bool
	err = db.QueryRowContext(context.Background(), queryString, lc.Username, lc.Password).Scan(&exists)
	if err != nil {
		if err.Error() != "no rows in result set" {
			responses.GeneralSuccess(writer, "User does not exist")
			return
		} else {
			responses.GeneralSystemFailure(writer, "Failed query")
			log.Error(err)
			return
		}
	}

	if exists {
		responses.GeneralSuccess(writer, "User Exists")
	} else {
		responses.GeneralSuccess(writer, "User does not exist")
	}
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}
