package flashcards

import uuid "github.com/satori/go.uuid"

// ID is unique identifier
type ID string

// NewID generates a new unique identifer
func NewID() ID {
	return ID(uuid.NewV4().String())
}
