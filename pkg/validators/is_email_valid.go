package validators

import "regexp"

func isEmailValid(email string) bool {
	var regexEmail = regexp.MustCompile(".*@.*")
	if len(email) < 3 || len(email) > 64 || !regexEmail.MatchString(email) {
		return false
	}
	return true
}
