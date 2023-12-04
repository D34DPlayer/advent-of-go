package main

import "testing"

type someTest struct {
	input string
	match bool
}

func (XXX someTest) test(t *testing.T) {
}

func TestNumberText(t *testing.T) {
	tests := []someTest{}

	for _, test := range tests {
		test.test(t)
	}
}
