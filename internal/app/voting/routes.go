package voting

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/vote/{candidate_id}", VoteForCandidate).Methods("POST")
	router.HandleFunc(basePath+"/vote", GetVotesForCandidates).Methods("GET")
	router.HandleFunc(basePath+"/vote/candidate/{candidate_id}", GetVotesForCandidate).Methods("GET")
}
