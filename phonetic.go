package flashcards

type Phonetic struct {
	ID ID `json:"id"`

	Language `json:"language,omitempty"`
	Value    string `json:"value,omitempty"`
}

func NewPhonetic(l Language, value string) *Phonetic {
	return &Phonetic{
		ID:       NewID(),
		Language: l,
		Value:    value,
	}
}
