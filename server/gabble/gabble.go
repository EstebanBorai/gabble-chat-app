package gabble

import (
	"github.com/whizzes/gabble/server/gabble/chat"
	"github.com/whizzes/gabble/server/gabble/config"
	"github.com/whizzes/gabble/server/gabble/logger"
	"github.com/whizzes/gabble/server/gabble/server"
)

// Gabble instance
type Gabble struct {
	chatHub *chat.Hub
	server  *server.Server
}

// New creates a new Gabble instance
func New(conf config.Config) (*Gabble, error) {
	var instance *Gabble = new(Gabble)
	var ch *chat.Chat
	var serv *server.Server
	var logs *logger.Logger

	if conf.GetLogLevel() != 0 {
		logs = logger.NewLogger(int8(conf.GetLogLevel()))
	}

	ch, hub, err := chat.NewChat(conf, logs)

	if err != nil {
		return nil, err
	}

	chatRequestHandler := ch.GetHandler()

	serv, err = server.NewServer(conf, chatRequestHandler, logs)

	if err != nil {
		return nil, err
	}

	instance.server = serv
	instance.chatHub = hub

	return instance, nil
}

// Run initializes a Gabble server
func (s *Gabble) Run() {
	go s.chatHub.Await()
	s.server.Listen()
}
