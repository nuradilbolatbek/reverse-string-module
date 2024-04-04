package _func

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"12345", "54321"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			reversed := Reverse(test.input)
			if reversed != test.expected {
				t.Errorf("Reverse(%s): expected %s, got %s", test.input, test.expected, reversed)
			}
		})
	}
}

func TestCountSymbols(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"world", 5},
		{"12345", 5},
		{"", 0},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			count := CountSymbols(test.input)
			if count != test.expected {
				t.Errorf("CountSymbols(%s): expected %d, got %d", test.input, test.expected, count)
			}
		})
	}
}
