package flashcards

import uuid "github.com/satori/go.uuid"

type ID string

func NewID() ID {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return ID(u.String())
}
