package administrator

import (
	"exploreur/backend/register"
)

func AddModeratorForAdmin(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("moderator", nicknameUser)
	}
}

func DeleteModeratorForAdmin(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("user", nicknameUser)
	}
}

func BanUserForAdmin(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("", nicknameUser)
	}
}

func UnbanUserForAdmin(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("user", nicknameUser)
	}
}
