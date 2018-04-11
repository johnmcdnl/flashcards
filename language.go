package flashcards

import (
	"strings"
	"github.com/sirupsen/logrus"
)

//Language is a ISO language
type Language struct {
	name string
	Code string `json:"langCode"`
}

// Predefined Languages
var (
	En = GetLang("en")
	Fr = GetLang("fr")
	Ru = GetLang("ru")
)

var languages = []Language{
	{name: "english", Code: "en"},
	{name: "french", Code: "fr"},
	{name: "russian", Code: "ru"},
}

//GetLang gets a Language based on its ISO code
func GetLang(code string) Language {
	code = strings.Trim(code, " ")
	for _, l := range languages {
		if l.Code == code{
			return l
		}
	}
	logrus.Errorln(code)
	panic("unknown")
	return Language{name: "unknown", Code: code}
}

// String represents a Language
func (l *Language) String() string {
	return jsonString(l)
}
