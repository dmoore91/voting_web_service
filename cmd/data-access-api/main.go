package main

import (
	"net/http"
	"strings"
	"time"
	"voting_web_service/internal/app/ping"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{"context": "main"})

func main() {
	BasePath := "/voting"
	router := mux.NewRouter()
	router.Use(Middleware)
	AttachSwaggerDocs(router, BasePath)
	// Initialize all Routes
	InitializeRoutes(router, BasePath)

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
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Compare("/dal/ping", r.URL.Path) != 0 {
			log.Info(r.Method + " " + r.URL.Path)
		}
		h.ServeHTTP(w, r)
	})
}

func AttachSwaggerDocs(router *mux.Router, basePath string) {
	sh := http.StripPrefix(basePath+"/swagger/", http.FileServer(http.Dir("./docs/swagger/")))
	router.PathPrefix(basePath + "/swagger/").Handler(sh)
}
