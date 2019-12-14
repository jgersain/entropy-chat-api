package middlewares

import (
	"errors"
	"github.com/jgersain/entropy-chat-api/api/auth"
	"github.com/jgersain/entropy-chat-api/api/utils"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
			return
		}
		next(w, r)
	}
}
