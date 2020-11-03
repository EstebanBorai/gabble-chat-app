package gabble

import (
	"github.com/EstebanBorai/gabble-chat-app/server/gabble/chat"
	"github.com/EstebanBorai/gabble-chat-app/server/gabble/config"
	"github.com/EstebanBorai/gabble-chat-app/server/gabble/logger"
	"github.com/EstebanBorai/gabble-chat-app/server/gabble/server"
)

// Gabble instance
type Gabble struct {
	server *server.Server
}

// New creates a new Gabble instance
func New(conf config.Config) (*Gabble, error) {
	var instance *Gabble = new(Gabble)
	var cht *chat.Chat
	var serv *server.Server
	var logs *logger.Logger

	if conf.GetLogLevel() != 0 {
		logs = logger.NewLogger(int8(conf.GetLogLevel()))
	}

	cht, err := chat.New(conf, logs)

	if err != nil {
		return nil, err
	}

	serv, err = server.NewServer(conf, cht.GetHandler(), logs)

	if err != nil {
		return nil, err
	}

	instance.server = serv

	return instance, nil
}

// Run initializes a Gabble server
func (s *Gabble) Run() {
	s.server.Listen()
}
