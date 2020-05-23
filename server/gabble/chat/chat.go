package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/estebanborai/whizzes-gabble/server/gabble/logger"
)

type Chat struct {
	hub      *Hub
	upgrader *websocket.Upgrader
	logs     *logger.Logger
}

// NewChat creates a new Chat instance
func NewChat(conf Config, logs *logger.Logger) (*Chat, error) {
	var chat *Chat = new(Chat)

	if logs != nil {
		chat.logs = logs
	}

	// Create upgrader
	chat.upgrader = makeSocketUpgrader(conf)
	chat.hub = newHub()

	return chat, nil
}

// Start a Chat instance
func (c *Chat) Start() {
	HubStart()
}

func (c *Chat) GetHandler() http.HandlerFunc {
	return c.makeRequestsHandler()
}

func makeSocketUpgrader(conf Config) *websocket.Upgrader {
	upgrader := new(websocket.Upgrader)

	upgrader.WriteBufferSize = 1024
	upgrader.ReadBufferSize = 1024

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	return upgrader
}

func (c *Chat) makeRequestsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := c.upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{
			hub:  c.hub,
			conn: conn,
			send: make(chan []byte, 256),
		}

		client.hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}
