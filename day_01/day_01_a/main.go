package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func calcLine(s string) int {
	firstSet := false
	first := 0
	last := 0

	for _, r := range s {
		if isNumber(r) {
			n := int(r - '0') // rune is a i32, so we can take its distance to '0' as its value
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
