package chat

import (
	socketio "github.com/googollee/go-socket.io"
)

const (
	// PingEvent Ping request event
	PingEvent string = "ping"
)

// InitHealtEvents Basic health events
func InitHealtEvents(server *SocketIOServer) {
	server.Socket.OnEvent(server.DefaultSocketPath, PingEvent, func(so socketio.Conn, msg string) {
		so.Emit(PingEvent, msg)
	})
}
