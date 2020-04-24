package chat

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	chatEvents "github.com/whizzes/gabble/server/src/chat/events"
)

const (
	// NOTICE represents the "notice" event
	NOTICE string = "notice"
	// MESSAGE represents the "message" event
	MESSAGE string = "message"
	// BYE represents the "bye" event
	BYE string = "bye"
)

// SocketIOServer encapsulates functionality for
// running a ScoketIO server
type SocketIOServer struct {
	socket *socketio.Server
	port   string
}

// MakeServer creates a new SocketIOServer
// and returns it
func MakeServer(port string) (*SocketIOServer, error) {
	socket, err := socketio.NewServer(nil)

	if err != nil {
		return nil, err
	}

	server := new(SocketIOServer)

	server.port = port
	server.socket = socket
	server.setEventListeners()

	return server, nil
}

// Start serves the SockeIO server
func (server *SocketIOServer) Start() {
	go server.socket.Serve()
	defer server.socket.Close()

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		server.socket.ServeHTTP(w, r)
	})

	log.Println("Serving at http://0.0.0.0" + server.port)
	log.Fatal(http.ListenAndServe(server.port, nil))
}

// setEventListeners create setups the SocketIOServer events
func (server *SocketIOServer) setEventListeners() {
	chatEvents.InitBasicEvents(server.socket)
	chatEvents.InitHealtEvents(server.socket)
}
