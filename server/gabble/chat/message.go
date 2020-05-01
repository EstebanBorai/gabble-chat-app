package chat

// Message represents a chat client
type Message string

// MakeMessage creates a new Client
func MakeMessage(from string) *Message {
	message := Message(from)
	return &message
}

func (m *Message) String() string {
	return string(*m)
}
