package administrator

import (
	"exploreur/backend/register"
)

func AddModeratorByAdminController(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("moderator", nicknameUser)
	}
}

func DeleteModeratorByAdminController(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("user", nicknameUser)
	}
}

func BanUserByAdminController(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("", nicknameUser)
	}
}

func UnbanUserByAdminController(nicknameUser string) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "admin" {
		register.UpdateUserRoleController("user", nicknameUser)
	}
}

func DeletePostByAdminController() {

}

func DeleteCommentByAdminController() {

}

func AddCategoryByAdminController() {

}

func UpdateCategoryByAdminController() {

}

func DeleteCategoryByAdminController() {

}
