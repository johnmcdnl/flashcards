package flashcards

import uuid "github.com/satori/go.uuid"

// ID is unique identifier
type ID string

// NewID generates a new unique identifer
func NewID() ID {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return ID(u.String())
}
