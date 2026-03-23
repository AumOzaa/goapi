package middleware

import (
	"errors"
	"net/http"

	"fmt"
	"github.com/AumOzaa/goapi/api"
	"github.com/AumOzaa/goapi/internal/tools"
	// "github.com/go-chi/chi"
	log "github.com/sirupen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("In the authorization")
		var username string = r.URL.Query().Get("username")
		fmt.Printf("\nTHe username is : %v\n", username)
		var token = r.Header.Get("Authorization")
		fmt.Printf("\nThe token value is : %v", token)
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		fmt.Printf("\nRUN-TIME : The value of the current database after initilization is %p\n", database)

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		//query the database :
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		fmt.Printf("\nRUN-TIME : This login details is a shared memory one %p\n", loginDetails)

		// if not found the client
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r) // calls the next middleware in line THi'd call the GetCOinBalance func
	})
}
