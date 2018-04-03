package flashcards

type Card struct {
	Phrase *Phrase `json:"phrase,omitempty"`
}

func NewCard() *Card {
	return &Card{}
}
