package flashcards_test

import (
	"testing"
	"github.com/johnmcdnl/flashcards"
	"github.com/stretchr/testify/assert"
)

func TestNewPhrase(t *testing.T) {
	var p = flashcards.NewPhrase(nil)
	assert.Len(t, p.Translations, 0)

	var pWithVal = flashcards.NewPhrase(flashcards.NewTranslation(flashcards.Eng, "hello"))
	assert.Len(t, pWithVal.Translations, 1)
	assert.Equal(t,flashcards.Eng, pWithVal.GetTranslation(flashcards.Eng).Language)
	assert.Equal(t,"hello", pWithVal.GetTranslation(flashcards.Eng).Value)
}

func TestPhrase_WithTranslation(t *testing.T) {
	var p = flashcards.NewPhrase(nil)

	assert.Empty(t, p.Translations)
	assert.Nil(t, p.GetTranslation(flashcards.Eng))

	p.WithTranslation(
		flashcards.NewTranslation(flashcards.Eng, "hello")).WithTranslation(
		flashcards.NewTranslation(flashcards.Rus, "Здравствуйте")).WithTranslation(
		flashcards.NewTranslation(flashcards.Fra, "bonjour"))

	assert.Len(t, p.Translations, 3)

	p.WithTranslation(flashcards.NewTranslation(flashcards.Eng, "hello"))
	assert.Len(t, p.Translations, 3)

}

func TestPhrase_GetTranslation(t *testing.T) {
	var p = new(flashcards.Phrase)
	p.WithTranslation(
		flashcards.NewTranslation(flashcards.Eng, "hello")).WithTranslation(
		flashcards.NewTranslation(flashcards.Rus, "Здравствуйте")).WithTranslation(
		flashcards.NewTranslation(flashcards.Fra, "bonjour"))

	assert.Equal(t, flashcards.Eng, p.GetTranslation(flashcards.Eng).Language)
	assert.Equal(t, "hello", p.GetTranslation(flashcards.Eng).Value)

	assert.Equal(t, flashcards.Fra, p.GetTranslation(flashcards.Fra).Language)
	assert.Equal(t, "bonjour", p.GetTranslation(flashcards.Fra).Value)

	assert.Equal(t, flashcards.Rus, p.GetTranslation(flashcards.Rus).Language)
	assert.Equal(t, "Здравствуйте", p.GetTranslation(flashcards.Rus).Value)

	// There should only be a single value per language
	p.WithTranslation(flashcards.NewTranslation(flashcards.Eng, "helloV2"))
	assert.Equal(t, flashcards.Eng, p.GetTranslation(flashcards.Eng).Language)
	assert.Equal(t, "helloV2", p.GetTranslation(flashcards.Eng).Value)

	assert.Equal(t, flashcards.Fra, p.GetTranslation(flashcards.Fra).Language)
	assert.Equal(t, "bonjour", p.GetTranslation(flashcards.Fra).Value)

	assert.Equal(t, flashcards.Rus, p.GetTranslation(flashcards.Rus).Language)
	assert.Equal(t, "Здравствуйте", p.GetTranslation(flashcards.Rus).Value)

}


func TestPhrase_Attempt(t *testing.T) {
	var p = new(flashcards.Phrase)
	p.WithTranslation(
		flashcards.NewTranslation(flashcards.Eng, "hello")).WithTranslation(
		flashcards.NewTranslation(flashcards.Rus, "Здравствуйте")).WithTranslation(
		flashcards.NewTranslation(flashcards.Fra, "bonjour"))

	assert.True(t, p.Attempt(flashcards.Eng, "Hello"))
	assert.True(t, p.Attempt(flashcards.Eng, "hello"))
	assert.False(t, p.Attempt(flashcards.Eng, "Здравствуйте"))
	assert.True(t, p.Attempt(flashcards.Rus, "Здравствуйте"))
}