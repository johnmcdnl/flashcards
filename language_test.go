package flashcards

import (
	"reflect"
	"testing"
)

func TestGetLang(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want Language
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLang(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguage_String(t *testing.T) {
	type fields struct {
		name string
		Code string
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
			l := &Language{
				name: tt.fields.name,
				Code: tt.fields.Code,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("Language.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
