package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/johnmcdnl/flashcards"
)

var deck *flashcards.Deck

func init() {
	deck = flashcards.NewDeckWithSize("deck.db", 100, 125, false)
	deck.Know = flashcards.English
	deck.Learning = flashcards.Russian
}

func NextCard(w http.ResponseWriter, r *http.Request) {
	if err := render.Render(w, r, NewCardResponse(deck.Next())); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

type CardResponse struct {
	*flashcards.Card
}

func NewCardResponse(Card *flashcards.Card) *CardResponse {
	return &CardResponse{Card: Card}
}

func (rd *CardResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
