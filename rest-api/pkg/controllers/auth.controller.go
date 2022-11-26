package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nealarch01/payaco/rest-api/pkg/authentication"
	"github.com/nealarch01/payaco/rest-api/pkg/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// If the request body is not x-www-form-urlencoded, return an error
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Invalid content type. Must be application/x-www-form-urlencoded"}`))
		return
	}

	// Get the userIdentifier (username or email) and password from the request body
	r.ParseForm()
	userIdentifier := r.Form.Get("user_identifier")
	password := r.Form.Get("password")

	// Get the account from the database
	accountExists, err := models.CheckAccountExists(userIdentifier)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Something went wrong. Try again."}`))
		return
	}

	if !accountExists {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No account with that username or email exists"}`))
		return
	}

	// Now that we verified an associated account exists, attempt to authenticate
	accountID, err := models.Login(userIdentifier, password)
	if err != nil || accountID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid credentials"}`))
		return
	}

	// If the account exists and the password is correct, generate a token
	token := authentication.CreateToken(accountID)

	// Return successful login and token
	w.WriteHeader(http.StatusOK)
	data := make(map[string]interface{})
	data["message"] = "Successfully logged in"
	data["token"] = token
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}

func Register(w http.ResponseWriter, r *http.Request) {
	// If the request body is not x-www-form-urlencoded, return an error
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Invalid content type. Must be application/x-www-form-urlencoded"}`))
		return
	}

	// Get the form data
	r.ParseForm()
	username := r.Form.Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Username cannot be empty"}`))
		return
	}
	firstName := r.Form.Get("first_name")
	if firstName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "First name cannot be empty"}`))
		return
	}
	lastName := r.Form.Get("last_name")
	if firstName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "First name cannot be empty"}`))
		return
	}
	phone_number := r.Form.Get("phone_number")
	if phone_number == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Phone number cannot be empty"}`))
		return
	}
	email := r.Form.Get("email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Email cannot be empty"}`))
		return
	}
	password := r.Form.Get("password")
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Password cannot be empty"}`))
		return
	}

	// Check if the username or email already exists
	usernameExists, err := models.CheckAccountExists(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Something went wrong. Try again."}`))
		return
	}

	emailExists, err := models.CheckAccountExists(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Something went wrong. Try again."}`))
		return
	}

	if usernameExists {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"message": "Username is already taken"}`))
		return
	}

	if emailExists {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"message": "Email is already taken"}`))
		return
	}

	// Create the account
	accountID, err := models.Register(models.Account{Username: username, FirstName: firstName, LastName: lastName, PhoneNumber: phone_number, Email: email, Password: password})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Something went wrong. Try again."}`))
		return
	}

	// Return successful registration and account ID
	w.WriteHeader(http.StatusCreated)
	data := make(map[string]interface{})
	data["message"] = "Successfully registered"
	data["account_id"] = accountID
	// Create a JWT
	token := authentication.CreateToken(accountID)
	data["token"] = token
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}
