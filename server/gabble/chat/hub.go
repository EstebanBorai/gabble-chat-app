package chat

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// HubStart creates a new chat hub and
// start listening
func HubStart() {
	hub := newHub()
	go hub.await()
}

func newHub() *Hub {
	var hub *Hub = new(Hub)

	hub.broadcast = make(chan []byte)
	hub.register = make(chan *Client)
	hub.unregister = make(chan *Client)
	hub.clients = make(map[*Client]bool)

	return hub
}

func (h *Hub) await() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
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
		}
	}
}
