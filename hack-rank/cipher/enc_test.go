package main

import "testing"

var testCases = []struct {
	message   string
	rotate    int
	encrypted string
}{
	{
		"abz",
		2,
		"cdb",
	},
	{
		"abcdefghijklmnopqrstuvwxyz",
		3,
		"defghijklmnopqrstuvwxyzabc",
	},
	{
		"middle-Outz",
		2,
		"okffng-Qwvb",
	},
}

func TestEncrypt(t *testing.T) {
	for _, tc := range testCases {
		ae := encrypt(tc.message, tc.rotate)
		if ae != tc.encrypted {
			t.Errorf("Expected %s but got %s for orig message %s\n", tc.encrypted, ae, tc.message)
		}
	}
}
