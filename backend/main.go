package main

import (
	"exploreur/backend/register"
	"exploreur/backend/server"
)

func main() {
	register.Init()
	server.Server()
}
