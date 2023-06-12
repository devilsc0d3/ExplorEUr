package main

import (
	"exploreur/backend/register"
	"exploreur/backend/server"
)

func main() {
	register.Init()
	register.AddUser("Sasha", "sasha@sasha.sasha", "Sasha.63", "admin")
	register.CheckNicknameAndPassword("Sasha", "Sasha.63hg")
	server.Server()
}
