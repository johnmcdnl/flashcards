package flashcards

type Translation struct {
	Language  `json:"language,omitempty"`
	Value     string      `json:"value,omitempty"`
	Phonetics []*Phonetic `json:"phonetics,omitempty"`
}

func NewTranslation() *Translation {
	return &Translation{}
}
