package main

import (
	"exploreur/backend/database"
	"exploreur/backend/server"
)

func main() {
	database.Init()
	server.Server()
}
