package register

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

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

func CheckPassword(password string) bool {

	check1, _ := regexp.MatchString(`[^\w]`, password)
	check2, _ := regexp.MatchString(`[\w]`, password)
	check3 := len(password) >= 8

	if check1 && check2 && check3 {
		return true
	} else {
		return false
	}

}

func CheckNicknameAndPassword(nickname string, password string) (bool, *User) {
	if IfNicknameExist(nickname) {
		id, _ := GetIDByNickname(nickname)
		user := GetUserByID(id)
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err == nil {
			return true, user
		}
	}
	return false, nil
}
