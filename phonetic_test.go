package flashcards

import (
	"reflect"
	"testing"
)

func TestNewPhonetic(t *testing.T) {
	type args struct {
		l Language
		v string
	}
	tests := []struct {
		name string
		args args
		want *Phonetic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPhonetic(tt.args.l, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhonetic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhonetic_String(t *testing.T) {
	type fields struct {
		Language Language
		Value    string
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
			p := &Phonetic{
				Language: tt.fields.Language,
				Value:    tt.fields.Value,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Phonetic.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
