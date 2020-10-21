package candidate

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/candidate/{user}/{party}", CreateCandidate).Methods("POST")
	router.HandleFunc(basePath+"/candidate", GetCandidates).Methods("GET")
	router.HandleFunc(basePath+"/candidate/{candidate_id}", GetCandidate).Methods("GET")
}
