package chat

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	config "github.com/whizzes/gabble/server/src/chat/config"
)

const (
	// PingEvent Ping request event
	PingEvent string = "ping"
)

// InitHealtEvents Basic health events
func InitHealtEvents(server *socketio.Server) {
	fmt.Println(PingEvent)
	server.OnEvent(config.PATH, PingEvent, func(so socketio.Conn, msg string) {
		so.Emit(PingEvent, msg)
	})
}
