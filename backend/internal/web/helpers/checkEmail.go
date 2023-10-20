package helpers

import "regexp"

func IsValidEmail(email string) bool {
	// Паттерн для проверки формата email

	pattern := `^(|[\w.%+-]+@[\w.-]+\.[a-zA-Z]{2,})$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}
