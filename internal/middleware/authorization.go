package middleware

import (
	"errors"
	"net/http"

	"github.com/AumOzaa/goapi/api"
	"github.com/AumOzaa/goapi/internal/tools"
	// "github.com/go-chi/chi"
	log "github.com/sirupen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		//query the database :
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		// if not found the client
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r) // calls the next middleware in line THi'd call the GetCOinBalance func
	})
}
