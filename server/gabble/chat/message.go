package chat

import (
	"encoding/json"
	"time"
)

// Message represents a chat client
type Message struct {
	Author   Client    `json:"author"`
	Message  string    `json:"message"`
	IssuedAt time.Time `json:"issuedAt"`
}

// MakeMessage creates a new Client
func MakeMessage(from string) *Message {
	var message *Message = new(Message)

	message.Author = "HARDCODED"
	message.IssuedAt = time.Now()
	message.Message = from

	return message
}

// String converts a Message type into a JSON string
func (m *Message) String() (string, error) {
	str, err := json.Marshal(m)

	if err != nil {
		return "", err
	}

	return string(str), nil
}

// MustParseString converts a Message type into JSON string.
// If an error occurs it panics
func (m *Message) MustParseString() string {
	str, err := m.String()

	if err != nil {
		panic(err)
	}

	return str
}
