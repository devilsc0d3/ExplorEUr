package main

import (
	"exploreur/backend/register"
	"exploreur/backend/server"
)

func main() {
	register.Init()
	register.AddUser("Sasha", "aujzf@azga.fazgf", "1234", "admin")
	server.Server()
}
