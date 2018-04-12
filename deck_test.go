package flashcards

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	type args struct {
		p *Phrase
	}
	tests := []struct {
		name string
		args args
		want *Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeck(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_With(t *testing.T) {
	type fields struct {
		Phrases []*Phrase
	}
	type args struct {
		p *Phrase
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deck{
				Phrases: tt.fields.Phrases,
			}
			if got := d.With(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deck.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Next(t *testing.T) {
	type fields struct {
		Phrases []*Phrase
	}
	tests := []struct {
		name   string
		fields fields
		want   *Phrase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deck{
				Phrases: tt.fields.Phrases,
			}
			if got := d.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deck.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Answer(t *testing.T) {
	type fields struct {
		Phrases []*Phrase
	}
	type args struct {
		p   *Phrase
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
			d := &Deck{
				Phrases: tt.fields.Phrases,
			}
			if err := d.Answer(tt.args.p, tt.args.l, tt.args.ans); (err != nil) != tt.wantErr {
				t.Errorf("Deck.Answer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeck_SeedFromCSV(t *testing.T) {
	type fields struct {
		Phrases []*Phrase
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deck{
				Phrases: tt.fields.Phrases,
			}
			d.SeedFromCSV()
		})
	}
}

func TestDeck_String(t *testing.T) {
	type fields struct {
		Phrases []*Phrase
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
			d := &Deck{
				Phrases: tt.fields.Phrases,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("Deck.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
