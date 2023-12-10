package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
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

func (p *parser) parseAlphaNumeric() string {
	out := ""
	for r, ok := p.current(); ok; r, ok = p.next() {
		if r >= '0' && r <= 'z' {
			out += string(r)
		} else {
			break
		}
	}
	return out
}

func (p *parser) parseTuple() []string {
	p.next() // TODO: check '(' && ok
	l := p.parseAlphaNumeric()
	p.next() // TODO: check ',' && ok
	p.skip()
	r := p.parseAlphaNumeric() // TODO: check ')'
	return []string{l, r}
}

func (p *parser) parseNode() (string, []string) {
	n := p.parseAlphaNumeric()
	p.skip()
	p.next() // TODO: check '=' && ok
	p.skip()
	t := p.parseTuple()
	return n, t
}

func startingPaths(nodes map[string][]string) [][]string {
	out := make([][]string, 0)
	for k, v := range nodes {
		r, _ := utf8.DecodeLastRuneInString(k)
		//fmt.Println(k, string(r))
		if r == 'A' {
			out = append(out, v)
		}
	}
	return out
}

func isNodeSolved(s string) bool {
	r, _ := utf8.DecodeLastRuneInString(s)
	return r == 'Z'
}

/* Bruteforce
func solve(nodes map[string][]string, lr []int) int {
	count := 0
	paths := startingPaths(nodes)
	for {
		for _, step := range lr {
			count++
			solved := 0
			for i, p := range paths {
				nextNode := p[step]
				n, ok := nodes[nextNode]
				if !ok {
					panic("Couldn't find node " + nextNode)
				}
				paths[i] = n
				if isNodeSolved(nextNode) {
					solved++
				}
			}
			if solved == len(paths) {
				return count
			}
		}
	}
}
*/

// INSIGHTS: for each path always the same solution and more importantly, always the same distance between solutions
func solveOne(startingNode []string, nodes map[string][]string, lr []int) int {
	count := 0
	path := startingNode
	ok := true
	for {
		for _, step := range lr {
			count++
			nextNode := path[step]
			if isNodeSolved(nextNode) {
				return count
			}
			path, ok = nodes[nextNode]
			if !ok {
				panic("Couldn't find node " + nextNode)
			}
		}
	}
}

// Stolen from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], nil)
	}

	return result
}

func solve(nodes map[string][]string, lr []int) int {
	paths := startingPaths(nodes)
	solutions := make([]int, len(paths))

	for i, p := range paths {
		solutions[i] = solveOne(p, nodes, lr)
	}

	//fmt.Println(solutions)
	return LCM(solutions[0], solutions[1], solutions[2:])
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
