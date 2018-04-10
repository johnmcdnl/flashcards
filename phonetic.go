package flashcards

type Phonetic struct {
	Language
	Value string
}

func (p *Phonetic) String() string {
	return jsonString(p)
}
