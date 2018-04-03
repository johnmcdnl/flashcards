package flashcards

type Translation struct {
	ID        ID `json:"id"`
	Language  `json:"language,omitempty"`
	Value     string      `json:"value,omitempty"`
	Phonetics []*Phonetic `json:"phonetics,omitempty"`
	Stats     *Stats      `json:"stats,omitempty"`
}

func NewTranslation(l Language, value string) *Translation {
	return &Translation{
		ID:       NewID(),
		Language: l,
		Value:    value,
		Stats:    new(Stats),
	}
}

func (t *Translation) WithPhonetic(p *Phonetic) *Translation {
	t.Phonetics = append(t.Phonetics, p)
	return t
}

func (t *Translation) GetPhonetic(l Language) *Phonetic {
	for _, p := range t.Phonetics {
		if p.Language == l {
			return p
		}
	}
	return &Phonetic{}
}
