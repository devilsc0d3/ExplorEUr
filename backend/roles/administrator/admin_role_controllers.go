package administrator

import (
	"exploreur/backend/database/category"
	"exploreur/backend/database/comment"
	"exploreur/backend/post"
	"exploreur/backend/register"
)

func AddModeratorByAdminController(nicknameUser string) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	register.UpdateUserRoleController("moderator", nicknameUser)
	return ""
}

func DeleteModeratorByAdminController(nicknameUser string) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	register.UpdateUserRoleController("user", nicknameUser)
	return ""
}

func BanUserByAdminController(nicknameUser string) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	register.UpdateUserRoleController("", nicknameUser)
	return ""
}

func UnbanUserByAdminController(nicknameUser string) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	register.UpdateUserRoleController("user", nicknameUser)
	return ""
}

func DeletePostByAdminController(postID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	post.DeletePost(postID)
	return ""
}

func DeleteCommentByAdminController(commentID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	comment.DeleteComment(commentID)
	return ""
}

func AddCategoryByAdminController(categoryName string) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	category.AddCategory(categoryName)
	return ""
}

func UpdateCategoryByAdminController(newName string, categoryID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	category.UpdateCategoryName(newName, categoryID)
	return ""
}

func DeleteCategoryByAdminController(categoryID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	category.DeleteCategory(categoryID)
	return ""
}
