package check

import "regexp"

func IsEmailValid(e string) bool {
	emailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegexp.MatchString(e)
}

