package flashcards

import "github.com/sirupsen/logrus"

//Stats represents how accurate attempts have been
type Stats struct {
	Source   Language
	Target   Language
	Attempts int
	Correct  int
}

//NewStats generates a new stats
func NewStats(src, tgt Language) *Stats {
	logrus.Debugln(`func NewStats(src, tgt Language)*Stats{`, src, tgt)

	return &Stats{
		Source: src,
		Target: tgt,
	}
}
