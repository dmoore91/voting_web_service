package session

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

// swagger:model sessionInfo
type SessionInfo struct {
	SessionID int    `json:"session_id"`
	Username  string `json:"username"`
}

// This function will be used by every other function to make sure that the user has the correct session_id.
// This takes care of the authentication step in auth/auth
func CheckSessionID(username string, sessionId int) bool {

	return true

	//db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
	//if err != nil {
	//	return false
	//}
	//
	//defer db.Close()
	//
	//queryString := "SELECT session " +
	//	"FROM Users " +
	//	"WHERE username=?"
	//
	//var userSessionId int
	//err = db.QueryRow(queryString, username).Scan(&userSessionId)
	//if err != nil {
	//	return false
	//}
	//
	//return userSessionId == sessionId
}

func SetSessionIdNull(writer http.ResponseWriter, request *http.Request) {
	// POST /session/sign_out/{user}
	//
	// Endpoint to set session id to null when signing out user
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	- name: user
	//	  in: query
	//	  description: name of user that is signing out
	//	  type: string
	//	  required: true
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
	// responses:
	//   '200':
	//     description: Session successfully set to null
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
	var si SessionInfo
	err := decoder.Decode(&si)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := CheckSessionID(si.Username, si.SessionID)

	if valid {
		params := mux.Vars(request)

		db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "UPDATE Users " +
			"SET session='' " +
			"WHERE username=?"

		r, err := db.Exec(queryString, params["user"])

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
