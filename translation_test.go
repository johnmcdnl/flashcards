package flashcards_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/johnmcdnl/flashcards"
)

func TestNewTranslation(t *testing.T) {
	var engHello = flashcards.NewTranslation(flashcards.Eng, "hello")
	assert.Equal(t, flashcards.Eng, engHello.Language)
	assert.Equal(t, "hello", engHello.Value)

	var rusHello = flashcards.NewTranslation(flashcards.Rus, "Здравствуйте")
	assert.Equal(t, flashcards.Rus, rusHello.Language)
	assert.Equal(t, "Здравствуйте", rusHello.Value)
}
