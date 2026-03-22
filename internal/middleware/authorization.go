package middleware

import (
	"errors"
	"net/http"

	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	log "github.com/sirupen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
		}
	})
}
