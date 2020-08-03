package main

import "testing"

var testCases = []struct {
	camelCase string
	count     int
}{
	{
		"saveMeOhMyGod",
		5,
	}, {
		"killMeSoftly",
		3,
	},
	{
		"saveChangesInTheEditor",
		5,
	},
	{
		"",
		0,
	},
}

func TestWordCount(t *testing.T) {
	for _, tc := range testCases {
		awc := wordCount(tc.camelCase)
		if awc != tc.count {
			t.Errorf("Expected %d but got %d for camelcase: %s\n", tc.count, awc, tc.camelCase)
		}

	}
}
