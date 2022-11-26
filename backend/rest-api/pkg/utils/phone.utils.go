package utilities

import (
	"regexp"
)

func IsPhoneNumber(phoneNumber string) bool {
	// Check if the phone number is valid
	// Format: +1-123-456-7890
	// Regex: ^\+1-\d{3}-\d{3}-\d{4}$
	re := regexp.MustCompile(`^\+1-\d{3}-\d{3}-\d{4}$`)
	return re.MatchString(phoneNumber)
}
