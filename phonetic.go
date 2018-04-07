package flashcards

// Phonetic is a helper on how to pronounce a phrase
type Phonetic struct {
	ID       ID `json:"-"`
	Language `json:"language,omitempty"`
	Value    string `json:"value,omitempty"`
}

// NewPhonetic craetes a new Phonetic
func NewPhonetic(l Language, value string) *Phonetic {
	return &Phonetic{
		ID:       NewID(),
		Language: l,
		Value:    value,
	}
}
