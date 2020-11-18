package voting

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/vote/{username}", VoteForCandidate).Methods("POST")
	router.HandleFunc(basePath+"/vote", GetVotesForCandidates).Methods("GET")
	router.HandleFunc(basePath+"/vote/candidate/{username}", GetVotesForCandidate).Methods("GET")
}
