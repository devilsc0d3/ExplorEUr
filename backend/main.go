package main

import (
	"exploreur/backend/fill"
	"exploreur/backend/post"
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
	post.Init()
	fill.DataFill()
	//fmt.Println("test1")
	server.Server()
}
