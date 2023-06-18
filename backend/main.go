package main

import (
	"exploreur/backend/database/comment"
	"exploreur/backend/fill"
	"exploreur/backend/post"
	"exploreur/backend/register"
	"exploreur/backend/server"
)

func main() {
	register.Init()
	post.Init()
	comment.Init()
	fill.DataFill()
	server.Server()
}
