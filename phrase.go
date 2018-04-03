package flashcards

type Phrase struct {
	ID           ID             `json:"-"`
	Translations []*Translation `json:"translations,omitempty"`
}

func NewPhrase() *Phrase {
	return &Phrase{ID: NewID()}
}

func (p *Phrase) WithTranslation(t *Translation) *Phrase {
	p.Translations = append(p.Translations, t)
	return p
}

func (p *Phrase) Language(l Language) *Translation {
	for _, t := range p.Translations {
		if t.Language == l {
			return t
		}
	}
	return &Translation{}
}
