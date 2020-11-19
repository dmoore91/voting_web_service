package session

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/session/sign_out/{user}", SetSessionIdNull).Methods("PUT")
	router.HandleFunc(basePath+"/session/{username}/{session_id}", VerifySessionID).Methods("GET")
}
