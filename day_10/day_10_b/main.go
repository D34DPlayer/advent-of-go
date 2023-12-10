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
	if x != 0 && enterDirection != Down && currPipe[Up] && pipes[x-1][y][Down] {
		return x - 1, y, Up
	}
	if y < maxY && enterDirection != Left && currPipe[Right] && pipes[x][y+1][Left] {
		return x, y + 1, Right
	}
	if x < maxX && enterDirection != Up && currPipe[Down] && pipes[x+1][y][Up] {
		return x + 1, y, Down
	}
	if y != 0 && enterDirection != Right && currPipe[Left] && pipes[x][y-1][Right] {
		return x, y - 1, Left
	}
	panic(fmt.Sprintf("Couldn't find pipe after: %d, %d", x, y))
}

func newWallMap(maxX, maxY int) (wallMap [][]wall) {
	for i := 0; i <= maxX; i++ {
		wallMap = append(wallMap, make([]wall, maxY+1))
	}
	return
}

type wall struct {
	connectsUp   bool
	connectsDown bool
	isWall       bool
}

func newWall(p Pipe, enterDirection int) wall {
	if enterDirection == Up {
		return wall{true, false, true}
	} else if enterDirection == Down {
		return wall{false, true, true}
	}
	return wall{p[Up], p[Down] && !p[Up], true}
}
func wallDirection(w wall) int {
	if w.connectsUp {
		return Up
	} else if w.connectsDown {
		return Down
	}
	return -1
}

func scanWallRow(wallRow []wall) (count int) {
	isIn := false
	direction := -1

	for _, w := range wallRow {
		//switch {
		//case !w.isWall && isIn:
		//	fmt.Print("*")
		//case !w.isWall:
		//	fmt.Print(".")
		//case w.connectsUp && w.connectsDown:
		//	fmt.Print("|")
		//case w.connectsUp:
		//	fmt.Print("^")
		//case w.connectsDown:
		//	fmt.Print("v")
		//default:
		//	fmt.Print("-")
		//}

		if !w.isWall {
			if isIn {
				count++
			}
		} else {
			if (direction == -1 && wallDirection(w) != -1) ||
				(direction == Up && wallDirection(w) == Down) ||
				(direction == Down && wallDirection(w) == Up) {
				isIn = !isIn
				direction = wallDirection(w)
			}
		}
	}
	//fmt.Print("\n")

	return
}

func solve(pipes [][]Pipe, x, y int) (sum int) {
	enterDirection := -1
	maxX := len(pipes) - 1
	maxY := len(pipes[0]) - 1
	wallsMap := newWallMap(maxX, maxY)
	wallsMap[x][y] = newWall(pipes[x][y], enterDirection)
	for {
		x, y, enterDirection = nextPipe(pipes, x, y, enterDirection)
		wallsMap[x][y] = newWall(pipes[x][y], enterDirection)
		if isStart(pipes[x][y]) {
			break
		}
	}

	for _, w := range wallsMap {
		sum += scanWallRow(w)
		//fmt.Println(sum)
	}
	return
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
	fmt.Println("Solution:", solve(pipeMatrix, startCoords[0], startCoords[1]))
	fmt.Println("Time:", time.Since(start))
}
