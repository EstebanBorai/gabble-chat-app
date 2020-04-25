package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/whizzes/gabble/server/src/chat"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverHost, serverPort := os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")

	if serverHost == "" || serverPort == "" {
		log.Fatal("SERVER_HOST or SERVER_PORT shouldnt be empty")
	}

	server, err := chat.MakeServer(serverHost, serverPort)

	if err != nil {
		fmt.Printf("socket new server err: %v\n", err)
		return
	}

	server.Start()
}
