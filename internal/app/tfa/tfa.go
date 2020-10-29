package tfa

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/exec"
	"voting_web_service/internal/app/responses"
)

type TFA struct {
	Secret string `json:"secret"`
	Data   string `json:"data"`
}

func GetTfa(writer http.ResponseWriter, request *http.Request) {

	out, err := exec.Command("/bin/bash", "internal/app/tfa/sample.sh").Output()
	mydir, _ := os.Getwd()
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed to get 2FA secret and data "+mydir)
		log.Error(err)
		return
	}

	r := bytes.NewReader(out)

	decoder := json.NewDecoder(r)
	var t TFA
	err = decoder.Decode(&t)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Decode Failed")
		log.Error(err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(t)
}

func Validate(writer http.ResponseWriter, request *http.Request) {
	token := "GET IT FROM THE POST DATA that is sent through the frontend"
	secret := "GET IT FROM THE USER TABLE"
	out, err := exec.Command("/bin/bash", "internal/app/tfa/validate.sh " + secret + " " + token).Output()
}
