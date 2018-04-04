package flashcards

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Stats struct {
	ID         ID      `json:"-"`
	Attempts   int     `json:"attempts"`
	Correct    int     `json:"correct"`
	Percentage float64 `json:"percentage"`
	Weighting  float64 `json:"weighting"`
}

func NewStats() *Stats {
	s := &Stats{ID: NewID(), Weighting: 50}
	s.update()
	return s
}

func (s *Stats) String() string {
	return fmt.Sprintf("Percentage: %.1f %s\t | Weight: %.1f", s.Percentage*100, "%", s.Weighting)
}

func (s *Stats) CorrectAttempt() {
	s.Attempts++
	s.Correct++
	s.update()
	logrus.Info("Correct \t | ", s.String())
}

func (s *Stats) WrongAttempt() {
	s.Attempts++
	s.update()
	logrus.Info("Wrong!!! \t | ", s.String())
}

func (s *Stats) update() {
	s.updatePercentage()
	s.updateWeighting()
}

func (s *Stats) updatePercentage() {
	if s.Attempts == 0 {
		s.Percentage = 0
		return
	}
	if s.Correct == 0 {
		s.Percentage = 0
		return
	}
	s.Percentage = float64(s.Correct) / float64(s.Attempts)
}

func (s *Stats) updateWeighting() {

	s.Weighting = (1 - s.Percentage) * 100

	if s.Attempts <= 3 {
		s.Weighting = 50
	}

	if s.Weighting == 0 {
		s.Weighting = 1
	}
}
