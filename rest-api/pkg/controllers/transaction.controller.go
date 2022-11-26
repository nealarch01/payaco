package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nealarch01/payaco/rest-api/pkg/authentication"
	"github.com/nealarch01/payaco/rest-api/pkg/models"
	utilities "github.com/nealarch01/payaco/rest-api/pkg/utils"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Get the form data
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	userID := authentication.GetIdFromToken(token) // Get the ID from the token
	if userID <= 0 {
		w.WriteHeader(http.StatusUnauthorized)
		// return json
		fmt.Fprint(w, `{"message": "Invalid token. User ID not found"}`)
		return
	}
	r.ParseForm()

	receiver := r.Form.Get("receiver")
	var receiverAccount models.Account
	// Check if receiver is email, phone or username
	if utilities.IsPhoneNumber(receiver) {
		acc, err := models.GetAccountByPhone(receiver)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"message": "No account found with that phone number"}`)
			return
		}
		receiverAccount = acc
	} else if utilities.IsEmail(receiver) {
		acc, err := models.GetAccountByEmail(receiver)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"message": "No account associated with email provided"}`)
			return
		}
		receiverAccount = acc
	} else {
		acc, err := models.GetAccountByUsername(receiver)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "No account associated with username provided")
			return
		}
		receiverAccount = acc
	}

	receiverID := receiverAccount.Id

	if receiverID == userID {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You cannot send money to yourself")
		return
	}

	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Invalid amount"}`)
		return
	}

	if userID <= 0 || receiverID <= 0 || amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Invalid value was entered"}`)
		return
	}

	// Check if the user has enough balance
	balance, err := models.GetBalance(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Error getting balance. Transaction cancelled"}`)
		return
	}

	if balance < amount {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Insufficient balance"}`)
		return
	}

	createID, err := models.CreateTransaction(models.Transaction{Sender: userID, Receiver: receiverID, Amount: amount})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Transaction failed. No charges were made"}`)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
	data := make(map[string]interface{})
	data["create_id"] = createID
	data["message"] = "Transaction successful"

	// Convert the data to JSON string
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}

func GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")

	// Get the ID from the token
	userID := authentication.GetIdFromToken(token)
	if userID <= 0 {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "Invalid token. User ID not found"}`)
		return
	}

	transactions, err := models.GetTransactionHistory(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Error getting transaction history"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := make(map[string]interface{})
	data["transactions"] = transactions
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	userID := authentication.GetIdFromToken(token) // Get the ID from the token
	if userID <= 0 {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "Invalid token. User ID not found"}`)
		return
	}

	r.ParseForm()
	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Invalid amount"}`)
		return
	}

	if amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Invalid amount"}`)
		return
	}

	balance, err := models.GetBalance(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Error getting balance. Transaction cancelled"}`)
		return
	}

	// Calculate the new balance
	newBalance := balance + amount
	err = models.UpdateBalance(userID, newBalance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Error updating balance. Transaction cancelled"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := make(map[string]interface{})
	data["message"] = "Deposit successful"
	data["new_balance"] = newBalance

	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	userID := authentication.GetIdFromToken(token) // Get the ID from the token
	if userID <= 0 {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid token. User ID not found")
		return
	}

	r.ParseForm()
	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Invalid amount"}`)
		return
	}

	if amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Invalid amount"}`)
		return
	}

	balance, err := models.GetBalance(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Error getting balance. Transaction cancelled"}`)
		return
	}

	// Check if the user's balance is enough to withdraw
	if amount > balance {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Insufficient balance"}`)
		return
	}

	// Calculate the new balance
	newBalance := balance - amount
	err = models.UpdateBalance(userID, newBalance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Error updating balance. Transaction cancelled"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := make(map[string]interface{})
	data["message"] = "Withdraw successful"
	data["new_balance"] = newBalance

	jsonData, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonData))
}
