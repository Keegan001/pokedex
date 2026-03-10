package main 

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input   string
		expected []string
	}{
		{
			input: "   hello   world   ",
			expected: []string{ "hello", "world"},
		},
	}
	for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice against the expected slice
	if len(actual) != len(c.expected) {
		t.Errorf("Length mismatch, expected %v, got %v",len(c.input),len(actual))
	}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("Word mismatch, expected %v, got %v", expectedWord, word)
			}
		}
	}
}