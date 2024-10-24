package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth = jwtauth.New("HS256", []byte("secretkey"), nil)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		} else {
			http.Error(w, "Authorization header must be of type Bearer", http.StatusUnauthorized)
			return
		}

		_, err := tokenAuth.Decode(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
