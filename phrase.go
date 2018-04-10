package flashcards

type Phrase struct {
	ID                          `json:"id"`
	Translations []*Translation `json:"translations"`
}

func NewPhrase(t *Translation) *Phrase {
	var p Phrase
	p.ID = newID()
	return p.With(t)
}

func (p *Phrase) String() string {
	return jsonString(p)
}

func (p *Phrase) With(t *Translation) *Phrase {
	if t == nil {
		return p
	}
	p.Translations = append(p.Translations, t)
	return p
}
