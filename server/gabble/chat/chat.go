package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/whizzes/gabble/server/gabble/logger"
)

// Chat represents a chat session
type Chat struct {
	clients  map[*websocket.Conn]Client
	session  chan Message
	upgrader *websocket.Upgrader
	logs     *logger.Logger
}

// Config represents a chat config
type Config interface {
	GetHost() string
	GetPort() string
	GetClientHost() string
	GetClientPort() string
}

// New creates a new Chat instance
func New(conf Config, logs *logger.Logger) (*Chat, error) {
	var chat *Chat = new(Chat)

	if logs != nil {
		chat.logs = logs
	}

	chat.clients = make(map[*websocket.Conn]Client)
	chat.upgrader = makeSocketUpgrader(conf)
	chat.session = make(chan Message)

	go chat.listenMessages()

	return chat, nil
}

// GetHandler creates a new http request websocket upgrader
func (chat *Chat) GetHandler() http.HandlerFunc {
	return chat.makeRequestsHandler()
}

// Broadcast a message to all clients
func (chat *Chat) Broadcast(message *Message) {
	for client := range chat.clients {
		err := client.WriteJSON(message)

		if err != nil {
			if chat.logs != nil {
				chat.logs.Warning("Error writting message!")
				chat.logs.Error(err)
			}

			client.Close()
			delete(chat.clients, client)
		}
	}
}

func makeSocketUpgrader(conf Config) *websocket.Upgrader {
	upgrader := new(websocket.Upgrader)

	// allowedOrigin := "http://" + conf.GetClientHost() + ":" + conf.GetClientPort()

	upgrader.CheckOrigin = func(r *http.Request) bool {
		// TODO: Should compare the request origin with the allowed origin to
		// return true

		return true
	}

	return upgrader
}

// For every message that arrives to this session
// broadcast the message to clients
func (chat *Chat) listenMessages() {
	for {
		message := <-chat.session

		if chat.logs != nil {
			chat.logs.Info("Received message: " + message.MustParseString())
		}

		chat.Broadcast(&message)
	}
}

func (chat *Chat) makeRequestsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := chat.upgrader.Upgrade(w, r, nil)

		if err != nil {
			if chat.logs != nil {
				chat.logs.Error(err)
			}
			// should be handled someway
			panic(err)
		}

		defer ws.Close()

		newClient := NewClient()

		chat.clients[ws] = newClient

		if chat.logs != nil {
			chat.logs.Info("New client registered in session")
		}

		chat.Broadcast(MakeMessage("Client " + string(newClient) + " joined the chat session!"))

		for {
			var message Message

			err := ws.ReadJSON(&message)

			// An error ocurred reading a message from a client
			if err != nil {
				if chat.logs != nil {
					chat.logs.Warning("Removing client")
				}

				delete(chat.clients, ws)
				break
			}

			chat.session <- message
		}
	}
}
