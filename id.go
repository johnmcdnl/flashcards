package flashcards

import (
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// ID is a unique identifier
type ID string

func newID() ID{
	logrus.Debugln(`func newID() ID{`)
	return ID(uuid.Must(uuid.NewV4()).String())
}

// String represents a ID
func (id *ID) String() string {
	return jsonString(id)
}
