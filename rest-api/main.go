package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/nealarch01/payaco/rest-api/pkg/middlewares"
	"github.com/nealarch01/payaco/rest-api/pkg/models"
	"github.com/nealarch01/payaco/rest-api/pkg/routes"
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
	return connection != nil
}

func main() {
	portNumber := 8001
	// connectionStr := "user=nealarchival dbname=banking-app port=4323"
	if isDatabaseUp() {
		fmt.Println("Database is up")
	} else {
		fmt.Println("Failed to make a database connection. Terminating program.")
		os.Exit(1)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/", entryPoint)        // Entry point, display an HTML page
	r.Use(middlewares.LoggingMiddleware) // Middleware that logs all incoming requests.

	authRoutes := r.PathPrefix("/auth").Subrouter()
	routes.AuthRoutes(authRoutes)

	// Public routes
	publicAccountRoutes := r.PathPrefix("/public/accounts").Subrouter() // Public routes, accessible to everyone
	routes.AccountRoutes(publicAccountRoutes)

	// Private routes
	securedAccountRoutes := r.PathPrefix("/account").Subrouter() // Secured routes, accessible only to authenticated users
	securedAccountRoutes.Use(middlewares.AuthenticationTokenValid)
	routes.SecuredAccountRoutes(securedAccountRoutes)

	// Private routes
	transactionRoutes := r.PathPrefix("/transaction").Subrouter() // Secured routes, accessible only to authenticated users
	transactionRoutes.Use(middlewares.AuthenticationTokenValid)
	routes.TransactionRoutes(transactionRoutes)

	fmt.Println("Server is running on port:", portNumber)
	http.ListenAndServe(":"+strconv.Itoa(portNumber), r) // r is the Router
}
