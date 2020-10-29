package tfa

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"os"
	log "github.com/sirupsen/logrus"
	"voting_web_service/internal/app/responses"
)

type TFAResponse struct {
	TFA string `json:"TFA"`
}

func GetTfa(writer http.ResponseWriter, request *http.Request) {

	out, err := exec.Command("/bin/bash", "internal/app/tfa/sample.sh").Output()
	mydir, _ := os.Getwd()
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed to get 2FA secret and data " + mydir )
		log.Error(err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(&TFAResponse{TFA: string(out)})
}