package main

import (
	"github.com/whizzes/gabble/server/gabble"
	"github.com/whizzes/gabble/server/gabble/config"
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
