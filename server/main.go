package main

import (
	"github.com/whizzes/gabble/server/gabble"
	"github.com/whizzes/gabble/server/gabble/config"
)

func main() {
	conf, err := config.FromEnv()

	if err != nil {
		panic(err)
	}

	server, err := gabble.NewGabble(conf)

	if err != nil {
		panic(err)
	}

	server.Run()
}
