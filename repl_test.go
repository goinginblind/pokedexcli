package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  123hello 		wORRRrld  ",
			expected: []string{"123hello", "worrrrld"},
		},
		{
			input:    "pokeballs are cool 2 me l 	o 	l		",
			expected: []string{"pokeballs", "are", "cool", "2", "me", "l", "o", "l"},
		},
		{
			input:    "Mr. House sends his regards, Benny",
			expected: []string{"mr.", "house", "sends", "his", "regards,", "benny"},
		},
		// add more cases here
	}

	var successes int = 0
	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %v - got: %v", len(c.expected), len(actual))
			t.Fail()
		}

		for i := range actual {
			actualWord, expectedWord := actual[i], c.expected[i]
			if actualWord != expectedWord {
				t.Errorf("expected: %v - got: %v", expectedWord, actualWord)
				t.Fail()
			}
		}
		successes++
	}
	fmt.Printf("%v out of %v passed", successes, len(cases))
}
