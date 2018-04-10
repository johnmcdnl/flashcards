package flashcards

import (
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type ID string

func newID() ID{
	logrus.Debugln(`func newID() ID{`)
	return ID(string(uuid.Must(uuid.NewV4()).String()[:4]))
}

func (id *ID) String() string {
	return jsonString(id)
}
