package chat

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

// InitBasicEvents Listening primitives events
func InitBasicEvents(server *SocketIOServer) {
	server.Socket.OnConnect(server.DefaultSocketPath, func(so socketio.Conn) error {
		so.SetContext("")
		log.Printf("Connection Opened with ID: %s", so.ID())

		return nil
	})

	server.Socket.OnError(server.DefaultSocketPath, func(conn socketio.Conn, e error) {
		// If the error is different to EOF, print and stop the server
		if e.Error() != "EOF" {
			log.Fatal(e)
		}
	})

	server.Socket.OnDisconnect(server.DefaultSocketPath, func(conn socketio.Conn, reason string) {
		log.Printf("Closed connection ID: %s, reason: %s", conn.ID(), reason)
	})
}
