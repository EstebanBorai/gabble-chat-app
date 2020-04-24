package chat

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

// InitBasicEvents Listening primitives events
func InitBasicEvents(server *socketio.Server) {
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		log.Printf("Connection Opened with ID: %s", so.ID())

		return nil
	})

	server.OnError("/", func(conn socketio.Conn, e error) {
		// print the error message and stop
		log.Fatal(e)
	})

	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
}
