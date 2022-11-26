package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nealarch01/payaco/rest-api/pkg/authentication"
	"github.com/nealarch01/payaco/rest-api/pkg/models"
	utilities "github.com/nealarch01/payaco/rest-api/pkg/utils"
)

func GetPublicAccountData(w http.ResponseWriter, r *http.Request) {
	// Get the id from the url
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "id was not provided"}`)
		return
	}

	// Get the account from the database
	account, err := models.GetAccountByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"message": "Account does not exist"}`)
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

// Requires a valid token
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
		fmt.Fprint(w, `{"message": "Internal server error"}`)
		return
	}

	data := make(map[string]interface{})
	data["accounts"] = publicAccounts

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}


func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	// Get the ID from the token
	userID := authentication.GetIdFromToken(token)

	// Get the password from the form
	r.ParseForm()
	// Get the old password and compare
	oldPassword := r.FormValue("old_password")
	if oldPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "old_password was not provided"}`)
		return
	}

	// Get the new password
	newPassword := r.Form.Get("new_password")
	if newPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "New password was not provided"}`)
		return
	}

	// Compare the old password with the one in the database
	account, err := models.GetAccountByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Internal server error. Try again."}`)
		return
	}

	if account.Password != oldPassword {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "Incorrect password"}`)
		return
	}

	// Update the password
	err = models.UpdatePassword(userID, newPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Internal server error"}`)
		return
	}

	// Create a new token and blacklist the old one
	newToken := authentication.CreateToken(userID)
	authentication.BlacklistToken(token)


	w.WriteHeader(http.StatusOK)
	data := make(map[string]interface{})
	data["message"] = "Password updated"
	data["token"] = newToken
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}


func UpdateEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	// Get the ID from the token
	userID := authentication.GetIdFromToken(token)

	// Get the email from the body
	r.ParseForm()
	email := r.Form.Get("email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Email was not provided"}`)
		return
	}

	// Update the email
	err := models.UpdateEmail(userID, email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Internal server error"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Email updated"}`)
}

func UpdatePhone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	// Get the ID from the token
	userID := authentication.GetIdFromToken(token)

	// Get the phone from the form
	r.ParseForm()
	phone := r.Form.Get("phone")
	if phone == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Phone was not provided"}`)
		return
	}

	// Update the phone
	err := models.UpdatePhone(userID, phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Internal server error"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Phone updated"}`)
}

func UpdateName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	// Get the ID from the token
	userID := authentication.GetIdFromToken(token)

	// Get the first name and last name from the form
	r.ParseForm()
	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	if firstName == "" || lastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "First name or last name was not provided"}`)
		return
	}

	// Update the first and last name
	err := models.UpdateName(userID, firstName, lastName)	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Internal server error"}`)
		return
	}

	// Return successful update)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Name successfully updated"}`)
}
