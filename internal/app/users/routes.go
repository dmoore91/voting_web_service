package users

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, basePath string) {
	router.HandleFunc(basePath+"/user", AddUser).Methods("POST")
	router.HandleFunc(basePath+"/user/{username}", UpdateUser).Methods("PUT")
	router.HandleFunc(basePath+"/user/{username}", GetUser).Methods("GET")
	router.HandleFunc(basePath+"/user/login", LoginUser).Methods("POST")

	// Order below is very important because if /user/permission/{username} isn't first it will never get called
	// because the other requests will suck it up. Go matched these requests in a top down order and takes the first
	// thing to match
	router.HandleFunc(basePath+"/user/permission/{username}", GetPermissionsForUser).Methods("GET")
	router.HandleFunc(basePath+"/user/{username}/{permission}", AddPermissionForUser).Methods("POST")
	router.HandleFunc(basePath+"/user/{username}/{permission}", RemovePermissionForUser).Methods("DELETE")
}
