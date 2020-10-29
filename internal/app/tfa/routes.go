package tfa

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/tfa", GetTfa).Methods("GET")
	router.HandleFunc(basePath+"/tfa_validate", Validate).Methods("POST")
}
