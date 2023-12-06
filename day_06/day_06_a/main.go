package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
}

func newParser(input string) parser {
	p := parser{[]rune(input), -1}
	p.next()
	return p
}

func (p *parser) next() (rune, bool) {
	if p.pos == len(p.input)-1 {
		return p.input[p.pos], false
	}
	p.pos++
	return p.input[p.pos], true
}

func (p *parser) current() (rune, bool) {
	return p.input[p.pos], p.pos != len(p.input)-1
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

func (p *parser) parseIntArr() []int {
	out := make([]int, 0)
	p.skip()
	for {
		out = append(out, p.parseInt())
		skipped := p.skip()

		if skipped == 0 {
			p.next()
			break
		}
	}
	return out
}

func calculateRadical(a, b, c int) float64 {
	// Good old $\sqrt{B^2-4AC}$
	return math.Sqrt(math.Pow(float64(b), 2) - float64(4*a*c))
}

func solve(time, distance int) int {
	// $Amount = ceil(T/2+Rad/2)-floor(T/2-Rad/2) - 2$
	radical := calculateRadical(1, time, distance) // actually -1, time, -distance but minuses cancel each other
	smallest := math.Floor((float64(time) - radical) / 2)
	biggest := math.Ceil((float64(time) + radical) / 2)
	return int(biggest - smallest - 1)
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	// Times
	scanner.Scan()
	p := newParser(strings.TrimPrefix(scanner.Text(), "Time:"))
	times := p.parseIntArr()
	// Distances
	scanner.Scan()
	p = newParser(strings.TrimPrefix(scanner.Text(), "Distance:"))
	distances := p.parseIntArr()

	product := 1
	for i, t := range times {
		d := distances[i]
		s := solve(t, d)
		product *= s
		//fmt.Println("Time", i, ":", t)
		//fmt.Println("Distance", i, ":", d)
		//fmt.Println("Solution", i, ":", s)
	}
	fmt.Println("Result:", product)
	fmt.Println("Time:", time.Since(start))
}
