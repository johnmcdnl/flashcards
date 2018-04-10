package flashcards_test

import (
	"testing"
	"github.com/johnmcdnl/flashcards"
	"github.com/stretchr/testify/assert"
)

func TestStats_Record(t *testing.T) {
	var s flashcards.Stats

	assert.Equal(t,0, s.Attempts )
	assert.Equal(t,0, s.Correct )

	s.Record(true)
	assert.Equal(t,1, s.Attempts )
	assert.Equal(t,1, s.Correct )
	s.Record(true)
	assert.Equal(t,2, s.Attempts )
	assert.Equal(t,2, s.Correct )
	s.Record(false)
	assert.Equal(t,3, s.Attempts )
	assert.Equal(t,2, s.Correct )
	s.Record(false)
	assert.Equal(t,4, s.Attempts )
	assert.Equal(t,2, s.Correct )
}