package main

import "testing"

type numberTextTest struct {
	nt    numberText
	input string
	match bool
}

func (ntt numberTextTest) test(t *testing.T) {
	t.Logf("Evaluating '%s' with input '%s'", string(ntt.nt.text), ntt.input)

	match := false
	for _, r := range ntt.input {
		match = match || ntt.nt.evaluateRune(r)
	}

	if match != ntt.match {
		t.Errorf("Match failure, expected %t, got %t", ntt.match, match)
	}
}

func TestNumberText(t *testing.T) {
	tests := []numberTextTest{
		{newNumberText("three", 3), "dqthreexqd", true},
		{newNumberText("three", 3), "adqthres21d", false},
		{newNumberText("two", 3), "6five3two", true},
	}

	for _, test := range tests {
		test.test(t)
	}
}

type calcLineTest struct {
	input string
	value int
}

func (clt calcLineTest) test(t *testing.T) {
	t.Logf("Evaluating '%s'", clt.input)

	value := calcLine(clt.input)

	if value != clt.value {
		t.Errorf("Match failure, expected %d, got %d", clt.value, value)
	}
}

func TestCalcLine(t *testing.T) {
	tests := []calcLineTest{
		{"6twofive3two", 62},
		{"9one43ninedrtznff", 99},
		{"9onenine", 99},
		{"4fiveeight", 48},
	}

	for _, test := range tests {
		test.test(t)
	}
}
