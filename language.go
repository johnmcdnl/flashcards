package flashcards

import (
	"strings"
	"github.com/sirupsen/logrus"
)

type Language struct {
	name string
	Code string `json:"langCode"`
}

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

func (l *Language) String() string {
	return jsonString(l)
}
