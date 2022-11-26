# Backend

### Disclaimer: This is all play/fake money!!


# Table Of Contents:
1. [Technology Stack](#techstack)
2. [Database](#initializing-the-database)
3. [Backend](#running-the-backend)
4. [Endpoints](#api-endpoints)
   - [Accounts](#account-public)
   - [Authentication](#authentication)
   - [Transactions](#transaction)
5. [Useful Resources](#useful-resources)


---

## Tech Stack
- **Language:** Go
- **Framework:** GorillaMux
- **Database:** PostgreSQL (Version: 15.0)
- **Additional Technologies:** JSON Web Tokens (JWTs)
- **Architecture:** Model-View-Controller (MVC)


---


## Initializing the database
- PostgreSQL Instance
  - Instance running at port 4323
  - Running ```./setup.sh``` will create a database named "development", tables, and some sample account data
- Navigate into ```./backend/database``` and run ```./setup.sh```
  - Note: ```setup.sh``` uses zsh by default, make sure to change this if you use bash.


---

## Running the backend
A shell script <b>run.sh</b> was provided to compile and execute ```./run.sh```
  - Note: ```run.sh``` uses zsh by default, make sure to change this if you use bash.

Alternatively, you may use ```go run main.go```

#### Additional Information:
- The server does not use HTTPS
- The default port (specified in main) is 8001
- Most endpoints besides "/" will return application/json


---


## API Endpoints
### Account (Public)
**Get user information**
```
/public/accounts/user/{id}

Response:
{
	"id": 1,
	"username": "nealarch01",
	"first_name": "Neal",
	"last_name": "A"
}
```

**Search for user**
```
/public/accounts/search?q=neal

Response:
{
	"accounts": [
		{
			"id": 1,
			"username": "nealarch01",
			"first_name": "Neal",
			"last_name": "A"
		},
		{
			"id": 13,
			"username": "nealarch03",
			"first_name": "Neal",
			"last_name": "A"
		}
	]
}
```

---

### Account (Private)
**Get account data (includes sensitive information)**
```
/account/data
\ -- Authorization: "token"

Response:
{
	"balance": 168.58,
	"email": "nealarch01@gmail.com",
	"first_name": "Neal",
	"id": 1,
	"last_name": "A",
	"username": "nealarch01"
}
```

---

### Authentication
**Logging in**
```
/auth/login
\ -- Content-Type: "x-www-form-urlencoded"
		"user_identifier": "usernameOrEmail"
		"password": "password"

Response:
{
	"message": "Successfully logged in",
	"token": "jwt"
}
```

**Create new account**
```
/auth/register
\ -- Content-Type: "x-www-form-urlencoded"
		"username": "username1"
		"email": "me@email.com"
		"phone_number": "+0-111-222-3333"
		"password": "password"
		"first_name": "firstName"
		"last_name": "lastName"

Response: 
{
	"account_id": 13,
	"message": "Successfully registered",
	"token": "ey.."
}
```

---

### Transaction
**Deposit**
```
/transaction/deposit
\ -- Authorization: "token"
\ -- Content-Type: "x-www-form-urlencoded"
		"amount": 1000

Response:
{
	"message": "Deposit successful",
	"new_balance": 768.58
}
```


**Withdraw**
```
/transaction/withdraw
\ -- Authorization: "token"
\ -- Content-Type: "x-www-form-urlencoded"
		"amount": 1000

Response:
{
	"message": "Withdraw successful",
	"new_balance": 668.58
}
```


**Create/Conduct Payment/Transaction**
```
/transaction/create
\ -- Authorization: "token"
\ -- Content-Type: "x-www-form-urlencoded"
				"receiver": "phonenum/email/username"

Response: 
{
	"create_id": 1,
	"message": "Transaction successful"
}
```

**History**
```
/transaction/history
\ -- Authorization: "token"

Response:
{
	"transactions": [
		{
			"id": 1,
			"sender": 1,
			"receiver": 13,
			"amount": 20
		}
	]
}
```

---

### Useful Resources:
- [Go Rest API youtube](https://youtu.be/jFfo23yIWac)
- [Gorilla/Mux](https://github.com/gorilla/mux#middleware)
- [Go database/sql documentation](https://pkg.go.dev/database/sql)
- [GoJWT Documentation](https://pkg.go.dev/github.com/golang-jwt/jwt/v4)
- [Storing currency PostgresSQL](https://stackoverflow.com/questions/15726535/which-datatype-should-be-used-for-currency)
- [URL Params](https://stackoverflow.com/questions/46045756/retrieve-optional-query-variables-with-gorilla-mux)
- [Middleware](https://www.turing.com/kb/building-middleware-for-node-js)
- [Sample data generator](https://www.mockaroo.com)