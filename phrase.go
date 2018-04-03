package flashcards

type Phrase struct {
	Value        string         `json:"value,omitempty"`
	Translations []*Translation `json:"translations,omitempty"`
}

func NewPhrase() *Phrase {
	return &Phrase{}
}

func (p *Phrase) WithTranslation(t *Translation) *Phrase {
	p.Translations = append(p.Translations, t)
	return p
}
