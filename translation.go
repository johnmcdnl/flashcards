package flashcards

type Translation struct {
	Language  `json:"language,omitempty"`
	Value     string      `json:"value,omitempty"`
	Phonetics []*Phonetic `json:"phonetics,omitempty"`
}

func NewTranslation(l Language, value string) *Translation {
	return &Translation{
		Language: l,
		Value:    value,
	}
}

func (t *Translation) WithPhonetic(p *Phonetic) *Translation {
	t.Phonetics = append(t.Phonetics, p)
	return t
}
