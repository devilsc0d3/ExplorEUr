package database

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
