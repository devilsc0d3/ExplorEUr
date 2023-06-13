package main

import (
	"exploreur/backend/register"
	"exploreur/backend/server"
	"fmt"
)

func main() {
	fmt.Println("test0")
	register.Init()
	fmt.Println("test1")
	server.Server()
}
