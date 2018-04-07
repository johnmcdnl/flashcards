package flashcards

// Translation represents how a specific phrase is representated in a given language
type Translation struct {
	ID        ID `json:"-"`
	Language  `json:"language,omitempty"`
	Value     string      `json:"value,omitempty"`
	Phonetics []*Phonetic `json:"phonetics,omitempty"`
	Stats     *Stats      `json:"stats,omitempty"`
}

// NewTranslation generates a new translation
func NewTranslation(l Language, value string) *Translation {
	return &Translation{
		ID:       NewID(),
		Language: l,
		Value:    value,
		Stats:    NewStats(),
	}
}

// WithPhonetic adds a phonetic for a translation
func (t *Translation) WithPhonetic(p *Phonetic) *Translation {
	t.Phonetics = append(t.Phonetics, p)
	return t
}

// GetPhonetic gets the phonetic for a given lanuage
func (t *Translation) GetPhonetic(l Language) *Phonetic {
	for _, p := range t.Phonetics {
		if p.Language == l {
			return p
		}
	}
	return &Phonetic{}
}
