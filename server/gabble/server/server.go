package server

import (
	"fmt"
	"net/http"

	"github.com/whizzes/gabble/server/gabble/logger"
)

// Config represents the Gabble Server configuration
type Config interface {
	GetHost() string
	GetPort() string
}

//Server represents a Gabble Server
type Server struct {
	httpServer *http.Server
	addr       string
	logs       *logger.Logger
}

// NewServer creates a new Gabble Server instance
func NewServer(conf Config, chatHander http.HandlerFunc, logs *logger.Logger) (*Server, error) {
	gabbleServer := new(Server)

	if logs != nil {
		gabbleServer.logs = logs
	}

	gabbleServer.addr = conf.GetHost() + conf.GetPort()

	serv := new(http.Server)
	serv.Addr = gabbleServer.addr
	serv.Handler = chatHander

	gabbleServer.httpServer = serv

	return gabbleServer, nil
}

// Listen serves the HTTP Gabble server
func (s *Server) Listen() {
	if s.logs != nil {
		s.logs.Info(fmt.Sprintf("Listening at %s", s.addr))
	}

	s.httpServer.ListenAndServe()
}
