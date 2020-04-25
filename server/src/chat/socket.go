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
	host   string
}

// MakeServer creates a new SocketIOServer
// and returns it
func MakeServer(host string, port string) (*SocketIOServer, error) {
	if host == "" || port == "" {
		log.Fatal("SERVER_HOST or SERVER_PORT shouldnt be empty")
	}

	socket, err := socketio.NewServer(nil)

	if err != nil {
		return nil, err
	}

	server := new(SocketIOServer)

	server.port = port
	server.socket = socket
	server.host = host
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
	host := server.host + ":" + server.port
	log.Println("Serving at http://" + host)
	log.Fatal(http.ListenAndServe(host, nil))
}

// setEventListeners create setups the SocketIOServer events
func (server *SocketIOServer) setEventListeners() {
	chatEvents.InitBasicEvents(server.socket)
	chatEvents.InitHealtEvents(server.socket)
}
