package main

import (
	"exploreur/backend/register"
	"exploreur/backend/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}
	register.Init()
	server.Server()
}
