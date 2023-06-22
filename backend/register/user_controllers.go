package register

func AddUserController(nickname string, email string, password string) string {
	if IfNicknameExist(nickname) {
		return "the nickname already exist"
	}
	if IfEmailExist(email) {
		return "the email already exist"
	}
	if !CheckPassword(password) {
		return "the password is incorrect"
	}
	if !CheckEmail(email) {
		return "the email is incorrect"
	}
	AddUser(nickname, email, password, "user")
	return ""
}

func UpdateUserRoleController(role string, nickname string) string {
	if IfNicknameExist(nickname) {
		return "the nickname already exist"
	}
	id, _ := GetIDByNickname(nickname)
	UpdateUserRole(role, id)
	return ""
}

func UpdateUserNicknameController(nickname string, newNickname string) string {
	if IfNicknameExist(newNickname) {
		return "the nickname already exist"
	}
	id, _ := GetIDByNickname(nickname)
	UpdateUserNickname(newNickname, id)
	return ""
}

func UpdateUserPasswordController(nickname string, password string) string {
	if !CheckPassword(password) {
		return "the password is incorrect"
	}
	id, _ := GetIDByNickname(nickname)
	UpdateUserPassword(password, id)
	return ""
}
