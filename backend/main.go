package main

import (
	"exploreur/backend/database/comment"
	"exploreur/backend/database/like_comment"
	"exploreur/backend/like_post"
	"exploreur/backend/post"
	"exploreur/backend/register"
	"exploreur/backend/report_post"
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
	comment.Init()
	like_post.Init()
	like_comment.Init()
	report_post.Init()
	server.Server()
}
