package routes

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/johnmcdnl/flashcards"
)

func ListTranslations(w http.ResponseWriter, r *http.Request) {
	p := deck.GetPhraseByID(chi.URLParam(r, "phraseID"))
	if err := render.RenderList(w, r, NewTranslationListResponse(p.Translations)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func GetTranslationByLang(w http.ResponseWriter, r *http.Request) {
	p := deck.GetPhraseByID(chi.URLParam(r, "phraseID"))
	lang := chi.URLParam(r, "langCode")
	var translation *flashcards.Translation
	if p != nil {
		for _, t := range p.Translations {
			if t.Language.Code == lang {
				translation = t
			}
		}
	}

	if err := render.Render(w, r, NewTranslationResponse(translation)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func NewTranslationListResponse(translations []*flashcards.Translation) []render.Renderer {
	var list []render.Renderer
	for _, translation := range translations {
		list = append(list, NewTranslationResponse(translation))
	}
	return list
}

func (rd *TranslationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewTranslationResponse(translation *flashcards.Translation) *TranslationResponse {
	resp := &TranslationResponse{Translation: translation}

	return resp
}

type TranslationResponse struct {
	*flashcards.Translation
}