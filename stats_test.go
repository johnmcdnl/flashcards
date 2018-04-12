package flashcards

import (
	"reflect"
	"testing"
)

func TestNewStats(t *testing.T) {
	type args struct {
		src Language
		tgt Language
	}
	tests := []struct {
		name string
		args args
		want *Stats
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStats(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
