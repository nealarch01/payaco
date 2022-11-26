package routes

import (
	"github.com/gorilla/mux"
	"github.com/nealarch01/payaco/pkg/controllers"
)

func AccountRoutes(router *mux.Router) {
	router.HandleFunc("/user/{id}", controllers.GetPublicAccountData).Methods("GET") // Get public user data by ID
	// Search query for users
	router.HandleFunc("/search", controllers.SearchAccounts).Methods("GET")
}

func SecuredAccountRoutes(router *mux.Router) {
	router.HandleFunc("/data", controllers.GetPrivateAccountData).Methods("GET") // Get private user data by ID
}
