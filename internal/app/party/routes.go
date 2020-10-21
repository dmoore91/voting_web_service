package party

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/party", CreateParty).Methods("POST")
	router.HandleFunc(basePath+"/party", GetParties).Methods("GET")
	router.HandleFunc(basePath+"/party/{user_id}", LinkUserAndParty).Methods("POST")
	router.HandleFunc(basePath+"/party/{user_id}", UpdateUserParty).Methods("PUT")
}
