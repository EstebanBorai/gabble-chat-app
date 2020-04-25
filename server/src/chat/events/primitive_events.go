package chat

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	config "github.com/whizzes/gabble/server/src/chat/config"
)

// InitBasicEvents Listening primitives events
func InitBasicEvents(server *socketio.Server) {
	server.OnConnect(config.PATH, func(so socketio.Conn) error {
		so.SetContext("")
		log.Printf("Connection Opened with ID: %s", so.ID())

		return nil
	})

	server.OnError(config.PATH, func(conn socketio.Conn, e error) {
		// If the error is different to EOF, print and stop the server
		if e.Error() != "EOF" {
			log.Fatal(e)
		}
	})

	server.OnDisconnect(config.PATH, func(conn socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
}
