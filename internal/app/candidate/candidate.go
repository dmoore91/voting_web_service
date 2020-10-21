package candidate

import (
	"database/sql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

func CreateCandidate(writer http.ResponseWriter, request *http.Request) {
	// POST /candidate/{user}/{party}
	//
	// Endpoint to add candidate to database
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	- name: party
	//	  in: query
	//	  description: name of party candidate belongs to
	//	  type: string
	//	  required: true
	//	- name: user
	//	  in: query
	//	  description: name of user that candidate is
	//	  type: string
	//	  required: true
	// responses:
	//   '200':
	//     description: If candidate created
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

	queryString := "SELECT party_id " +
		"FROM Party " +
		"WHERE party=?"

	var partyID int
	err = db.QueryRow(queryString, params["party"]).Scan(&partyID)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	queryString = "SELECT user_id " +
		"FROM Users " +
		"WHERE username=?"

	var userID int
	err = db.QueryRow(queryString, params["user"]).Scan(&userID)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	queryString = "INSERT INTO Candidate(user_id, party_id) " +
		"VALUES(?, ?)"

	_, err = db.Exec(queryString, partyID, userID)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	responses.GeneralSuccess(writer, "Success")
}

func GetCandidates(writer http.ResponseWriter, request *http.Request) {
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

	responses.GeneralNotImplemented(writer)
}

func GetCandidate(writer http.ResponseWriter, request *http.Request) {
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

	responses.GeneralNotImplemented(writer)
}
