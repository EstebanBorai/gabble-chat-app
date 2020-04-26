package chat

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

// SocketIOServer encapsulates functionality for
// running a ScoketIO server
type SocketIOServer struct {
	port              string
	host              string
	DefaultSocketPath string
	Socket            *socketio.Server
}

// SocketIOConfig represents the required
// configuration to build a SocketIOServer
// instance
type SocketIOConfig interface {
	GetHost() string
	GetPort() string
	GetDefaultSocketPath() string
}

// MakeServer creates a new SocketIOServer
// and returns it
func MakeServer(conf SocketIOConfig) (*SocketIOServer, error) {
	socket, err := socketio.NewServer(nil)

	if err != nil {
		return nil, err
	}

	server := new(SocketIOServer)

	server.port = conf.GetPort()
	server.host = conf.GetHost()
	server.Socket = socket
	server.DefaultSocketPath = conf.GetDefaultSocketPath()
	// server.setEventListeners()

	return server, nil
}

// Start serves the SockeIO server
func (server *SocketIOServer) Start() {
	go server.Socket.Serve()
	defer server.Socket.Close()

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		server.Socket.ServeHTTP(w, r)
	})
	host := server.host + ":" + server.port
	log.Println("Serving at http://" + host)
	log.Fatal(http.ListenAndServe(host, nil))
}

// setEventListeners create setups the SocketIOServer events
func (server *SocketIOServer) setEventListeners() {
	InitBasicEvents(server)
	InitHealtEvents(server)
}
