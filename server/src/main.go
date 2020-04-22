package main

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		fmt.Printf("socket new server err: %v\n", err)
		return
	}

	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Printf("connected: %v\n", conn.ID())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(conn socketio.Conn, msg string) {
		fmt.Printf("closed, msg: %v\n", msg)
	})

	go server.Serve()
	defer server.Close()

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	})
	fmt.Println(http.ListenAndServe(":8000", nil))
}
