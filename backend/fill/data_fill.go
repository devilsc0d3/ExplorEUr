package fill

import (
	"exploreur/backend/database/category"
	"exploreur/backend/register"
)

func DataFill() {
	//comment.Clear()
	//post.Clear()
	//category.Clear()

	register.AddUserController("admin", "admin@admin.fr", "Viol1418.")
	register.AddUserController("jean", "jean.eude@hotmial.com", "Passw0rd.")
	register.AddUserController("adan", "adan135@gmail.fr", "Passw0rd.")

	//post.AddPost("L'eau c'est mieux en cannette, et bas non !", 0, 3)
	//post.AddPost("La terre est plate!", 2, 1)
	//post.AddPost("L'eau c'est de l'alcool", 1, 3)
	//post.AddPost("jesus transforme l'eau en vin", 1, 2)

	category.AddCategory("Place")
	category.AddCategory("Tools")
	category.AddCategory("Information")
}
