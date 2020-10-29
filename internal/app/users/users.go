package users

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"voting_web_service/internal/app/responses"
	"voting_web_service/internal/app/session"
)

type User struct {
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Party     string `json:"party"`
}

// swagger:model newUserInfo
type InputUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Party     string `json:"party"`
	SecretKey string `json:"secret_key"`
}

// swagger:model updateUserInfo
type UpdateUserStruct struct {
	Email        string              `json:"email"`
	FirstName    string              `json:"first name"`
	LastName     string              `json:"last name"`
	Party        string              `json:"party"`
	SessionCreds session.SessionInfo `json:"session"`
}

// swagger:model loginCreds
type LoginCreds struct {
	Username     string              `json:"username"`
	Password     string              `json:"password"`
	SessionCreds session.SessionInfo `json:"session"`
}

type Permission struct {
	Permission string `json:"permission"`
}

type PermissionsStruct struct {
	Permissions []Permission `json:"permissions"`
}

type SecretKeyStruct struct {
	SecretKey string `json:"secret_key"`
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func isCorrectPassword(writer http.ResponseWriter, byteHash []byte, pwd []byte) bool {

	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		responses.GeneralNoContent(writer, "User does not exist")
		return false
	}
	return true
}

func getPartyIdForParty(writer http.ResponseWriter, party string) int {

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return -1
	}

	defer db.Close()

	queryString := "SELECT party_id " +
		"FROM Party " +
		"WHERE party=?"

	var partyID int
	err = db.QueryRow(queryString, party).Scan(&partyID)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return -1
	}

	return partyID
}

func LoginUser(writer http.ResponseWriter, request *http.Request) {
	// POST /user/login
	//
	// Endpoint
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	 - name: login_info
	//	   in: body
	//	   description: username and password to login
	//	   schema:
	//	     "$ref": "#/definitions/loginCreds"
	//	   required: true
	// responses:
	//   '200':
	//     description: if user exists
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '204':
	//     description: if user doesn't exists
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	decoder := json.NewDecoder(request.Body)
	var lc LoginCreds
	err := decoder.Decode(&lc)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(lc.SessionCreds.Username, lc.SessionCreds.SessionID)

	if valid {

		db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "SELECT hashed_password, secret_key " +
			"FROM Users " +
			"WHERE username=?"

		var hashedPass string
		var secretKey string

		err = db.QueryRow(queryString, lc.Username).Scan(&hashedPass, &secretKey)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Failed query")
			log.Error(err)
			return
		}

		isCorrect := isCorrectPassword(writer, []byte(hashedPass), []byte(lc.Password))

		if isCorrect {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(200)
			_ = json.NewEncoder(writer).Encode(SecretKeyStruct{SecretKey: secretKey})
		} else {
			responses.GeneralNoContent(writer, "User does not exist")
		}
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	// GET /user/{username}
	//
	// Endpoint
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	// - name: username
	//   in: path
	//   description: username for user
	//   type: string
	//   required: true
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
	// responses:
	//   '200':
	//     description: if user is logged in
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	params := mux.Vars(request)

	decoder := json.NewDecoder(request.Body)
	var si session.SessionInfo
	err := decoder.Decode(&si)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(si.Username, si.SessionID)

	if valid {
		db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	// PUT /user/{username}
	//
	// Endpoint to update user info. Cannot change username, password or user_id
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	// 	- name: username
	//   in: path
	//   description: username for user
	//   type: string
	//   required: true
	//	- name: user
	//	  in: body
	//	  description: info to update user
	//	  schema:
	//	    "$ref": "#/definitions/updateUserInfo"
	//	  required: true
	// responses:
	//   '200':
	//     description: user updated
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	params := mux.Vars(request)

	decoder := json.NewDecoder(request.Body)
	var u UpdateUserStruct
	err := decoder.Decode(&u)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(u.SessionCreds.Username, u.SessionCreds.SessionID)

	if valid {

		db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "UPDATE Users " +
			"SET email=?, first_name=?, last_name=?, party_id=? " +
			"WHERE username=?"

		partyID := getPartyIdForParty(writer, u.Party)

		//If it's failed we've already returned an error message so all we need to do is exit this function
		if partyID != -1 {

			r, err := db.Exec(queryString, u.Email, u.FirstName, u.LastName, partyID, params["username"])
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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	// POST /user
	//
	// Endpoint to add user
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	- name: user
	//	  in: body
	//	  description: new user info
	//	  schema:
	//	    "$ref": "#/definitions/newUserInfo"
	//	  required: true
	// responses:
	//   '200':
	//     description: user added
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	decoder := json.NewDecoder(request.Body)
	var u InputUser
	err := decoder.Decode(&u)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "INSERT INTO Users(username, hashed_password, email, first_name, last_name, party_id, secret_key)  " +
		"VALUES(?, ?, ?, ?, ?, ?, ?)"

	partyID := getPartyIdForParty(writer, u.Party)

	//If it's failed we've already returned an error message so all we need to do is exit this function
	if partyID != -1 {
		r, err := db.Exec(queryString, u.Username, hashAndSalt([]byte(u.Password)), u.Email, u.FirstName,
			u.LastName, partyID, u.SecretKey)
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
}

func GetPermissionsForUser(writer http.ResponseWriter, request *http.Request) {
	// GET /user/permission/{username}
	//
	// Gets list of permissions for user
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	// 	- name: username
	//   in: path
	//   description: username for user we want permissions for
	//   type: string
	//   required: true
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
	// responses:
	//   '200':
	//     description: permission we got
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	params := mux.Vars(request)

	decoder := json.NewDecoder(request.Body)
	var si session.SessionInfo
	err := decoder.Decode(&si)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(si.Username, si.SessionID)

	if valid {

		db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func AddPermissionForUser(writer http.ResponseWriter, request *http.Request) {
	// POST /user/{username}/{permission}
	//
	// Add permission represented by permission and link to username
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	// 	- name: username
	//   in: path
	//   description: username to add permissions for
	//   type: string
	//   required: true
	// 	- name: permission
	//   in: path
	//   description: permission to link with user
	//   type: string
	//   required: true
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
	// responses:
	//   '200':
	//     description: permission we got
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	params := mux.Vars(request)

	decoder := json.NewDecoder(request.Body)
	var si session.SessionInfo
	err := decoder.Decode(&si)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(si.Username, si.SessionID)

	if valid {

		db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func RemovePermissionForUser(writer http.ResponseWriter, request *http.Request) {
	// DELETE /user/{username}/{permission}
	//
	// Delete permission represented by permission from username
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	// 	- name: username
	//   in: path
	//   description: username to delete permissions for
	//   type: string
	//   required: true
	// 	- name: permission
	//   in: path
	//   description: permission to delete
	//   type: string
	//   required: true
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
	// responses:
	//   '200':
	//     description: delete permission
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	params := mux.Vars(request)

	decoder := json.NewDecoder(request.Body)
	var si session.SessionInfo
	err := decoder.Decode(&si)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(si.Username, si.SessionID)

	if valid {

		db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
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

		queryString = "DELETE FROM User_Permissions " +
			"WHERE user_id=? AND permission_id=?"

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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}
