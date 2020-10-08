package users

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/user", AddUser).Methods("POST")
	router.HandleFunc(basePath+"/user", UpdateUser).Methods("PUT")
	router.HandleFunc(basePath+"/user/{user_id}", GetUser).Methods("GET")
	router.HandleFunc(basePath+"/user/login", LoginUser).Methods("POST")
}
