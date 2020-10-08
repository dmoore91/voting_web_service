package users

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

type User struct {
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Party     string `json:"party"`
}

type InputUser struct {
	Username       string `json:"username"`
	HashedPassword string `json:"password"`
	Email          string `json:"email"`
	FirstName      string `json:"first name"`
	LastName       string `json:"last name"`
	Party          string `json:"party"`
}

type UpdateUserStruct struct {
	Email     string `json:"email"`
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Party     string `json:"party"`
}

type LoginCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Permission struct {
	Permission string `json:"permission"`
}

type PermissionsStruct struct {
	Permissions []Permission `json:"permissions"`
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
	err = db.QueryRow(queryString, lc.Username, lc.Password).Scan(&exists)
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

	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/test")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "SELECT user_id, username, email, first_name, last_name, party_id " +
		"FROM Users " +
		"WHERE username=?"

	var user User
	err = db.QueryRow(queryString, params["username"]).Scan(&user.UserId, &user.Username, &user.Email, &user.FirstName,
		&user.LastName, &user.Party)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	decoder := json.NewDecoder(request.Body)
	var u UpdateUserStruct
	err := decoder.Decode(&u)
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

	queryString := "UPDATE Users " +
		"SET email=?, first_name=?, last_name=?, party_id=? " +
		"WHERE username=?"

	//TODO Need to change this to not be hardcoded
	r, err := db.Exec(queryString, u.Email, u.FirstName, u.LastName, 1, params["username"])
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	rowsAffected, err := r.RowsAffected()

	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	if rowsAffected == 0 {
		responses.GeneralSystemFailure(writer, "Query Failed")
		return
	}

	responses.GeneralSuccess(writer, "Success")
}

func AddUser(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var u InputUser
	err := decoder.Decode(&u)
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

	queryString := "INSERT INTO Users(username, hashed_password, email, first_name, last_name, party_id)  " +
		"VALUES(?, ?, ?, ?, ?, ?)"

	//TODO Need to change this to not be hardcoded
	r, err := db.Exec(queryString, u.Username, u.HashedPassword, u.Email, u.FirstName, u.FirstName, 1)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	rowsAffected, err := r.RowsAffected()

	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	if rowsAffected == 0 {
		responses.GeneralSystemFailure(writer, "Query Failed")
		return
	}

	responses.GeneralSuccess(writer, "Success")
}

func GetPermissionsForUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/test")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "SELECT permission " +
		"FROM Permissions " +
		"INNER JOIN User_Permissions ON User_Permissions.permission_id = Permissions.permission_id " +
		"INNER JOIN Users ON Users.user_id = User_Permissions.user_id " +
		"WHERE username=?"

	rows, err := db.Query(queryString, params["username"])
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return
	}

	var permissions []Permission

	defer rows.Close()

	for rows.Next() {
		var p = Permission{}
		err = rows.Scan(&p.Permission)

		if err != nil {
			responses.GeneralSystemFailure(writer, "Get Failed")
			log.Error(err)
			return
		}

		permissions = append(permissions, p)
	}

	resp := PermissionsStruct{Permissions: permissions}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(resp)
}

func AddPermissionForUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/test")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "SELECT user_id " +
		"FROM Users " +
		"WHERE username=?"

	var userId int
	err = db.QueryRow(queryString, params["username"]).Scan(&userId)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return
	}

	queryString = "SELECT permission_id " +
		"FROM Permissions " +
		"WHERE permission=?"

	var permissionID int
	err = db.QueryRow(queryString, params["permission"]).Scan(&permissionID)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return
	}

	queryString = "INSERT INTO User_Permissions(permission_id, user_id)  " +
		"VALUES(?, ?)"

	r, err := db.Exec(queryString, userId, permissionID)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	rowsAffected, err := r.RowsAffected()

	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	if rowsAffected == 0 {
		responses.GeneralSystemFailure(writer, "Query Failed")
		return
	}

	responses.GeneralSuccess(writer, "Success")
}

func RemovePermissionForUser(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}
