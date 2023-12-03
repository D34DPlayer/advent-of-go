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

type symbol struct {
	r rune
	x int
	y int
}

func (s symbol) touches(n number) bool {
	yTouches := n.y >= s.y-1 && n.y <= s.y+1
	xTouches := n.xStart <= s.x+1 && n.xEnd >= s.x-1
	return yTouches && xTouches
}

type number struct {
	n      int
	xStart int
	xEnd   int
	y      int
}

type parser struct {
	input []rune
	pos   int
	r     rune
	y     int
}

func (p *parser) next() (rune, bool) {
	if p.pos == len(p.input)-1 {
		return p.r, false
	}
	p.pos += 1
	p.r = p.input[p.pos]
	return p.r, true
}

func (p *parser) current() (rune, bool) {
	return p.r, p.pos != len(p.input)-1
}

func (p *parser) skip() {
	for r, ok := p.current(); ok; r, ok = p.next() {
		if r != '.' {
			break
		}
	}
}

func (p *parser) parseNumber() number {
	out := ""
	xStart := p.pos
	xEnd := xStart - 1 // -1 for looping purposes

	for r, ok := p.current(); ok; r, ok = p.next() {
		if r >= '0' && r <= '9' {
			out += string(r)
			xEnd++
		} else {
			break
		}
	}

	n, err := strconv.Atoi(out)

	if err != nil {
		panic(err)
	}
	return number{n, xStart, xEnd, p.y}
}

func (p *parser) parseSymbol() symbol {
	r, _ := p.current()
	defer p.next()

	return symbol{r, p.pos, p.y}
}

func (p *parser) parseLine() ([]symbol, []number) {
	symbols := make([]symbol, 0)
	numbers := make([]number, 0)
	for {
		p.skip()
		r, ok := p.current()
		if !ok {
			break
		}
		if r >= '0' && r <= '9' {
			numbers = append(numbers, p.parseNumber())
		} else {
			symbols = append(symbols, p.parseSymbol())
		}
	}
	return symbols, numbers
}

func newParser(input string, y int) parser {
	p := parser{[]rune(input), -1, '\x00', y}
	p.next()
	return p
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	y := 0
	sum := 0
	prevSyms := make([]symbol, 0)
	prevNbrs := make([]number, 0)
	for scanner.Scan() {
		line := scanner.Text()
		p := newParser(line, y)
		syms, nbrs := p.parseLine()
		//fmt.Println("Line:", line)
		//fmt.Println("Symbols:", syms)
		//fmt.Println("Numbers:", nbrs)

		for _, s := range syms {
			for _, n := range prevNbrs {
				if s.touches(n) {
					sum += n.n
				}
			}
			for _, n := range nbrs {
				// TODO: mark the numbers taken here so that they don't count twice
				if s.touches(n) {
					sum += n.n
				}
			}
		}

		for _, n := range nbrs {
			for _, s := range prevSyms {
				if s.touches(n) {
					sum += n.n
				}
			}
		}

		prevSyms = syms
		prevNbrs = nbrs
		y++
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Time:", time.Since(start))
}
