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

const (
	Up int = iota
	Right
	Down
	Left
)

type Pipe = [4]bool

func newPipe(r rune) Pipe {
	switch r {
	case '7':
		return Pipe{false, false, true, true}
	case '|':
		return Pipe{true, false, true, false}
	case 'J':
		return Pipe{true, false, false, true}
	case '-':
		return Pipe{false, true, false, true}
	case 'L':
		return Pipe{true, true, false, false}
	case 'F':
		return Pipe{false, true, true, false}
	case '.':
		return Pipe{false, false, false, false}
	case 'S':
		return Pipe{true, true, true, true}
	default:
		panic("Unexpected pipe character: " + string(r))
	}
}

func isStart(p Pipe) bool {
	for _, d := range p {
		if !d {
			return false
		}
	}
	return true
}

func nextPipe(pipes [][]Pipe, x, y, enterDirection int) (int, int, int) {
	currPipe := pipes[x][y]
	maxX := len(pipes) - 1
	maxY := len(pipes[0]) - 1
	if x != 0 && enterDirection != Up && currPipe[Up] && pipes[x-1][y][Down] {
		return x - 1, y, Down
	}
	if y < maxY && enterDirection != Right && currPipe[Right] && pipes[x][y+1][Left] {
		return x, y + 1, Left
	}
	if x < maxX && enterDirection != Down && currPipe[Down] && pipes[x+1][y][Up] {
		return x + 1, y, Up
	}
	if y != 0 && enterDirection != Left && currPipe[Left] && pipes[x][y-1][Right] {
		return x, y - 1, Right
	}
	panic(fmt.Sprintf("Couldn't find pipe after: %d, %d", x, y))
}

func loopLength(pipes [][]Pipe, x, y int) int {
	count := 0
	enterDirection := -1
	for {
		count++
		x, y, enterDirection = nextPipe(pipes, x, y, enterDirection)
		fmt.Println("Next:", x, y)
		if isStart(pipes[x][y]) {
			break
		}
	}
	return count
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

func (p *parser) parsePipe() Pipe {
	r, _ := p.current()
	defer p.next()
	return newPipe(r)
}

func (p *parser) parsePipes() (pipes []Pipe, startIndex int) {
	startIndex = -1
	i := 0
	for r, ok := p.current(); ok; r, ok = p.next() {
		if r == 'S' {
			startIndex = i
		}
		pipes = append(pipes, newPipe(r))
		i++
	}
	return
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	var pipeMatrix [][]Pipe
	startCoords := [2]int{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		p := newParser(line)
		pipes, j := p.parsePipes()
		pipeMatrix = append(pipeMatrix, pipes)
		//fmt.Println("Line:", line)
		//fmt.Println("Pipes:", pipes)

		if j != -1 {
			startCoords = [2]int{i, j}
		}
		i++
	}
	//fmt.Println("Start coords:", startCoords)
	fmt.Println("Solution:", loopLength(pipeMatrix, startCoords[0], startCoords[1])/2)
	fmt.Println("Time:", time.Since(start))
}
