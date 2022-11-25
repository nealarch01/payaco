package routes

import (
	"github.com/gorilla/mux"
	"github.com/nealarch01/payaco/pkg/controllers"
)

func AccountRoutes(router *mux.Router) {
	router.HandleFunc("/{id}", controllers.GetPublicAccountData).Methods("GET") // Get public user data by ID
}

func SecuredAccountRoutes(router *mux.Router) {
	router.HandleFunc("/{id}", controllers.GetPrivateAccountData).Methods("GET") // Get private user data by ID
}
