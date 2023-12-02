package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type numberText struct {
	text    []rune
	value   int
	counter int
	length  int
}

func newNumberText(text string, value int) numberText {
	return numberText{[]rune(text), value, 0, len(text)}
}

func newNumbers() []numberText {
	return []numberText{
		newNumberText("0", 0), //newNumberText("zero", 0),
		newNumberText("1", 1), newNumberText("one", 1),
		newNumberText("2", 2), newNumberText("two", 2),
		newNumberText("3", 3), newNumberText("three", 3),
		newNumberText("4", 4), newNumberText("four", 4),
		newNumberText("5", 5), newNumberText("five", 5),
		newNumberText("6", 6), newNumberText("six", 6),
		newNumberText("7", 7), newNumberText("seven", 7),
		newNumberText("8", 8), newNumberText("eight", 8),
		newNumberText("9", 9), newNumberText("nine", 9),
	}
}

func (nt *numberText) reset(r rune) {
	nt.counter = 0
	if r == nt.text[0] {
		nt.counter = min(1, nt.length-1) // This prevents the counter getting out of bounds
	}
}

func (nt *numberText) evaluateRune(r rune) (out bool) {
	expectedRune := nt.text[nt.counter]

	if r != expectedRune || nt.counter+1 == nt.length {
		defer nt.reset(r)
	}

	if r == expectedRune {
		nt.counter += 1
	}

	return nt.counter == nt.length
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumber(r rune, numbers []numberText) (n int, ok bool) {
	n = -1
	ok = false
	for i, nt := range numbers {
		res := nt.evaluateRune(r)
		numbers[i] = nt // n is a copy of the struct!
		if res {
			n = nt.value // Don't return here, as we need to invalidate all numbers
			ok = true
		}
	}
	return
}

func calcLine(s string) int {
	numbers := newNumbers()
	firstSet := false
	first := 0
	last := 0

	for _, r := range s {
		if n, ok := isNumber(r, numbers); ok {
			if !firstSet {
				first = n
				firstSet = true
			}
			last = n
		}
	}

	return 10*first + last
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	sum := 0
	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	for scanner.Scan() {
		line := scanner.Text()
		value := calcLine(line)
		sum += value
		// fmt.Println("Line:", line)
		// fmt.Println("Value:", value)
	}
	check(scanner.Err())
	fmt.Println("Sum:", sum)
	fmt.Println("Time:", time.Since(start))
}
