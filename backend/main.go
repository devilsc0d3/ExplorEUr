package main

import (
	"exploreur/backend/registerDB"
	"exploreur/backend/server"
)

func main() {
	registerDB.Init()
	server.Server()
}
