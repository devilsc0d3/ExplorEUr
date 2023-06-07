package database

import (
	"regexp"
)

// TODO register

func CheckEmail(str string) bool {
	re := regexp.MustCompile(`^[0-9a-z!#$%&'*+–/=?^_.{|}~]{1,64}@[a-z]{1,63}\.[a-z]{1,3}$`)
	return re.MatchString(str)
}
