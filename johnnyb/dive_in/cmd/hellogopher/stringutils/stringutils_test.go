package stringutils_test

import (
	"stringutils"
	"testing"
)

func TestUpper(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"empty", "", ""},
		{"simple", "hello", "HELLO"},
		{"mixed", "Hello, World", "HELLO, WORLD"},
		{"special", "Hello, 世界", "HELLO, 世界"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringutils.Upper(tt.arg); got != tt.want {
				t.Errorf("Upper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"empty", "", ""},
		{"simple", "HELLO", "hello"},
		{"mixed", "Hello, World", "hello, world"},
		{"special", "Hello, 世界", "hello, 世界"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringutils.Lower(tt.arg); got != tt.want {
				t.Errorf("Lower() = %v, want %v", got, tt.want)
			}
		})
	}
}
