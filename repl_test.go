package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch.\nInput: %q\nExpected: %v\nGot: %v",
				c.input, c.expected, actual)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Word mismatch.\nInput: %q\nExpected: %v\nGot: %v",
					c.input, c.expected, actual)
				break
			}
		}
	}
}
