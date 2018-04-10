package flashcards

type Phonetic struct {
	//ID           `json:"id"`
	Language
	Value string `json:"value"`
}

func NewPhonetic(l Language, v string) *Phonetic {
	return &Phonetic{
		//ID:       newID(),
		Language: l,
		Value:    v,
	}
}

func (p *Phonetic) String() string {
	return jsonString(p)
}
