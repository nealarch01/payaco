package controllers

import (
	"fmt"
	"github.com/nealarch01/payaco/pkg/authentication"
	"github.com/nealarch01/payaco/pkg/models"
	"net/http"
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
	accountID, err := models.VerifyLogin(userIdentifier, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid credentials"}`))
		return
	}
	
	// If the account exists and the password is correct, generate a token
	token := authentication.CreateToken(*accountID)

	// Return successful login and token
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message": "Successfully logged in", "token": "%s"}`, token)))
}
