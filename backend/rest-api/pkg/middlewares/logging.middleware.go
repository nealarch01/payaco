package middlewares

import (
	"fmt"
	"net/http"
)

// Logs all requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the client ip, method, url, and headers
		fmt.Printf("Client Address: %s \n %s %s %s \nHeaders:%s", r.RemoteAddr, r.Method, r.URL, r.Proto, r.Header)
		next.ServeHTTP(w, r)
	})
}
