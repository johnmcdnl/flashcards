package flashcards

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Card is a single card in the deck
type Card struct {
	ID     ID      `json:"-"`
	Phrase *Phrase `json:"phrase,omitempty"`
}

// ShuffleCards reorders items in the slice of cards using a PRNG
func ShuffleCards(cards []*Card) {
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

// NewCard generates a new card to represdent a single phrase
func NewCard(p *Phrase) *Card {
	return &Card{
		ID:     NewID(),
		Phrase: p,
	}
}

// PrintQuestion is a pretty formatter for command line version of rendering the card
func (c *Card) PrintQuestion(deck *Deck) {
	if deck.Know == "" || deck.Learning == "" {
		panic(fmt.Sprintf("Something isn't set -- Know: %s -- Learning: %s", deck.Know, deck.Learning))
	}
	var options []*Card
	options = append(options, deck.GetIncorrectGuesses(c, 3)...)
	options = append(options, c)
	ShuffleCards(options)

	learning := c.Phrase.Language(deck.Learning)
	percentage := learning.Stats.Percentage * 100
	weighting := learning.Stats.Weighting
	msg := fmt.Sprintf("Translate: %s  %s     (%.1f %s   %.1f)",
		learning.Value, learning.GetPhonetic(deck.Know).Value, percentage, "%", weighting)
	for i, o := range options {
		msg = fmt.Sprintf("%s \n \t %d) %s", msg, i+1, o.Phrase.Language(deck.Know).Value)
	}

	fmt.Println(msg)
}

// AttemptAnswer accepts an answer from a user and checks whether it is correct
func (c *Card) AttemptAnswer(know, learning Language, attempt string) {

	if attempt == "" {
		attempt = readString()
	}

	phrase := c.Phrase

	knowT := phrase.Language(know)
	learningT := phrase.Language(learning)

	if strings.EqualFold(attempt, knowT.Value) {
		c.Phrase.Language(learning).Stats.CorrectAttempt()
	} else {
		c.Phrase.Language(learning).Stats.WrongAttempt()
		logrus.Errorf("\t | \t %s \t | \t %s \t | \t %s \t | \t %s",
			attempt,
			learningT.Value,
			learningT.GetPhonetic(know).Value,
			knowT.Value,
		)
	}

}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your answer: ")
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, " ")
	input = strings.Trim(input, "\n")
	return input
}
