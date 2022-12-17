package models

import (
	"fmt"
	"strings"
	"context"
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

type PublicAccount struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (publicAccount *PublicAccount) ToPublicAccount(account Account) {
	publicAccount.Id = account.Id
	publicAccount.Username = account.Username
	publicAccount.FirstName = account.FirstName
	publicAccount.LastName = account.LastName
}

func GetAccountByID(id int) (Account, error) {
	var account Account
	db := GetConnection()
	if db == nil {
		return account, fmt.Errorf("database connection failed")
	}

	err := db.QueryRowContext(context.Background(), "SELECT * FROM account WHERE id = $1", id).Scan(&account.Id, &account.Username, &account.FirstName, &account.LastName, &account.PhoneNumber, &account.Email, &account.Password, &account.Balance)
	if err != nil {
		fmt.Println(err)
		return account, err
	}

	return account, nil
}

func GetAccountByEmail(email string) (Account, error) {
	var account Account
	db := GetConnection()
	if db == nil {
		return account, fmt.Errorf("database connection failed")
	}

	err := db.QueryRowContext(context.Background(), "SELECT * FROM account WHERE email = $1", email).Scan(&account.Id, &account.Username, &account.FirstName, &account.LastName, &account.PhoneNumber, &account.Email, &account.Password, &account.Balance)
	if err != nil {
		fmt.Println(err)
		return account, err
	}

	return account, nil
}

func GetAccountByPhone(phone string) (Account, error) {
	var account Account
	db := GetConnection()
	if db == nil {
		return account, fmt.Errorf("database connection failed")
	}

	err := db.QueryRowContext(context.Background(), "SELECT * FROM account WHERE phone_number = $1", phone).Scan(&account.Id, &account.Username, &account.FirstName, &account.LastName, &account.PhoneNumber, &account.Email, &account.Password, &account.Balance)
	if err != nil {
		fmt.Println(err)
		return account, err
	}

	return account, nil
}

func GetAccountByUsername(username string) (Account, error) {
	var account Account
	db := GetConnection()
	if db == nil {
		return account, fmt.Errorf("database connection failed")
	}

	err := db.QueryRowContext(context.Background(), "SELECT * FROM account WHERE username = $1", username).Scan(&account.Id, &account.Username, &account.FirstName, &account.LastName, &account.PhoneNumber, &account.Email, &account.Password, &account.Balance)
	if err != nil {
		fmt.Println(err)
		return account, err
	}

	return account, nil
}

func AccountsCount() (int, error) {
	db := GetConnection()
	if db == nil {
		return 0, fmt.Errorf("database connection failed")
	}

	var count int
	err := db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM account").Scan(&count)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return count, nil
}

func SearchUsername(username string) ([]Account, error) {
	db := GetConnection()
	if db == nil {
		return nil, fmt.Errorf("database connection failed")
	}

	var accounts []Account
	rows, err := db.QueryContext(context.Background(), "SELECT * FROM account WHERE username LIKE $1", username+"%") // % is wildcard
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var account Account
		err = rows.Scan(&account.Id, &account.Username, &account.FirstName, &account.LastName, &account.PhoneNumber, &account.Email, &account.Password, &account.Balance)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func CheckAccountExists(userIdentifier string) (bool, error) {
	db := GetConnection()
	if db == nil {
		return false, fmt.Errorf("database connection failed")
	}

	// Checks if an email or username exists
	queryString := ""
	if strings.Contains(userIdentifier, "@") {
		queryString = "SELECT COUNT(*) FROM account WHERE email = $1"
	} else {
		queryString = "SELECT COUNT(*) FROM account WHERE username = $1"
	}

	var count int
	err := db.QueryRowContext(context.Background(), queryString, userIdentifier).Scan(&count)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return count > 0, nil
}

func GetBalance(id int) (float64, error) {
	db := GetConnection()
	if db == nil {
		return 0, fmt.Errorf("database connection failed")
	}

	var balance float64
	err := db.QueryRowContext(context.Background(), "SELECT balance FROM account WHERE id = $1", id).Scan(&balance)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return balance, nil
}

// Update Functions

func UpdateBalance(id int, newBalance float64) error {
	db := GetConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.ExecContext(context.Background(), "UPDATE account SET balance = $1 WHERE id = $2", newBalance, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}


func UpdateEmail(id int, newEmail string) error {
	db := GetConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.ExecContext(context.Background(), "UPDATE account SET email = $1 WHERE id = $2", newEmail, id)
	if err != nil {
		return err
	}

	return nil
}


func UpdatePhone(id int, newPhone string) error {
	db := GetConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.ExecContext(context.Background(), "UPDATE account SET phone_number = $1 WHERE id = $2", newPhone, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePassword(id int, newPassword string) error {
	db := GetConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.ExecContext(context.Background(), "UPDATE account SET password = $1 WHERE id = $2", newPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFirstName(id int, newFirstName string) error {
	db := GetConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.ExecContext(context.Background(), "UPDATE account SET first_name = $1 WHERE id = $2", newFirstName, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateLastName(id int, newLastName string) error {
	db := GetConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.ExecContext(context.Background(), "UPDATE account SET last_name = $1 WHERE id = $2", newLastName, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateName(id int, newFirstName string, newLastName string) error {
	db := GetConnection()
	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := db.ExecContext(context.Background(), "UPDATE account SET first_name = $1, last_name = $2 WHERE id = $3", newFirstName, newLastName, id)
	if err != nil {
		return err
	}

	return nil
}

// Authentication functions

func Login(userIdentifier string, password string) (int, error) {
	db := GetConnection()
	if db == nil {
		return 0, fmt.Errorf("database connection failed")
	}

	// Checks if an email or username exists
	queryString := ""
	if strings.Contains(userIdentifier, "@") {
		queryString = "SELECT id FROM account WHERE email = $1 AND password = $2"
	} else {
		queryString = "SELECT id FROM account WHERE username = $1 AND password = $2"
	}

	var id int = 0
	err := db.QueryRowContext(context.Background(), queryString, userIdentifier, password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func Register(account Account) (int, error) {
	db := GetConnection()
	if db == nil {
		return 0, fmt.Errorf("database connection failed")
	}

	account.Balance = 0

	var id int
	err := db.QueryRowContext(context.Background(), "INSERT INTO account (username, first_name, last_name, phone_number, email, password, balance) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", account.Username, account.FirstName, account.LastName, account.PhoneNumber, account.Email, account.Password, account.Balance).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}