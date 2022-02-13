package validators

import "regexp"

func IsEmailValid(email string) bool {
	var regexEmail = regexp.MustCompile(".*@.*")
	if len(email) < 3 || len(email) > 64 || !regexEmail.MatchString(email) {
		return false
	}
	return true
}
