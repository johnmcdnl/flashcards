package flashcards_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnmcdnl/flashcards"
)

func TestShuffleCards(t *testing.T) {
	var cards []*flashcards.Card
	for i := 0; i <= 10; i++ {
		c := flashcards.NewCard(flashcards.NewPhrase())
		c.ID = flashcards.ID(fmt.Sprintf("c%d", i))
		cards = append(cards, c)
	}

	flashcards.ShuffleCards(cards)

	assert.True(t,
		cards[0].ID.String() != "c0" ||
			cards[1].ID.String() != "c1" ||
			cards[2].ID.String() != "c2" ||
			cards[3].ID.String() != "c3" ||
			cards[4].ID.String() != "c4" ||
			cards[5].ID.String() != "c5" ||
			cards[6].ID.String() != "c6" ||
			cards[7].ID.String() != "c7" ||
			cards[8].ID.String() != "c8" ||
			cards[9].ID.String() != "c9",
	)
}

func TestPrintQuestion(t *testing.T) {
	d := NewTestDeck("TestPrintQuestion")
	c := d.Next()
	assert.NotEmpty(t, c.PrintQuestion(d))
}

func TestAttemptAnswer_Correct(t *testing.T) {
	d := NewTestDeck("TestAttemptAnswer_Correct")
	c := d.Next()
	correctAnswer := c.Phrase.Language(d.Know).Value

	origAttempt := c.Phrase.Language(d.Learning).Stats.Attempts
	origCorrect := c.Phrase.Language(d.Learning).Stats.Correct
	origPercentage := c.Phrase.Language(d.Learning).Stats.Percentage
	c.AttemptAnswer(d.Know, d.Learning, correctAnswer)
	afterAttempt := c.Phrase.Language(d.Learning).Stats.Attempts
	afterCorrect := c.Phrase.Language(d.Learning).Stats.Correct
	afterPercentage := c.Phrase.Language(d.Learning).Stats.Percentage

	assert.Equal(t, origAttempt+1, afterAttempt)
	assert.Equal(t, origCorrect+1, afterCorrect)
	assert.True(t, afterPercentage > origPercentage)
}

func TestAttemptAnswer_Incorrect(t *testing.T) {
	d := NewTestDeck("TestAttemptAnswer_Incorrect")
	c := d.Next()

	correctAnswer := c.Phrase.Language(d.Know).Value
	origAttempt := c.Phrase.Language(d.Learning).Stats.Attempts
	origCorrect := c.Phrase.Language(d.Learning).Stats.Correct
	origPercentage := c.Phrase.Language(d.Learning).Stats.Percentage
	c.AttemptAnswer(d.Know, d.Learning, correctAnswer+"ABC")
	afterAttempt := c.Phrase.Language(d.Learning).Stats.Attempts
	afterCorrect := c.Phrase.Language(d.Learning).Stats.Correct
	afterPercentage := c.Phrase.Language(d.Learning).Stats.Percentage

	assert.Equal(t, origAttempt+1, afterAttempt)
	assert.Equal(t, origCorrect, afterCorrect)
	if origPercentage > 0 {
		assert.True(t, afterPercentage < origPercentage)
	} else {
		assert.True(t, afterPercentage == origPercentage)
	}

}
