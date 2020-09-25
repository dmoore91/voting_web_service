package ping

import (
	"encoding/json"
	"net/http"
)

type pongResponse struct {
	Ping string `json:"ping"`
}

func Ping(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(&pongResponse{Ping: "Pong"})
}
