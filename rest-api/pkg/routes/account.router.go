package routes

import (
	"github.com/gorilla/mux"
	"github.com/nealarch01/payaco/rest-api/pkg/controllers"
)

func AccountRoutes(router *mux.Router) {
	router.HandleFunc("/user/{id}", controllers.GetPublicAccountData).Methods("GET") // Get public user data by ID
	// Search query for users
	router.HandleFunc("/search", controllers.SearchAccounts).Methods("GET")
}

func SecuredAccountRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.GetPrivateAccountData).Methods("GET") // Get private user data by ID
	router.HandleFunc("/password", controllers.UpdatePassword).Methods("POST") 
	router.HandleFunc("/name", controllers.UpdateName).Methods("PUT") // Should it be a PUT or a POST? I think PUT is more appropriate
	router.HandleFunc("/email", controllers.UpdateEmail).Methods("PUT")
	router.HandleFunc("/phone", controllers.UpdatePhone).Methods("PUT")
}
