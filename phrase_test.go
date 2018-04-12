package flashcards

import (
	"reflect"
	"testing"
)

func TestNewPhrase(t *testing.T) {
	type args struct {
		t *Translation
	}
	tests := []struct {
		name string
		args args
		want *Phrase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPhrase(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhrase_With(t *testing.T) {
	type fields struct {
		ID           ID
		Translations []*Translation
	}
	type args struct {
		t *Translation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Phrase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Phrase{
				ID:           tt.fields.ID,
				Translations: tt.fields.Translations,
			}
			if got := p.With(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Phrase.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhrase_Get(t *testing.T) {
	type fields struct {
		ID           ID
		Translations []*Translation
	}
	type args struct {
		l Language
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Translation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Phrase{
				ID:           tt.fields.ID,
				Translations: tt.fields.Translations,
			}
			if got := p.Get(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Phrase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhrase_Answer(t *testing.T) {
	type fields struct {
		ID           ID
		Translations []*Translation
	}
	type args struct {
		l   Language
		ans string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Phrase{
				ID:           tt.fields.ID,
				Translations: tt.fields.Translations,
			}
			if err := p.Answer(tt.args.l, tt.args.ans); (err != nil) != tt.wantErr {
				t.Errorf("Phrase.Answer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPhrase_String(t *testing.T) {
	type fields struct {
		ID           ID
		Translations []*Translation
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Phrase{
				ID:           tt.fields.ID,
				Translations: tt.fields.Translations,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Phrase.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
