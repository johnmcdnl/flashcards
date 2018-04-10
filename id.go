package flashcards

import "github.com/satori/go.uuid"

type ID string

func newID() ID{
	return ID(uuid.Must(uuid.NewV4()).String())
}