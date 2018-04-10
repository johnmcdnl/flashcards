package flashcards

import "github.com/satori/go.uuid"

type ID string

func (id *ID) String() string {
	return jsonString(id)
}

func newID() ID{
	return ID(string(uuid.Must(uuid.NewV4()).String()[:4]))
}