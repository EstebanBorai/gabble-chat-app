package chat

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/whizzes/gabble/server/gabble/logger"
)

type Chat struct {
	hub      *Hub
	upgrader *websocket.Upgrader
	logs     *logger.Logger
}

// NewChat creates a new Chat instance
func NewChat(conf Config, logs *logger.Logger) (*Chat, *Hub, error) {
	var chat *Chat = new(Chat)

	if logs != nil {
		chat.logs = logs
	}

	// Create upgrader
	chat.upgrader = makeSocketUpgrader(conf)
	chat.hub = newHub(logs)

	return chat, chat.hub, nil
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
		if c.logs != nil {
			c.logs.Info("Received request")
		}

		conn, err := c.upgrader.Upgrade(w, r, nil)

		if err != nil {
			if c.logs != nil {
				c.logs.Error(err)
			}

			panic(err)
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
