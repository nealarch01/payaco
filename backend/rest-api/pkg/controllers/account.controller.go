package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nealarch01/payaco/pkg/models"
)

func GetPublicAccountData(w http.ResponseWriter, r *http.Request) {
	// Get the id from the url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the account from the database
	account, err := models.GetAccountByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the account excluding the password
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"id": %d, "username": "%s", "first_name": "%s", "last_name": "%s", "email": "%s", "balance": %f}`, account.Id, account.Username, account.FirstName, account.LastName, account.Email, account.Balance)
}

func GetPrivateAccountData(w http.ResponseWriter, r *http.Request) {

}

func Authenticate(w http.ResponseWriter, r *http.Request) {

}
