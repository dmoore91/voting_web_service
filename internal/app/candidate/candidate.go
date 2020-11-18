package candidate

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
	"voting_web_service/internal/app/session"
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
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
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
		db, err := sql.Open("mysql", "root:secret@tcp(mysql_db:3306)/voting")
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
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func GetCandidates(writer http.ResponseWriter, request *http.Request) {
	// GET /candidate
	//
	// Endpoint
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	 - name: session_info
	//	   in: body
	//	   description: session info
	//	   schema:
	//	     "$ref": "#/definitions/sessionInfo"
	//	   required: true
	// responses:
	//   '200':
	//     description: List of candidates
	//     schema:
	//       "$ref": "#/definitions/candidateList"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

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

		db, err := sql.Open("mysql", "root:secret@tcp(mysql_db:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "SELECT candidate_id, username, party " +
			"FROM Candidate " +
			"INNER JOIN Users ON Candidate.user_id = Users.user_id " +
			"INNER JOIN Party ON Party.party_id = Candidate.party_id"

		rows, err := db.Query(queryString)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Failed query")
			log.Error(err)
			return
		}

		var candidates []Candidate

		defer rows.Close()

		for rows.Next() {
			var c = Candidate{}
			err = rows.Scan(&c.CandidateID, &c.Username, &c.Party)

			if err != nil {
				responses.GeneralSystemFailure(writer, "Get Failed")
				log.Error(err)
				return
			}

			candidates = append(candidates, c)
		}

		resp := CandidateList{Candidates: candidates}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		_ = json.NewEncoder(writer).Encode(resp)
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func GetCandidate(writer http.ResponseWriter, request *http.Request) {
	// GET /candidate/{candidate_id}
	//
	// Endpoint to get specific candidate
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	- name: candidate_id
	//	  in: query
	//	  description: id for candidate
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
	//     description: candidate
	//     schema:
	//       "$ref": "#/definitions/candidate"
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

		db, err := sql.Open("mysql", "root:secret@tcp(mysql_db:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "SELECT candidate_id, username, party " +
			"FROM Candidate " +
			"INNER JOIN Users ON Candidate.user_id = Users.user_id " +
			"INNER JOIN Party ON Party.party_id = Candidate.party_id " +
			"WHERE candidate_id=?"

		var c Candidate

		err = db.QueryRow(queryString, params["candidate_id"]).Scan(&c.CandidateID, &c.Username, &c.Party)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Failed query")
			log.Error(err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		_ = json.NewEncoder(writer).Encode(c)
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}
