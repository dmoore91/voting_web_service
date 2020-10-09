package permission

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
	"voting_web_service/internal/app/users"
)

type UsersStruct struct {
	Users []users.User `json:"users"`
}

func AddPermission(writer http.ResponseWriter, request *http.Request) {
	// POST /permission/{permission}
	//
	// Add new permission to database
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	// 	- name: permission
	//   in: path
	//   description: permission to add
	//   type: string
	//   required: true
	// responses:
	//   '200':
	//     description: permission added
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

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "INSERT INTO Permissions(permission) " +
		"VALUES(?)"

	r, err := db.Exec(queryString, params["permission"])
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

func GetUsersForPermission(writer http.ResponseWriter, request *http.Request) {
	// GET /permission/{permission}
	//
	// Get users for permission
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	// 	- name: permission
	//   in: path
	//   description: Get users that have permission
	//   type: string
	//   required: true
	// responses:
	//   '200':
	//     description: Return list of users
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

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "SELECT user_id, username, email, first_name, last_name, party_id " +
		"FROM Users " +
		"INNER JOIN User_Permissions ON User_Permissions.user_id = Users.user_id " +
		"INNER JOIN Permissions ON Permissions.permission_id = User_Permissions.permission_id" +
		"WHERE permission=?"

	rows, err := db.Query(queryString, params["permission"])
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return
	}

	var userList []users.User

	defer rows.Close()

	for rows.Next() {
		var u = users.User{}
		err = rows.Scan(&u.UserId, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.Party)

		if err != nil {
			responses.GeneralSystemFailure(writer, "Get Failed")
			log.Error(err)
			return
		}

		userList = append(userList, u)
	}

	resp := UsersStruct{Users: userList}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(resp)
}
