package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type parser struct {
	input []rune
	pos   int
	done  bool
}

func calcDiff(a []int) []int {
	out := make([]int, len(a)-1)
	for i := 1; i < len(a); i++ {
		diff := a[i] - a[i-1]
		out[i-1] = diff
	}
	return out
}

func newParser(input string) parser {
	p := parser{[]rune(input), -1, false}
	p.next()
	return p
}

func (p *parser) next() (rune, bool) {
	if p.pos == len(p.input)-1 {
		p.done = true
		return p.input[p.pos], false
	}
	p.pos++
	return p.input[p.pos], true
}

func (p *parser) current() (rune, bool) {
	return p.input[p.pos], !p.done
}

func (p *parser) skip() int {
	amount := 0
	for r, ok := p.current(); ok; r, ok = p.next() {
		if !(r == ' ') {
			break
		}
		amount++
	}
	return amount
}

func (p *parser) parseInt() int {
	r, _ := p.current()
	out := string(r)
	for r, ok := p.next(); ok; r, ok = p.next() {
		if r >= '0' && r <= '9' {
			out += string(r)
		} else {
			break
		}
	}

	id, err := strconv.Atoi(out)

	if err != nil {
		panic(err)
	}
	return id
}

func (p *parser) parseIntArray() []int {
	out := make([]int, 0)
	for {
		_, ok := p.current()
		if !ok {
			break
		}
		out = append(out, p.parseInt())
		p.skip()
	}
	return out
}

func isAllZeroes(a []int) bool {
	for _, i := range a {
		if i != 0 {
			return false
		}
	}
	return true
}

func solve(a []int) int {
	firstValues := []int{a[0]}
	diffs := calcDiff(a)
	for !isAllZeroes(diffs) {
		firstValues = append(firstValues, diffs[0])
		diffs = calcDiff(diffs)
	}
	fmt.Println(firstValues)

	diff := 0
	maxIndex := len(firstValues) - 1
	for i, _ := range firstValues {
		diff = -diff + firstValues[maxIndex-i]
	}
	return diff
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		p := newParser(line)
		values := p.parseIntArray()
		s := solve(values)
		//fmt.Println("Line:", line)
		//fmt.Println("Values:", values)
		//fmt.Println("Solution:", s)
		sum += s
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Time:", time.Since(start))
}
