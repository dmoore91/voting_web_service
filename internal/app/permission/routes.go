package permission

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/permission", AddPermission).Methods("POST")
	router.HandleFunc(basePath+"/permission/{permission}", GetUsersForPermission).Methods("GET")
}
