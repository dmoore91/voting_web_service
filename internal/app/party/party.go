package party

import (
	"database/sql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

func CreateParty(writer http.ResponseWriter, request *http.Request) {
	// POST /party
	//
	// Endpoint to add party to database
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

	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "INSERT INTO Party(party)  " +
		"VALUES(?)"

	r, err := db.Exec(queryString, params["party"])
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

func GetParties(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}

func LinkUserAndParty(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}

func UpdateUserParty(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}
