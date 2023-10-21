package auth

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	// Паттерн для проверки формата email

	pattern := `^(|[\w.%+-]+@[\w.-]+\.[a-zA-Z]{2,})$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func (ctrl Controller) alreadyEmail(email string) (string, bool) {

	_, err := ctrl.studentService.GetLogin(email)
	if err == nil {
		return "student", true
	}

	_, err = ctrl.teacherService.GetLogin(email)
	if err == nil {
		return "teacher", true
	}

	_, err = ctrl.adminService.GetLogin(email)
	if err == nil {
		return "admin", true
	}

	return "", false

}
