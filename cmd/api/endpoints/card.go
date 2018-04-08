package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/johnmcdnl/flashcards"
)

var deck *flashcards.Deck

func init() {
	deck = flashcards.NewDeckWithSize(flashcards.DeckName, 100, 125, false)
	deck.Know = flashcards.English
	deck.Learning = flashcards.Russian
}

// NextCard is a handler to get the next card
func NextCard(w http.ResponseWriter, r *http.Request) {
	if err := render.Render(w, r, NewCardResponse(deck.Next())); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// CardResponse is a caard
type CardResponse struct {
	*flashcards.Card
}

//NewCardResponse is a new card response
func NewCardResponse(Card *flashcards.Card) *CardResponse {
	return &CardResponse{Card: Card}
}

// Render renders the cardresponse
func (rd *CardResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
