package database

import "regexp"

// TODO register

func IfNicknameExist(nickname string) bool {
	id, err := GetIDByNickname(nickname)
	if err != nil {
		panic(err)
	}
	if id == -1 {
		return false
	}
	return true
}

func IfEmailExist(email string) bool {
	id, err := GetIDByEmail(email)
	if err != nil {
		panic(err)
	}
	if id == -1 {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	re := regexp.MustCompile(`^[0-9a-z!#$%&'*+–/=?^_.{|}~]{1,64}@[a-z]{1,63}\.[a-z]{1,20}$`)
	return re.MatchString(email)
}
