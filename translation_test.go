package flashcards

import (
	"reflect"
	"testing"
)

func TestNewTranslation(t *testing.T) {
	type args struct {
		l Language
		v string
	}
	tests := []struct {
		name string
		args args
		want *Translation
	}{
		{"Simple 1", args{En, "hello"}, &Translation{Language: En, Value: "hello"}},
		{"Simple 2", args{Ru, "hello"}, &Translation{Language: Ru, Value: "hello"}},
		{"Value Required", args{Ru, ""}, nil},
		{"Lang Required", args{Language{}, "hello"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTranslation(tt.args.l, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTranslation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslation_With(t *testing.T) {
	type fields struct {
		Language  Language
		Value     string
		Phonetics []*Phonetic
	}
	type args struct {
		p *Phonetic
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
			tr := &Translation{
				Language:  tt.fields.Language,
				Value:     tt.fields.Value,
				Phonetics: tt.fields.Phonetics,
			}
			if got := tr.With(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translation.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslation_Get(t *testing.T) {
	type fields struct {
		Language  Language
		Value     string
		Phonetics []*Phonetic
	}
	type args struct {
		l Language
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Phonetic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Translation{
				Language:  tt.fields.Language,
				Value:     tt.fields.Value,
				Phonetics: tt.fields.Phonetics,
			}
			if got := tr.Get(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translation.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslation_String(t *testing.T) {
	type fields struct {
		Language  Language
		Value     string
		Phonetics []*Phonetic
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
			tr := &Translation{
				Language:  tt.fields.Language,
				Value:     tt.fields.Value,
				Phonetics: tt.fields.Phonetics,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("Translation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
