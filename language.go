package flashcards

type Language struct {
	name string
	Code string `json:"langCode"`
}


var (
	En = Language{name: "english", Code: "en"}
	Fr = Language{name: "french", Code: "fr"}
	Ru = Language{name: "russian", Code: "ru"}
)

func (l *Language) String() string {
	return jsonString(l)
}
