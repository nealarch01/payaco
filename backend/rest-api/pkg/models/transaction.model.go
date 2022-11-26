package models 

import (
	"fmt"
)

/* Transaction Schema
CREATE TABLE transactions (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    sender INT REFERENCES account(id),
    receiver INT REFERENCES account(id),
    amount DECIMAL(12, 2)
);
*/

type Transaction struct {
	Id int `json:"id"`
	Sender int `json:"sender"`
	Receiver int `json:"receiver"`
	Amount float64 `json:"amount"`
}


func CreateTransaction(transaction Transaction) (int, error) {
	db := InitConnection()

	if db == nil {
		return -1, fmt.Errorf("database connection failed")
	}

	// Get the balance of both accounts
	senderBalance, err := GetBalance(transaction.Sender)	
	if err != nil {
		return -1, err
	}
	senderNewBalance := senderBalance - transaction.Amount

	receiverBalance, err := GetBalance(transaction.Receiver)
	if err != nil {
		return -1, nil 
	}
	receiverNewBalance := receiverBalance + transaction.Amount
	
	// Add to the transaction table and fetch the new id
	var newId int
	err = db.QueryRow("INSERT INTO transactions (sender, receiver, amount) VALUES ($1, $2, $3) RETURNING id", transaction.Sender, transaction.Receiver, transaction.Amount).Scan(&newId)

	if err != nil {
		return -1, err
	}
	
	// Update the balance of both accounts
	senderUpdateStatus := UpdateBalance(transaction.Sender, senderNewBalance)
	receiverUpdateStatus := UpdateBalance(transaction.Receiver, receiverNewBalance)
	// If either update fails, rollback the transaction
	if senderUpdateStatus != nil || receiverUpdateStatus != nil {
		// Undo the newly created transaction
		_, _ = db.Exec("DELETE FROM transactions WHERE sender = $1 AND receiver = $2", transaction.Sender, transaction.Receiver)
		return -1, fmt.Errorf("transaction failed")
	}

	return newId, nil
}

func GetTransactionHistory(userID int) ([]Transaction, error) {
	var transactions []Transaction = make([]Transaction, 0)
	db := InitConnection()

	if db == nil {
		return transactions, fmt.Errorf("database connection failed")
	}
	rows, err := db.Query("SELECT * FROM transactions WHERE sender = $1", userID)
	if err != nil {
		return transactions, err
	}

	tempTransaction := Transaction{}
	for rows.Next() {
		err := rows.Scan(&tempTransaction.Id, &tempTransaction.Sender, &tempTransaction.Receiver, &tempTransaction.Amount)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, tempTransaction)
	}

	return transactions, nil
}


func Deposit(userID int, amount float64) error {
	db := InitConnection()

	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	// Get the balance of the account
	balance, err := GetBalance(userID)	
	if err != nil {
		return err
	}
	newBalance := balance + amount
	
	// Update the balance of the account
	updateStatus := UpdateBalance(userID, newBalance)
	if updateStatus != nil {
		return updateStatus
	}

	return nil
}

func Withdraw(userID int, amount float64) error {
	db := InitConnection()

	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	// Get the balance of the account
	balance, err := GetBalance(userID)	
	if err != nil {
		return err
	}
	newBalance := balance - amount
	
	// Update the balance of the account
	updateStatus := UpdateBalance(userID, newBalance)
	if updateStatus != nil {
		return updateStatus
	}

	return nil
}