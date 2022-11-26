package routes 

import (
	"github.com/nealarch01/payaco/pkg/controllers"
	"github.com/gorilla/mux"
)


func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
}