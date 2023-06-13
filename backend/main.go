package main

import (
	"exploreur/backend/registerDB"
	"exploreur/backend/server"
)

func main() {
	registerDB.ResetDatabase()
	registerDB.Init()
	server.Server()
}
