package flashcards

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Card struct {
	Phrase *Phrase `json:"phrase,omitempty"`
}

func ShuffleCards(cards []*Card) {
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

func NewCard(p *Phrase) *Card {
	return &Card{
		Phrase: p,
	}
}

func (c *Card) PrintQuestion(deck *Deck) {
	if deck.Know == "" || deck.Learning == "" {
		panic(fmt.Sprintf("Something isn't set -- Know: %s -- Learning: %s", deck.Know, deck.Learning))
	}
	var options []*Card
	options = append(options, deck.GetIncorrectGuesses(c, 3)...)
	options = append(options, c)
	ShuffleCards(options)

	learning := c.Phrase.Language(deck.Learning)

	msg := fmt.Sprintf("Translate: %s  %s ", learning.Value, learning.GetPhonetic(deck.Know).Value)
	for i, o := range options {
		msg = fmt.Sprintf("%s \n \t %d) %s", msg, i+1, o.Phrase.Language(deck.Know).Value)
	}

	fmt.Println(msg)
}

func (c *Card) AttemptAnswer(know, learning Language, attempt string) {

	if attempt == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Type your answer: ")
		attempt, _ = reader.ReadString('\n')
		attempt = strings.Trim(attempt, " ")
		attempt = strings.Trim(attempt, "\n")
	}

	correctValue := c.Phrase.Language(know).Value

	if strings.EqualFold(correctValue, attempt) {
		fmt.Println("Correct!")
		c.Phrase.Language(learning).Stats.CorrectAttempt()
	} else {
		fmt.Println("Incorrect!! ", correctValue)
		c.Phrase.Language(learning).Stats.WrongAttempt()
	}

}
