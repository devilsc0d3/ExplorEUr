package user

import (
	"regexp"
)

func VerifUserPassword(password string) bool {

	verif1, _ := regexp.MatchString(`[^\w]`, password)
	verif2, _ := regexp.MatchString(`[\w]`, password)

	if verif1 && verif2 {
		return true
	} else {
		return false
	}

}
