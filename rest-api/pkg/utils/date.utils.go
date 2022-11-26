package utilities 

import (
	"time"
)

// Returns the current date and time
func CurrentDateTime() string {
	return time.Now().Format("2022-01-01 12:00:00")
}