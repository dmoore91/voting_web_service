package voting

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

// swagger:model votes
type VotesStruct struct {
	Votes int `json:"votes"`
}

// swagger:model votesForCandidates
type VotesForCandidate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Party     string `json:"party"`
	Votes     int    `json:"votes"`
}

// swagger:model votesForCandidatesList
type VotesForCandidateList struct {
	Candidates []VotesForCandidate `json:"candidates"`
}

func VoteForCandidate(writer http.ResponseWriter, request *http.Request) {
	// POST /vote/{candidate_id}
	//
	// Endpoint to vote for candidate
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	- name: candidate_id
	//	  in: query
	//	  description: id for candidate that's being voted for
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
	//     description: Vote counted for candidate
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
	var lc users.LoginCreds
	err := decoder.Decode(&lc)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(lc.SessionCreds.Username, lc.SessionCreds.SessionID)

	if valid {
		params := mux.Vars(request)

		db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "SELECT user_id " +
			"FROM Users " +
			"WHERE username=?"

		var userID int
		err = db.QueryRow(queryString, params["username"]).Scan(&userID)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Query Failed")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString = "UPDATE Candidate " +
			"SET votes = votes + 1 " +
			"WHERE user_id=?"

		_, err = db.Exec(queryString, userID)
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

func GetVotesForCandidate(writer http.ResponseWriter, request *http.Request) {
	// GET /vote/candidate/{candidate_id}
	//
	// Endpoint to get votes for candidate
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	- name: candidate_id
	//	  in: query
	//	  description: id for candidate to get votes for
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
	//     description: Number of votes for candidate
	//     schema:
	//       "$ref": "#/definitions/votes"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/genlocalhost:8880/voting/vote/1eralResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	decoder := json.NewDecoder(request.Body)
	var lc users.LoginCreds
	err := decoder.Decode(&lc)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(lc.SessionCreds.Username, lc.SessionCreds.SessionID)

	if valid {
		params := mux.Vars(request)

		db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "SELECT user_id " +
			"FROM Users " +
			"WHERE username=?"

		var userID int
		err = db.QueryRow(queryString, params["username"]).Scan(&userID)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Query Failed")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString = "SELECT votes " +
			"FROM Candidate " +
			"WHERE user_id=?"

		var votes int
		err = db.QueryRow(queryString, userID).Scan(&votes)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Query Failed")
			log.Error(err)
			return
		}

		resp := VotesStruct{Votes: votes}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		_ = json.NewEncoder(writer).Encode(resp)
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}

func GetVotesForCandidates(writer http.ResponseWriter, request *http.Request) {
	// POST /vote
	//
	// Endpoint to vote for all candidates
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
	//     description: Got votes for candidates
	//     schema:
	//       "$ref": "#/definitions/votesForCandidatesList"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server error
	//     schema:
	//       "$ref": "#/definitions/generalResponse"

	decoder := json.NewDecoder(request.Body)
	var lc users.LoginCreds
	err := decoder.Decode(&lc)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	valid := session.CheckSessionID(lc.SessionCreds.Username, lc.SessionCreds.SessionID)

	if valid {
		db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
		if err != nil {
			responses.GeneralSystemFailure(writer, "Cannot connect to db")
			log.Error(err)
			return
		}

		defer db.Close()

		queryString := "SELECT first_name, last_name, party, votes " +
			"FROM Candidate " +
			"INNER JOIN Users ON Users.user_id = Candidate.user_id " +
			"INNER JOIN Party ON Party.party_id = Candidate.party_id"

		rows, err := db.Query(queryString)
		if err != nil {
			responses.GeneralSystemFailure(writer, "Failed query")
			log.Error(err)
			return
		}

		var candidates []VotesForCandidate

		defer rows.Close()

		for rows.Next() {
			var c = VotesForCandidate{}
			err = rows.Scan(&c.FirstName, &c.LastName, &c.Party, &c.Votes)

			if err != nil {
				responses.GeneralSystemFailure(writer, "Get Failed")
				log.Error(err)
				return
			}

			candidates = append(candidates, c)
		}

		resp := VotesForCandidateList{Candidates: candidates}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		_ = json.NewEncoder(writer).Encode(resp)
	} else {
		responses.GeneralBadRequest(writer, "Bad Session Token")
	}
}
