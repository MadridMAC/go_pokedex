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
			input:    "	hello	world	",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check actual and input lengths; if inequal, print error and fail test
		if len(actual) != len(c.expected) {
			t.Errorf("error: expected output length does not match actual output length | expected: %d, actual: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("error: expected word does not match actual word | expected: %s, actual: %s", expectedWord, word)
			}
		}
	}
}
