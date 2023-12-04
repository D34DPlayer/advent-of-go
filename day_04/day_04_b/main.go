package main

import (
	"bufio"
	"fmt"
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

type card struct {
	id      int
	winners []int
	own     []int
}

func (c card) matchesAmount() int {
	count := 0
	for _, o := range c.own {
		for _, w := range c.winners {
			if w == o {
				count++
				break
			}
		}
	}
	return count
}

type parser struct {
	input []rune
	pos   int
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
	for {
		out = append(out, p.parseInt())
		skipped := p.skip()
		r, _ := p.current()

		if skipped == 0 || r == '|' {
			p.next()
			break
		}
	}
	return out
}

func (p *parser) parseCard() card {
	id := p.parseInt()
	r, _ := p.current()
	if r != ':' {
		panic("Expected colon")
	}
	p.next()
	p.skip()
	winners := p.parseIntArr()
	p.skip()
	own := p.parseIntArr()

	return card{id, winners, own}
}

func parseLine(input string) int {
	input = strings.TrimPrefix(input, "Card")
	p := parser{[]rune(input), -1}
	p.next()
	p.skip()

	c := p.parseCard()
	return c.matchesAmount()
}

func makeWithValue(amount int, value int) []int {
	out := make([]int, amount)
	for i := range out {
		out[i] = value
	}
	return out
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	points := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		points = append(points, parseLine(line))
	}
	//fmt.Println("Points:", points)

	cards := makeWithValue(len(points), 1)
	sum := 0
	for i, p := range points {
		c := cards[i]
		sum += c
		for j := 1; j < p+1; j++ {
			cards[i+j] += c
		}
		//if i < 16 {
		//	fmt.Println("point", p)
		//	fmt.Println("card", c)
		//	fmt.Println("Loop", i, cards)
		//}
	}

	//fmt.Println("Cards:", cards)

	fmt.Println("Sum:", sum)
	fmt.Println("Time:", time.Since(start))
}
