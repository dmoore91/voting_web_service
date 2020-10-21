package voting

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

// swagger:model votes
type VotesStruct struct {
	Votes int `json:"votes"`
}

func VoteForCandidate(writer http.ResponseWriter, request *http.Request) {
	// POST /voting/{candidate_id}
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
	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "UPDATE Candidate " +
		"SET votes = votes + 1 " +
		"WHERE candidate_id=?"

	_, err = db.Exec(queryString, params["candidate_id"])
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	responses.GeneralSuccess(writer, "Success")
}

func GetVotesForCandidate(writer http.ResponseWriter, request *http.Request) {
	// GET /voting/candidate/{candidate_id}
	//
	// Endpoint to vote for candidate
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
	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "SELECT votes " +
		"FROM Candidate " +
		"WHERE candidate_id=?"

	var votes int
	err = db.QueryRow(queryString, params["candidate_id"]).Scan(&votes)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	resp := VotesStruct{Votes: votes}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(resp)
}

func GetVotesForCandidates(writer http.ResponseWriter, request *http.Request) {
	// POST /voting/{candidate_id}
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
	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:secret@tcp(0.0.0.0:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "UPDATE Votes " +
		"SET votes = votes + 1 " +
		"WHERE candidate_id=?"

	_, err = db.Exec(queryString, params["candidate_id"])
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	responses.GeneralSuccess(writer, "Success")
}
