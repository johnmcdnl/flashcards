package flashcards

type Phonetic struct {
	Language `json:"language,omitempty"`
	Value    string `json:"value,omitempty"`
}
