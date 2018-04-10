package flashcards

type Translation struct {
	//ID                    `json:"id"`
	Language
	Value     string      `json:"value"`
	Phonetics []*Phonetic `json:"phonetics,omitempty"`
}

func NewTranslation(l Language, v string) *Translation {
	return &Translation{
		//ID: newID(),
		Language: l,
		Value:    v,
	}
}

func (t *Translation) String() string {
	return jsonString(t)
}

func (t *Translation) With(p *Phonetic) *Translation {
	if p == nil {
		return t
	}
	t.Phonetics = append(t.Phonetics, p)
	return t
}
