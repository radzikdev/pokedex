package main

import "testing"

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
		input:    "word ",
		expected: []string{"word"},
	},
	{
		input:    "  HeLLo ,  world  ",
		expected: []string{"hello", ",", "world"},
	},
	{
		input:    " ",
		expected: []string{},
	},
   }

    for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("size does not match %v vs %v", len(actual), len(c.expected))
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("word does not match, %v != %v", word, expectedWord)
		}
	}
}
}
