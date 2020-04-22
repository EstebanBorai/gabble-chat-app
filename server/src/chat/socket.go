package chat

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
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
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		server.socket.ServeHTTP(w, r)
	})

	log.Println("Serving at http://127.0.0.1" + server.port)
	log.Fatal(http.ListenAndServe(server.port, nil))
}

// setEventListeners create setups the SocketIOServer events
func (server *SocketIOServer) setEventListeners() {
	server.socket.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		log.Printf("Connection Opened with ID: %s", conn.ID())

		return nil
	})

	server.socket.OnEvent("/", NOTICE, func(conn socketio.Conn, msg string) {
		conn.Emit("reply", "have "+msg)
		log.Printf("Event: %s\t: %v\n", NOTICE, msg)
	})

	server.socket.OnEvent("/chat", MESSAGE, func(conn socketio.Conn, msg string) string {
		conn.SetContext(msg)
		log.Printf("Event: %s\t: %v\n", MESSAGE, msg)

		return "Received " + msg
	})

	server.socket.OnEvent("/", BYE, func(conn socketio.Conn) string {
		last := conn.Context().(string)
		conn.Emit(BYE, last)
		conn.Close()

		log.Printf("Event: %s\t: Connection closed for: %s\n", BYE, last)

		return last
	})

	server.socket.OnError("/", func(conn socketio.Conn, e error) {
		// print the error message and stop
		log.Fatal(e)
	})

	server.socket.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
}
