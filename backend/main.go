package main

import (
	"exploreur/backend/database"
	"exploreur/backend/server"
	"fmt"
)

func main() {
	fmt.Println("test0")
	database.Init()
	fmt.Println("test1")
	server.Server()
}
