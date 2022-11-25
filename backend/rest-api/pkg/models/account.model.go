package models

import (
	"fmt"
	"strings"
)

/* account table schema
CREATE TABLE account (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(32) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(65) NOT NULL,
    BALANCE DECIMAL(12, 2) NOT NULL
);
*/

type Account struct {
	Id          int     `json:"id"`
	Username    string  `json:"username"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Balance     float64 `json:"balance"`
}

func GetAccountByID(id int) (Account, error) {
	var account Account
	db := InitConnection()
	if db == nil {
		return account, fmt.Errorf("database connection failed")
	}

	err := db.QueryRow("SELECT * FROM account WHERE id = $1", id).Scan(&account.Id, &account.Username, &account.FirstName, &account.LastName, &account.PhoneNumber, &account.Email, &account.Password, &account.Balance)
	if err != nil {
		fmt.Println(err)
		return account, err
	}

	return account, nil
}

func VerifyLogin(userIdentifier string, password string) (*int, error) {
	db := InitConnection()
	if db == nil {
		return nil, fmt.Errorf("database connection failed")
	}

	// Checks if an email or username exists
	queryString := ""
	if strings.Contains(userIdentifier, "@") {
		queryString = "SELECT id FROM accounts WHERE email = $1 AND password = $2"
	} else {
		queryString = "SELECT id FROM accounts WHERE username = $1 AND password = $2"
	}

	var id *int = nil
	err := db.QueryRow(queryString, userIdentifier, password).Scan(&id)
	if err != nil {
		print(err)
		return nil, err
	}

	return id, nil
}

func CheckAccountExists(userIdentifier string) (bool, error) {
	db := InitConnection()
	if db == nil {
		return false, fmt.Errorf("database connection failed")
	}

	// Checks if an email or username exists
	queryString := ""
	if strings.Contains(userIdentifier, "@") {
		queryString = "SELECT COUNT(*) FROM accounts WHERE email = $1"
	} else {
		queryString = "SELECT COUNT(*) FROM accounts WHERE username = $1"
	}

	var count int
	err := db.QueryRow(queryString, userIdentifier).Scan(&count)
	if err != nil {
		print(err)
		return false, err
	}

	return count > 0, nil
}

func GetBalance(id int) (float64, error) {
	db := InitConnection()
	if db == nil {
		return 0, fmt.Errorf("database connection failed")
	}

	var balance float64
	err := db.QueryRow("SELECT balance FROM account WHERE id = $1", id).Scan(&balance)
	if err != nil {
		print(err)
		return 0, err
	}

	return balance, nil
}

func UpdateBalance(id int, newBalance float64) error {
	db := InitConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.Exec("UPDATE account SET balance = $1 WHERE id = $2", newBalance, id)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func AccountsCount() (int, error) {
	db := InitConnection()
	if db == nil {
		return 0, fmt.Errorf("database connection failed")
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM accounts").Scan(&count)
	if err != nil {
		print(err)
		return 0, err
	}

	return count, nil
}
