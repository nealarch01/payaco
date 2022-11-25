package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/nealarch01/payaco/pkg/middlewares"
	"github.com/nealarch01/payaco/pkg/models"
	"github.com/nealarch01/payaco/pkg/routes"
)

func entryPoint(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, world! Server is running.</h1>")
	// Return a 200 OK status
	w.WriteHeader(http.StatusOK)
	// Add content type
	w.Header().Set("Content-Type", "text/html")
}

// Checks if the database is up when the server starts
func isDatabaseUp() bool {
	connection := models.InitConnection()
	if connection == nil {
		return false
	}
	connection.Close() // Close the initial connection
	return true
}

func main() {
	portNumber := 8001
	// connectionStr := "user=nealarchival dbname=banking-app port=4323"
	if isDatabaseUp() {
		fmt.Println("Database is up")
	} else {
		fmt.Println("Failed to make a database connection. Terminating program.")
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/", entryPoint)        // Entry point, display an HTML page
	r.Use(middlewares.LoggingMiddleware) // Middleware that logs all incoming requests.

	authRoutes := r.PathPrefix("/api/auth").Subrouter()
	routes.AuthRoutes(authRoutes)

	accountRoutes := r.PathPrefix("/api/accounts").Subrouter()
	routes.AccountRoutes(accountRoutes)

	fmt.Println("Server is running on port:", portNumber)
	http.ListenAndServe(":"+strconv.Itoa(portNumber), r) // r is the Router
}
