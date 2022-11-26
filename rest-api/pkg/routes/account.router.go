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
	router.HandleFunc("/data", controllers.GetPrivateAccountData).Methods("GET") // Get private user data by ID
	router.HandleFunc("/update/password", controllers.UpdatePassword).Methods("POST") 
	router.HandleFunc("/update/name", controllers.UpdateName).Methods("PUT") // Should it be a PUT or a POST? I think PUT is more appropriate
	router.HandleFunc("/update/email", controllers.UpdateEmail).Methods("PUT")
	router.HandleFunc("/update/phone", controllers.UpdatePhone).Methods("PUT")
}
