package flashcards

// Phrase is a phrase from a languages that contains some translations in different languages
type Phrase struct {
	ID           ID             `json:"-"`
	Translations []*Translation `json:"translations,omitempty"`
}

// NewPhrase generates a new phrase
func NewPhrase() *Phrase {
	return &Phrase{ID: NewID()}
}

// WithTranslation adds a translation for a phrase
func (p *Phrase) WithTranslation(t *Translation) *Phrase {
	p.Translations = append(p.Translations, t)
	return p
}

// Language gets the translation for the phrase in said language
func (p *Phrase) Language(l Language) *Translation {
	for _, t := range p.Translations {
		if t.Language == l {
			return t
		}
	}
	return &Translation{}
}
