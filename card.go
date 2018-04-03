package flashcards

type Card struct {
	Phrase *Phrase `json:"phrase,omitempty"`
}

func NewCard(p *Phrase) *Card {
	return &Card{
		Phrase: p,
	}
}
