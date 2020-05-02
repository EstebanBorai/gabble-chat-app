package chat

import "github.com/google/uuid"

// Client represents a chat client
type Client string

// NewClient creates a new Client
func NewClient() Client {
	id := uuid.New()

	return Client(id.String())
}
