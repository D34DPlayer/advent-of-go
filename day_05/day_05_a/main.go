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

type seedJourney struct {
	id   int
	step int
}

func newSeedJourney(id int) seedJourney {
	return seedJourney{id, 0}
}

type recipe struct {
	dest   int
	source int
	length int
}

func (r recipe) appliesTo(sj seedJourney) bool {
	return sj.id >= r.source && sj.id < r.source+r.length
}

func (r recipe) apply(sj seedJourney) int {
	return sj.id - r.source + r.dest
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

		if skipped == 0 {
			p.next()
			break
		}
	}
	return out
}

func newParser(input string) parser {
	p := parser{[]rune(input), -1}
	p.next()
	return p
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	seeds := make([]seedJourney, 0, 16)
	currentStep := 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		fmt.Println("Line:", line)
		if len(seeds) == 0 {
			line = strings.TrimPrefix(line, "seeds: ")
			p := newParser(line)
			for _, seedId := range p.parseIntArr() {
				seeds = append(seeds, newSeedJourney(seedId))
			}
			fmt.Println("Seeds:", seeds)
		} else if len(line) != 0 && line[len(line)-1] == ':' {
			currentStep++
			fmt.Println("Changed step:", currentStep)
			continue
		} else {
			p := newParser(line)
			values := p.parseIntArr()
			if len(values) != 3 {
				panic("Expected recipe with 3 elements")
			}
			r := recipe{values[0], values[1], values[2]}

			for i, s := range seeds {
				if s.step < currentStep && r.appliesTo(s) {
					seeds[i] = seedJourney{r.apply(s), currentStep}
				}
			}
		}
	}
	minId := seeds[0].id
	for _, s := range seeds {
		if s.id < minId {
			minId = s.id
		}
	}
	fmt.Println("Min:", minId)
	fmt.Println("Time:", time.Since(start))
}
