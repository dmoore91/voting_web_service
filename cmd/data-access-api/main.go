package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
	"voting_web_service/internal/app/candidate"
	"voting_web_service/internal/app/party"
	"voting_web_service/internal/app/permission"
	"voting_web_service/internal/app/ping"
	"voting_web_service/internal/app/users"
	"voting_web_service/internal/app/voting"
)

var log = logrus.WithFields(logrus.Fields{"context": "main"})

func main() {
	BasePath := "/voting"
	router := mux.NewRouter()
	router.Use(Middleware)
	AttachSwaggerDocs(router, BasePath)

	// Initialize all Routes
	InitializeRoutes(router, BasePath)

	// Connect to mysql server
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/voting")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fileServer := http.FileServer(http.Dir("./static")) // New code
	router.Handle("/", fileServer)                      // New code

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	//TODO Add HTTPS

	server := &http.Server{
		Addr:           ":8880",
		Handler:        router,
		TLSConfig:      nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("Ready to Handle Requests")
	log.Fatal(server.ListenAndServe())
}

func InitializeRoutes(router *mux.Router, basePath string) {
	ping.InitializeRoutes(router, basePath)
	users.InitializeRoutes(router, basePath)
	permission.InitializeRoutes(router, basePath)
	party.InitializeRoutes(router, basePath)
	candidate.InitializeRoutes(router, basePath)
	voting.InitializeRoutes(router, basePath)
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Compare("/voting/ping", r.URL.Path) != 0 {
			log.Info(r.Method + " " + r.URL.Path)
		}
		h.ServeHTTP(w, r)
	})
}

func AttachSwaggerDocs(router *mux.Router, basePath string) {
	sh := http.StripPrefix(basePath+"/swagger/", http.FileServer(http.Dir("./docs/swagger/")))
	router.PathPrefix(basePath + "/swagger/").Handler(sh)
}
