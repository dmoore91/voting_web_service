package party

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/party/link/{user}/{party}", LinkUserAndParty).Methods("POST") //Don't move this
	router.HandleFunc(basePath+"/party/{party}", CreateParty).Methods("POST")
	router.HandleFunc(basePath+"/party", GetParties).Methods("GET")
}
