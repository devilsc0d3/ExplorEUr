package fill

import (
	"exploreur/backend/database/category"
	"exploreur/backend/database/comment"
	"exploreur/backend/database/like_comment"
	"exploreur/backend/like_post"
	"exploreur/backend/post"
	"exploreur/backend/register"
	"exploreur/backend/report_post"
	"exploreur/backend/server"
)

func DataFill() {
	//post.Clear()
	//register.AddUserController("jean", "jean.eude@hotmial.com", "Passw0rd.")
	//register.AddUserController("adan", "adan135@gmail.fr", "Passw0rd.")
	//post.AddPost("L'eau c'est mieux en cannette, et bas non !", 0)
	//post.AddPost("La terre est plate!", 0)
	//post.AddPost("L'eau c'est de l'alcool", 1)
	//post.AddPost("jesus transforme l'eau en vin", 1)
	//comment.Clear()
	//category.AddCategory("Place")
	//category.AddCategory("Tools")
	//category.AddCategory("Information")
	//category.Clear()
	register.Init()
	post.Init()
	comment.Init()
	like_post.Init()
	like_comment.Init()
	report_post.Init()
	category.Init()
	server.Server()
}
