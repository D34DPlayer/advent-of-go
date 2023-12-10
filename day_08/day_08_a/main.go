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

func (p *parser) parseLR() []int {
	out := make([]int, 0)
	for r, ok := p.current(); ok; r, ok = p.next() {
		if r == 'L' {
			out = append(out, 0)
		} else if r == 'R' {
			out = append(out, 1)
		}
	}
	return out
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

func (p *parser) parseTuple() []string {
	p.next() // TODO: check '(' && ok
	l := p.parseAlpha()
	p.next() // TODO: check ',' && ok
	p.skip()
	r := p.parseAlpha() // TODO: check ')'
	return []string{l, r}
}

func (p *parser) parseNode() (string, []string) {
	n := p.parseAlpha()
	p.skip()
	p.next() // TODO: check '=' && ok
	p.skip()
	t := p.parseTuple()
	return n, t
}

func solve(nodes map[string][]string, lr []int) int {
	count := 0
	current, ok := nodes["AAA"]
	if !ok {
		panic("Couldn't find node AAA")
	}

	for {
		for _, step := range lr {
			count++
			nextNode := current[step]
			if nextNode == "ZZZ" {
				return count
			}
			current, ok = nodes[nextNode]
			if !ok {
				panic("Couldn't find node " + nextNode)
			}
		}
	}
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	scanner.Scan()
	line := scanner.Text()
	p := newParser(line)
	lr := p.parseLR()
	//fmt.Println("LR:", line)
	//fmt.Println("LR:", lr)
	scanner.Scan() // Empty line

	nodes := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		p := newParser(line)
		key, values := p.parseNode()
		nodes[key] = values
		//fmt.Println("Line:", line)
		//fmt.Println("Key:", key)
		//fmt.Println("values:", values)
	}

	fmt.Println("Solution:", solve(nodes, lr))
	fmt.Println("Time:", time.Since(start))
}
