package tfa

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"voting_web_service/internal/app/responses"
)

type TFA struct {
	Secret string `json:"secret"`
	Data   string `json:"data"`
}

// swagger:model validationInput
type ValidateStruct struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

func GetTfa(writer http.ResponseWriter, request *http.Request) {
	// GET /GetTfa
	//
	// Endpoint to get the TFA secret

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

func getSecretForUser(writer http.ResponseWriter, username string) (string, bool) {

	db, err := sql.Open("mysql", "root:VV@WF9Xf8C6!#Xy!@tcp(mysql_db:3306)/voting")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return "", false
	}

	defer db.Close()

	queryString := "SELECT secret_key " +
		"FROM Users " +
		"WHERE username=?"

	var secret string
	err = db.QueryRow(queryString, username).Scan(&secret)
	if err != nil {
		responses.GeneralSystemFailure(writer, "Failed query")
		log.Error(err)
		return "", false
	}

	return secret, true
}

func Validate(writer http.ResponseWriter, request *http.Request) {
	// POST /tfa_validate
	//
	// Endpoint to validate 2 factor authentication
	//
	// ---
	// produces:
	// - application/json
	//  parameters:
	//	 - name: validation_input
	//	   in: body
	//	   description: input for validating 2fa
	//	   schema:
	//	     "$ref": "#/definitions/validationInput"
	//	   required: true
	// responses:
	//   '200':
	//     description: User is valid
	//     schema:
	//       "$ref": "#/definitions/candidateList"
	//   '400':
	//     description: bad request
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	//   '500':
	//     description: server errorValidateStruct
	//     schema:
	//       "$ref": "#/definitions/generalResponse"
	decoder := json.NewDecoder(request.Body)
	var v ValidateStruct
	err := decoder.Decode(&v)
	if err != nil {
		responses.GeneralBadRequest(writer, "Decode Failed")
		log.Error(err)
		return
	}

	secret, legit := getSecretForUser(writer, v.Username)
	if legit {
		out, err := exec.Command("./internal/app/tfa/validate.sh", secret, v.Token).
			Output()
		fmt.Println("output", string(out))
		if err != nil {
			responses.GeneralBadRequest(writer, "Validation Script Failed")
			log.Error(err)
			return
		}

		if strings.EqualFold(string(out), "true") {
			responses.GeneralSuccess(writer, "Success")
		} else {
			responses.GeneralBadRequest(writer, "Token invalid")
		}
	} else {
		responses.GeneralSystemFailure(writer, "Failed to get user secret")
	}
}
