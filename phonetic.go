package flashcards

type Phonetic struct {
	Language `json:"language,omitempty"`
	Value    string `json:"value,omitempty"`
}

func NewPhonetic(l Language, value string) *Phonetic {
	return &Phonetic{
		Language: l,
		Value:    value,
	}
}
