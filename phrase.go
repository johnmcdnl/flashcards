package flashcards

type Phrase struct {
	Translations []*Translation `json:"translations,omitempty"`
}

func NewPhrase() *Phrase {
	return &Phrase{}
}
