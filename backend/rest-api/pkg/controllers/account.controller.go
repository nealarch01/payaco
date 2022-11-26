package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nealarch01/payaco/pkg/authentication"
	"github.com/nealarch01/payaco/pkg/models"
	utilities "github.com/nealarch01/payaco/pkg/utils"
)

func GetPublicAccountData(w http.ResponseWriter, r *http.Request) {
	// Get the id from the url
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "id was not provided"}`)
		return
	}

	// Get the account from the database
	account, err := models.GetAccountByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error": "Account does not exist"}`)
		return
	}

	// Return the account excluding the password
	w.WriteHeader(http.StatusOK)
	// Create a dictionary with keys
	data := make(map[string]interface{})
	data["username"] = account.Username
	data["first_name"] = account.FirstName
	data["last_name"] = account.LastName

	// Convert the dictionary to a json string
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}

func GetPrivateAccountData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	// Get the ID from the token
	userID := authentication.GetIdFromToken(token)

	account, err := models.GetAccountByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return ID, username, first name, last name, email, and balance
	w.WriteHeader(http.StatusOK)
	data := make(map[string]interface{})
	data["id"] = account.Id
	data["username"] = account.Username
	data["first_name"] = account.FirstName
	data["last_name"] = account.LastName
	data["email"] = account.Email
	data["balance"] = account.Balance

	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}

func SearchAccounts(w http.ResponseWriter, r *http.Request) {
	// Get the username from the url
	w.Header().Set("Content-Type", "application/json")
	// Get the userIdentifier from the params
	userIdentifier := r.URL.Query().Get("q")
	if userIdentifier == "" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"accounts": null}`)
		return
	}

	isEmail := utilities.IsEmail(userIdentifier)
	if isEmail {
		account, err := models.GetAccountByEmail(userIdentifier)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		data := make(map[string]interface{})
		data["username"] = account.Username
		data["first_name"] = account.FirstName
		data["last_name"] = account.LastName
		jsonData, _ := json.Marshal(data)
		fmt.Fprint(w, string(jsonData))
		return
	}

	accounts, err := models.SearchUsername(userIdentifier)
	publicAccounts := make([]models.PublicAccount, 0) // Initialize as an empty array

	for _, account := range accounts {
		var publicAccount models.PublicAccount
		publicAccount.ToPublicAccount(account)
		publicAccounts = append(publicAccounts, publicAccount)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error": "Internal server error"}`)
		return
	}

	data := make(map[string]interface{})
	data["accounts"] = publicAccounts

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}
