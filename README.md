# Backend

### Disclaimer: This is all play/fake money!!


# Table of Contents:
1. [Technology Stack](#techstack)
2. [Database](#initializing-the-database)
3. [Running the RestAPI](#running-the-rest-api)
4. [API Documentation](#api-endpoints)
5. [Useful Resources](#useful-resources)


## Tech Stack
- **Language:** Go
- **Framework:** GorillaMux
- **Database:** PostgreSQL (Version: 15.0)
- **Additional Technologies:** JSON Web Tokens (JWTs)
- **Architecture:** Model-View-Controller (MVC)

### Additional Information:
- The server does not use HTTPS
- The default port (specified in main) is 8001
- Most endpoints besides "/" will return application/json



## Initializing the database
- Create a PostgreSQL instance at any port (this project uses 4323)
- A shell script ```setup.sh``` was provided and can be found in the database directory. Change the port number to your instance and run it
	- Note: this shell script uses zsh, if you use bash, modify it.



## Running the REST API
You may use ```go run main.go```. Alternatively, another shell script was provided and compiles the program into an executable and immediately runs the server.

<b>Note:</b> navigate into ```./rest-api/pkg/models/database.go``` and change your connection string if you are using a different port or database name. 


## API Endpoints

### [API Documentation Postman](https://documenter.getpostman.com/view/21072555/2s8YsxtqoY)<br>
Note: More Endpoints will be added.


## Useful Resources:
- [Go Rest API youtube](https://youtu.be/jFfo23yIWac)
- [Gorilla/Mux](https://github.com/gorilla/mux#middleware)
- [Go database/sql documentation](https://pkg.go.dev/database/sql)
- [GoJWT Documentation](https://pkg.go.dev/github.com/golang-jwt/jwt/v4)
- [Storing currency PostgresSQL](https://stackoverflow.com/questions/15726535/which-datatype-should-be-used-for-currency)
- [URL Params](https://stackoverflow.com/questions/46045756/retrieve-optional-query-variables-with-gorilla-mux)
- [Middleware](https://www.turing.com/kb/building-middleware-for-node-js)
- [Sample data generator](https://www.mockaroo.com)