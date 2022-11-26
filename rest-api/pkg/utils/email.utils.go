package	utilities 

import (
	"strings"
)


func IsEmail(str string) bool {
	return strings.Contains(str, "@")
}