package party

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

type party struct {
	Id    int    `json:"id"`
	Party string `json:"party"`
}

type partyList struct {
	Parties []party `json:"parties"`
}

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

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "SELECT party_id, party " +
		"FROM Party"

	rows, err := db.Query(queryString)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return
	}

	var parties []party

	defer rows.Close()

	for rows.Next() {
		var p = party{}
		err = rows.Scan(&p.Id, &p.Party)

		if err != nil {
			responses.GeneralSystemFailure(writer, "Get Failed")
			log.Error(err)
			return
		}

		parties = append(parties, p)
	}

	resp := partyList{Parties: parties}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(resp)
}

func LinkUserAndParty(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}

func UpdateUserParty(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}
