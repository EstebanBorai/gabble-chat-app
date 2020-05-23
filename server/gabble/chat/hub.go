package chat

import (
	"log"

	"github.com/whizzes/gabble/server/gabble/logger"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	logs       *logger.Logger
}

func newHub(logs *logger.Logger) *Hub {
	var hub *Hub = new(Hub)

	hub.broadcast = make(chan []byte)
	hub.register = make(chan *Client)
	hub.unregister = make(chan *Client)
	hub.clients = make(map[*Client]bool)
	hub.logs = logs

	return hub
}

// Await manage iteractions with the web socket
// This function must run concurrently with the server
func (h *Hub) Await() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

			if h.logs != nil {
				h.logs.Info("Client registered into hub")
				log.Println(client)
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

			if h.logs != nil {
				h.logs.Info("Client unregistered from hub")
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

			if h.logs != nil {
				h.logs.Info("Received Message")
			}
		}
	}
}
