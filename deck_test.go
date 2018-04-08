package flashcards_test

import (
	"os"
	"testing"

	"github.com/johnmcdnl/flashcards"
	"github.com/stretchr/testify/assert"
)

const testDeckName = "flashcardstest"

func NewTestDeck(unique string) *flashcards.Deck {
	os.Remove(testDeckName + unique + ".db")
	return flashcards.NewDeck(testDeckName + unique + ".db")
}

func TestDeckNext(t *testing.T) {
	deck := NewTestDeck("TestDeckNext")

	assert.Nil(t, deck.Current)
	assert.Nil(t, deck.Last)

	card1 := deck.Next()
	assert.NotNil(t, card1)
	assert.Equal(t, card1, deck.Current)
	assert.Nil(t, deck.Last)

	card2 := deck.Next()
	assert.NotNil(t, card2)
	assert.Equal(t, card2, deck.Current)
	assert.Equal(t, card1, deck.Last)
}
