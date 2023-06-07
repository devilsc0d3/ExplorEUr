package database

func AddUserController(nickname string, email string, password string) {
	IfNicknameExist(nickname)
	IfEmailExist(email)
	//check password
	//check email
	AddUser(nickname, email, password, "user")
}

func UpdateUserRoleController(role string, nickname string) {
	IfNicknameExist(nickname)
	id, _ := GetIDByNickname(nickname)
	UpdateUserRole(role, id)
}

func UpdateUserNicknameController(nickname string) {
	IfNicknameExist(nickname)
	id, _ := GetIDByNickname(nickname)
	UpdateUserNickname(nickname, id)
}

func UpdateUserPasswordController() {
	//check password

}
