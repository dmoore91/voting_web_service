package party

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
	"voting_web_service/internal/app/session"
	"voting_web_service/internal/app/users"
)

// swagger:model party
type party struct {
	Id    int    `json:"id"`
	Party string `json:"party"`
}

// swagger:model partyList
type partyList struct {
	Parties []party `json:"parties"`
}

func CreateParty(writer http.ResponseWriter, request *http.Request) {
	// POST /party/{party}
	//
	// Endpoint to add party to database
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	- name: party
	//	  in: query
	//	  description: name of party to add
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
	//     description: party added
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

	//decoder := json.NewDecoder(request.Body)
	//	var lc users.LoginCreds
	//	err := decoder.Decode(&lc)
	//	if err != nil {
	//		responses.GeneralBadRequest(writer, "Decode Failed")
	//		log.Error(err)
	//		return
	//	}

	//valid := session.CheckSessionID(lc.SessionCreds.Username, lc.SessionCreds.SessionID)
	valid := true

	if valid {
		db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func GetParties(writer http.ResponseWriter, request *http.Request) {
	// GET /party
	//
	// Endpoint to get all parties in database
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
	// responses:
	//   '200':
	//     description: parties
	//     schema:
	//       "$ref": "#/definitions/partyList"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:

	//decoder := json.NewDecoder(request.Body)
	//	var lc users.LoginCreds
	//	err := decoder.Decode(&lc)
	//	if err != nil {
	//		responses.GeneralBadRequest(writer, "Decode Failed")
	//		log.Error(err)
	//		return
	//	}

	//valid := session.CheckSessionID(lc.SessionCreds.Username, lc.SessionCreds.SessionID)
	valid := true

	if valid {
		db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func LinkUserAndParty(writer http.ResponseWriter, request *http.Request) {
	// POST /party/link/{user}/{party}
	//
	// Endpoint to link user and party
	//
	// ---
	// produces:
	// - application/json
	//	- name: party
	//	  in: query
	//	  description: name of party to link
	//	  type: string
	//	  required: true
	//	- name: user
	//	  in: query
	//	  description: name of user to link
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
	//     description: user and party linked
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:

	//decoder := json.NewDecoder(request.Body)
	//	var lc users.LoginCreds
	//	err := decoder.Decode(&lc)
	//	if err != nil {
	//		responses.GeneralBadRequest(writer, "Decode Failed")
	//		log.Error(err)
	//		return
	//	}

	//valid := session.CheckSessionID(lc.SessionCreds.Username, lc.SessionCreds.SessionID)
	valid := true

	if valid {

		params := mux.Vars(request)

		db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
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

		queryString = "Update Users " +
			"SET party_id=? " +
			"WHERE user_id=?"

		_, err = db.Exec(queryString, partyID, userID)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Query Failed")
			log.Error(err)
			return
		}

		responses.GeneralSuccess(writer, "Success")
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}
