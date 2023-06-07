package database

import "fmt"

func AddUserController(nickname string, email string, password string) {
	if !IfNicknameExist(nickname) {
		if !IfEmailExist(email) {
			if CheckPassword(password) {
				if CheckEmail(email) {
					AddUser(nickname, email, password, "user")
				} else {
					fmt.Println("the email is incorrect")
				}
			} else {
				fmt.Println("the password is incorrect") //+bonne pratique
			}
		} else {
			fmt.Println("the email already exist")
		}
	} else {
		fmt.Println("the nickname already exist")
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

func UpdateUserPasswordController(nickname string, password string) {
	if CheckPassword(password) {
		id, _ := GetIDByNickname(nickname)
		UpdateUserPassword(password, id)
	} else {
		fmt.Println("the password is incorrect") //+bonne pratique
	}

}
