package flashcards_test

import (
	"testing"

	"github.com/johnmcdnl/flashcards"
	"github.com/stretchr/testify/assert"
)

func TestDeckNext(t *testing.T) {
	deck := flashcards.NewStandardDeck()

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
