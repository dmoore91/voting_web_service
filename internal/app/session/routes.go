package session

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/session/sign_out/{user}", SetSessionIdNull).Methods("PUT")
}
