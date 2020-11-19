package users

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/user", AddUser).Methods("POST")
	router.HandleFunc(basePath+"/user/{username}", UpdateUser).Methods("PUT")
	router.HandleFunc(basePath+"/user/login", LoginUser).Methods("POST")
	router.HandleFunc(basePath+"/user/permission/{username}", GetPermissionsForUser).Methods("GET")
	router.HandleFunc(basePath+"/user/{username}/{permission}", AddPermissionForUser).Methods("POST")
	router.HandleFunc(basePath+"/user/{username}/{permission}", RemovePermissionForUser).Methods("DELETE")
}
