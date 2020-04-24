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
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		server.socket.ServeHTTP(w, r)
	})

	log.Println("Serving at http://127.0.0.1" + server.port)
	log.Fatal(http.ListenAndServe(server.port, nil))
}

// setEventListeners create setups the SocketIOServer events
func (server *SocketIOServer) setEventListeners() {
	server.socket.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		log.Printf("Connection Opened with ID: %s", so.ID())

		return nil
	})

	server.socket.OnEvent("/", NOTICE, func(so socketio.Conn, msg string) {
		so.Emit("reply", "have "+msg)
		log.Printf("Event: [%s]\t: %v\n", NOTICE, msg)
	})

	server.socket.OnEvent("/", MESSAGE, func(so socketio.Conn, msg string) string {
		log.Printf("Event: [%s]\tFrom: %s\tBody: %s\n", MESSAGE, so.ID(), msg)

		return msg
	})

	server.socket.OnEvent("/", BYE, func(so socketio.Conn) string {
		last := so.Context().(string)
		so.Emit(BYE, last)
		so.Close()

		log.Printf("Event: [%s]\t: Connection closed for: %s\n", BYE, last)

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
