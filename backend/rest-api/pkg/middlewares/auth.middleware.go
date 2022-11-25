package middlewares

import (
	"fmt"
	"github.com/nealarch01/payaco/pkg/authentication"
	"net/http"
)

// Auth Middlware, checks if access token is valid
func AuthenticationTokenValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the header
		token := r.Header.Get("Authorization")
		if token == "" {
			fmt.Println("No token provided")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Token was not provided"}`))
			return
		}
		isValid := authentication.ValidateToken(token)
		if !isValid {
			fmt.Println("Invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			// Return JSON response
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "This token is invalid"}`))
			return
		}
		// If the token is valid, continue
		next.ServeHTTP(w, r)
	})
}
