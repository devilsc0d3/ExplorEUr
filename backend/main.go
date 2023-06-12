package main

import (
	"exploreur/backend/server"
	"exploreur/backend/userDB"
)

func main() {
	userDB.Init()
	server.Server()
}
