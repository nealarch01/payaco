package authentication

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nealarch01/payaco/pkg/models"
	"github.com/nealarch01/payaco/pkg/utils"
)


func CreateToken(accountID int) string {
	// Create claims
	claims := jwt.MapClaims{
		"id":       accountID,
		"issuedAt": utilities.CurrentDateTime(),
	}
	// Add claims to token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

// Checks if a token is valid
func ValidateToken(tokenString string) bool {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error decoding the token")
		}
		return []byte("secret"), nil
	})
	// Check if an error occured
	if err != nil {
		return false
	}

	// If the token is invalid
	if !token.Valid { // If expired or not valid
		return false
	}
	// Lastly, do a lookup to check if the token is blacklisted

	connection := models.InitConnection()
	if connection == nil {
		return false
	}

	var count int = 0
	connection.QueryRow("SELECT COUNT(*) FROM blacklist WHERE token = $1", tokenString).Scan(&count)
	return count == 0
}

func BlacklistToken(tokenString string) error {
	connection := models.InitConnection()
	if connection == nil {
		return fmt.Errorf("database connection failed")
	}

	_, err := connection.Exec("INSERT INTO blacklist (token) VALUES ($1)", tokenString)
	if err != nil {
		return err
	}
	return nil
}

