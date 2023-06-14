package main

import (
	"exploreur/backend/fill"
	"exploreur/backend/post"
	"exploreur/backend/register"
	"exploreur/backend/server"
)

func main() {
	//fmt.Println("test0")
	register.Init()
	post.Init()
	fill.DataFill()
	//fmt.Println("test1")
	server.Server()
}
