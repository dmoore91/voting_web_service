package users

import (
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

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
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
