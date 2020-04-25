package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/whizzes/gabble/server/src/chat"
)

func main() {
	godotenv.Load("../.env")

	server, err := chat.MakeServer(os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
