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

type game struct {
	id   int
	sets []set
}

func (g game) isSubset(s set) bool {
	out := true
	for _, gameSet := range g.sets {
		out = out && gameSet.isSubset(s)
	}
	return out
}

type set struct {
	r int
	g int
	b int
}

func (a set) isSubset(b set) bool {
	return a.r <= b.r && a.g <= b.g && a.b <= b.b
}

type parser struct {
	input []rune
	pos   int
	r     rune
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

func (p *parser) parseInt() int {
	out := ""
	for r, ok := p.current(); ok; r, ok = p.next() {
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

func (p *parser) skip() {
	for r, ok := p.current(); ok; r, ok = p.next() {
		if !(r == ' ' || r == ':' || r == ',') {
			break
		}
	}
}

func (p *parser) parseAlpha() string {
	out := ""
	for r, ok := p.current(); ok; r, ok = p.next() {
		if r >= 'A' && r <= 'z' {
			out += string(r)
		} else {
			break
		}
	}
	return out
}

func (p *parser) parseCube() (int, string) {
	n := p.parseInt()
	p.skip()
	color := p.parseAlpha()
	return n, color
}

func (p *parser) parseSet() set {
	red := 0
	green := 0
	blue := 0

	for {
		p.skip()
		r, ok := p.current()
		if !ok || r == ';' {
			p.next()
			break
		}
		n, c := p.parseCube()
		switch c {
		case "red":
			red += n
		case "green":
			green += n
		case "blue":
			blue += n
		default:
			panic(fmt.Sprintf("Unexpected color %s", c))
		}
	}

	return set{red, green, blue}
}

func (p *parser) parseGame() game {
	id := p.parseInt()

	sets := make([]set, 0, 1)

	for {
		p.skip()
		_, ok := p.current()
		if !ok {
			break
		}
		sets = append(sets, p.parseSet())
	}
	return game{id, sets}
}

func parseLine(s string) game {
	s = strings.TrimPrefix(s, "Game ")
	p := parser{[]rune(s), -1, '\x00'}
	p.next()

	return p.parseGame()
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	expectedSet := set{12, 13, 14}
	sum := 0

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	for scanner.Scan() {
		line := scanner.Text()
		g := parseLine(line)
		//fmt.Println(line)
		//fmt.Println(g)
		//fmt.Println(g.isSubset(expectedSet))
		if g.isSubset(expectedSet) {
			sum += g.id
		}
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Time:", time.Since(start))
}
