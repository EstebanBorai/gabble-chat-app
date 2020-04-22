package main

import (
	"fmt"

	"github.com/whizzes/gabble/server/src/chat"
)

func main() {
	server, err := chat.MakeServer(":8000")

	if err != nil {
		fmt.Printf("socket new server err: %v\n", err)
		return
	}

	server.Start()
}
