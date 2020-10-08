package permission

import (
	"database/sql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"voting_web_service/internal/app/responses"
)

func AddPermission(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/test")
	if err != nil {
		responses.GeneralSystemFailure(writer, "Cannot connect to db")
		log.Error(err)
		return
	}

	defer db.Close()

	queryString := "INSERT INTO Permissions(permission) " +
		"VALUES(?)"

	r, err := db.Exec(queryString, params["permission"])
	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	rowsAffected, err := r.RowsAffected()

	if err != nil {
		responses.GeneralSystemFailure(writer, "Query Failed")
		log.Error(err)
		return
	}

	if rowsAffected == 0 {
		responses.GeneralSystemFailure(writer, "Query Failed")
		return
	}

	responses.GeneralSuccess(writer, "Success")
}

func GetUsersForPermission(writer http.ResponseWriter, request *http.Request) {
	responses.GeneralNotImplemented(writer)
}
