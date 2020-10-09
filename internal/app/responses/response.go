package responses

import (
	"encoding/json"
	"net/http"
)

// Generalized response for API
//
// swagger:model generalResponse
type GeneralResponse struct {
	// message
	// Required: true
	// in: body
	Message string `json:"message"`
}

// Generalized response for when we need primary key
//
// swagger:model generalPrimaryKeyResponse
type generalPrimaryKey struct {
	// primary key
	// Required: true
	// in: body
	PrimaryKey int `json:"primary_key"`
}

func GeneralSuccess(writer http.ResponseWriter, message string) {
	respond(writer, http.StatusOK, message)
}

func GeneralNoContent(writer http.ResponseWriter, message string) {
	respond(writer, http.StatusNoContent, message)
}

func GeneralBadRequest(writer http.ResponseWriter, message string) {
	respond(writer, http.StatusBadRequest, message)
}

func GeneralSystemFailure(writer http.ResponseWriter, message string) {
	respond(writer, http.StatusInternalServerError, message)
}

func GeneralNotImplemented(writer http.ResponseWriter) {
	respond(writer, http.StatusNotImplemented, "")
}

func GeneralReturnPrimaryKey(writer http.ResponseWriter, primaryKey int) {
	tmp := generalPrimaryKey{PrimaryKey: primaryKey}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(tmp)
}

func respond(writer http.ResponseWriter, result int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(result)
	_ = json.NewEncoder(writer).Encode(&GeneralResponse{Message: message})
}
