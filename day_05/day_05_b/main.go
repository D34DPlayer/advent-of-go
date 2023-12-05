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
	idStart int
	length  int
	step    int
}

type recipe struct {
	dest   int
	source int
	length int
}

func (r recipe) apply(sj seedJourney, step int) []seedJourney {
	if sj.step >= step {
		return []seedJourney{sj}
	}

	newJourneys := make([]seedJourney, 0)
	if sj.idStart >= r.source && sj.idStart+sj.length <= r.source+r.length {
		//fmt.Println("Full match")
		newJourneys = append(newJourneys, seedJourney{sj.idStart - r.source + r.dest, sj.length, step})
	} else if sj.idStart >= r.source+r.length || sj.idStart+sj.length <= r.source {
		//fmt.Println("No match")
		newJourneys = append(newJourneys, sj)
	} else {
		if r.source > sj.idStart {
			//fmt.Println("Excess left")
			start := sj.idStart
			end := r.source
			newJourneys = append(newJourneys, seedJourney{start, end - start, sj.step})
		}
		if r.source+r.length < sj.idStart+sj.length {
			//fmt.Println("Excess right")
			start := r.source + r.length
			end := sj.idStart + sj.length
			newJourneys = append(newJourneys, seedJourney{start, end - start, sj.step})
		}
		start := max(r.source, sj.idStart)
		end := min(r.source+r.length, sj.idStart+sj.length)
		newJourneys = append(newJourneys, seedJourney{start - r.source + r.dest, end - start, step})
	}
	return newJourneys
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

func (p *parser) parseSeedRange() seedJourney {
	start := p.parseInt()
	p.skip()
	length := p.parseInt()
	return seedJourney{start, length, 0}
}

func (p *parser) parseRecipe() recipe {
	dest := p.parseInt()
	p.skip()
	source := p.parseInt()
	p.skip()
	length := p.parseInt()
	return recipe{dest, source, length}
}

func (p *parser) parseSeeds() []seedJourney {
	seeds := make([]seedJourney, 0)
	for {
		p.skip()
		_, ok := p.current()
		if !ok {
			break
		}
		seeds = append(seeds, p.parseSeedRange())
	}
	return seeds
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
	var seeds []seedJourney
	currentStep := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if seeds == nil {
			line = strings.TrimPrefix(line, "seeds: ")
			p := newParser(line)
			seeds = p.parseSeeds()
			//fmt.Println(seeds)
		} else if len(line) != 0 && line[len(line)-1] == ':' {
			currentStep++
			//fmt.Println("Changed step:", currentStep)
			//fmt.Println(seeds)
			continue
		} else {
			p := newParser(line)
			r := p.parseRecipe()
			newSeeds := make([]seedJourney, 0, len(seeds))

			for _, s := range seeds {
				appliedSeeds := r.apply(s, currentStep)
				//fmt.Println("Recipe", r)
				//fmt.Println("Seed", s)
				//fmt.Println("Applied seeds", appliedSeeds)
				newSeeds = append(newSeeds, appliedSeeds...)
			}

			seeds = newSeeds
		}
	}
	minId := seeds[0].idStart
	for _, s := range seeds {
		if s.idStart < minId {
			minId = s.idStart
		}
	}
	//fmt.Println("Seeds:", seeds)
	fmt.Println("Min:", minId)
	fmt.Println("Time:", time.Since(start))
}
