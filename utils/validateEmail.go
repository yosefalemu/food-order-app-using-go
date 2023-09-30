package utils

import "regexp"

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	matched, err := regexp.MatchString(pattern, email)
	return err == nil && matched
}
