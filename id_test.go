package flashcards_test

import (
	"testing"

	"github.com/johnmcdnl/flashcards"
	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	assert.NotEmpty(t, flashcards.NewID())
}

func TestNewID_IsUnique(t *testing.T) {
	assert.NotEqual(t, flashcards.NewID(), flashcards.NewID())
}
