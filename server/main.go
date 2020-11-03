package main

import (
	"github.com/EstebanBorai/gabble-chat-app/server/gabble"
	"github.com/EstebanBorai/gabble-chat-app/server/gabble/config"
)

func main() {
	conf, err := config.FromEnvFile()

	if err != nil {
		panic(err)
	}

	server, err := gabble.New(conf)

	if err != nil {
		panic(err)
	}

	server.Run()
}
