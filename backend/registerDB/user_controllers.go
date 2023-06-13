package registerDB

import (
	"fmt"
)

func AddUserController(nickname string, email string, password string) string {
	if !IfNicknameExist(nickname) {
		if !IfEmailExist(email) {
			if CheckPassword(password) {
				if CheckEmail(email) {
					AddUser(nickname, email, password, "user")
					return ""
				} else {
					return "the email is incorrect"
				}
			} else {
				return "the password is incorrect"
			}
		} else {
			return "the email already exist"
		}
	} else {
		return "the nickname already exist"
	}
}

func UpdateUserRoleController(role string, nickname string) {
	if !IfNicknameExist(nickname) {
		id, _ := GetIDByNickname(nickname)
		UpdateUserRole(role, id)
	} else {
		fmt.Println("the nickname already exist")
	}
}

func UpdateUserNicknameController(nickname string, newNickname string) {
	if !IfNicknameExist(newNickname) {
		id, _ := GetIDByNickname(nickname)
		UpdateUserNickname(newNickname, id)
	} else {
		fmt.Println("the nickname already exist")
	}
}

func UpdateUserPasswordController(nickname string, password string) string {
	if CheckPassword(password) {
		id, _ := GetIDByNickname(nickname)
		UpdateUserPassword(password, id)
	} else {
		return "the password is incorrect"
	}
	return ""
}
