package flashcards

import "github.com/sirupsen/logrus"

type Stats struct {
	Source   Language
	Target   Language
	Attempts int
	Correct  int
}

func NewStats(src, tgt Language) *Stats {
	logrus.Debugln(`func NewStats(src, tgt Language)*Stats{`, src, tgt)

	return &Stats{
		Source: src,
		Target: tgt,
	}
}
