package routes

import (
	"github.com/gorilla/mux"
	"github.com/nealarch01/payaco/pkg/controllers"
)


func TransactionRoutes(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateTransaction).Methods("POST")
	router.HandleFunc("/history", controllers.GetTransactionHistory).Methods("GET")
	router.HandleFunc("/deposit", controllers.Deposit).Methods("POST")
	router.HandleFunc("/withdraw", controllers.Withdraw).Methods("POST")
}
