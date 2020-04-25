package chat

import (
	socketio "github.com/googollee/go-socket.io"
	config "github.com/whizzes/gabble/server/src/chat/config"
)

const (
	// PingEvent Ping request event
	PingEvent string = "ping"
)

// InitHealtEvents Basic health events
func InitHealtEvents(server *socketio.Server) {
	server.OnEvent(config.PATH, PingEvent, func(so socketio.Conn, msg string) {
		so.Emit(PingEvent, msg)
	})
}
