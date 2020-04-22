package main

import (
	"log"

	"github.com/whizzes/gabble/server/src/chat"
)

func main() {
	server, err := chat.MakeServer(":8000")

	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
