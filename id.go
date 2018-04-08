package flashcards

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// ID is unique identifier
type ID string

// NewID generates a new unique identifer
func NewID() ID {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String())
}

// String is a string representation of the ID
func (id *ID) String() string {
	if id == nil {
		return ""
	}
	return string(*id)
}
