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
			input:    " Test  ME or I DONT really care",
			expected: []string{"test", "me", "or", "i", "dont", "really", "care"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  Go  is   Awesome  ",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    "  Spaces     Everywhere  ",
			expected: []string{"spaces", "everywhere"},
		},
		{
			input:    "  ",
			expected: []string{}, // Empty input should return an empty slice
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Wrong length of slice! expected: %v actual: %v", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			excpectedWord := c.expected[i]
			if word != excpectedWord {
				t.Errorf("Expected: %v insted recived: %v", excpectedWord, word)
			}
		}
	}
}
