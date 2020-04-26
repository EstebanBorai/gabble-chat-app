package gabble

import (
	"github.com/whizzes/gabble/server/gabble/chat"
	"github.com/whizzes/gabble/server/gabble/config"
)

// Server manages socket connections and
// server side state of a gabble instance
type Server struct {
	sockets *chat.SocketIOServer
}

// NewGabble creates a new GabbleServer
// instance
func NewGabble(conf config.Config) (*Server, error) {
	var server *Server = new(Server)

	sockets, err := chat.MakeServer(conf)

	if err != nil {
		return nil, err
	}

	server.sockets = sockets

	return server, nil
}

// Run initializes a Gabble server
func (s *Server) Run() {
	s.sockets.Start()
}
