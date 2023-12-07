package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Hand int

const (
	Undefined Hand = iota
	FiveOfAKind
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	Pair
	HighCard
)

type game struct {
	cards []int
	bid   int
	hand  Hand
}

func (g game) calcHand() Hand {
	cards := make(map[int]int)
	jokers := 0
	for _, c := range g.cards {
		if c == 1 {
			jokers++
			continue
		}
		n, ok := cards[c]
		if !ok {
			n = 0
		}
		cards[c] = n + 1
	}

	switch {
	case isFiveOfAKind(cards, jokers):
		return FiveOfAKind
	case isFourOfAKind(cards, jokers):
		return FourOfAKind
	case isFullHouse(cards, jokers):
		return FullHouse
	case isThreeOfAKind(cards, jokers):
		return ThreeOfAKind
	case isTwoPair(cards, jokers):
		return TwoPair
	case isPair(cards, jokers):
		return Pair
	default:
		return HighCard
	}
}

func isFiveOfAKind(n map[int]int, _ int) bool {
	return len(n) <= 1
}

func isFourOfAKind(n map[int]int, jokers int) bool {
	for _, v := range n {
		if v >= 4-jokers {
			return true
		}
	}
	return false
}

func isThreeOfAKind(n map[int]int, jokers int) bool {
	for _, v := range n {
		if v >= 3-jokers {
			return true
		}
	}
	return false
}

func isPair(n map[int]int, jokers int) bool {
	for _, v := range n {
		if v >= 2-jokers {
			return true
		}
	}
	return false
}

func isTwoPair(n map[int]int, jokers int) bool {
	count := 0
	for _, v := range n {
		if v >= 2-jokers {
			count++
		}
	}
	return count == 2
}

func isFullHouse(n map[int]int, _ int) bool {
	return len(n) == 2
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

func mapCard(r rune) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J': // JOKERS
		return 1
	case 'T':
		return 10
	default:
		i, err := strconv.Atoi(string(r))

		if err != nil {
			panic(err)
		}
		return i
	}
}

func (p *parser) parseCards() []int {
	r, _ := p.current()
	out := []int{mapCard(r)}
	for r, ok := p.next(); ok; r, ok = p.next() {
		if r == ' ' {
			p.next()
			break
		}
		out = append(out, mapCard(r))
	}
	return out
}

func (p *parser) parseGame() game {
	p.skip()
	cards := p.parseCards()
	p.skip()
	bid := p.parseInt()
	return game{cards, bid, Undefined}
}

func main() {
	start := time.Now()
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	games := make([]game, 0)
	for scanner.Scan() {
		line := scanner.Text()
		p := newParser(line)
		g := p.parseGame()
		g.hand = g.calcHand()
		//if strings.ContainsRune(line, 'J') {
		//	fmt.Println("Line:", line)
		//	fmt.Println("Game:", g)
		//}
		games = append(games, g)
	}
	sort.Slice(games, func(i, j int) bool {
		if games[i].hand != games[j].hand {
			return games[i].hand > games[j].hand
		} else {
			for idx, iCard := range games[i].cards {
				jCard := games[j].cards[idx]
				if iCard != jCard {
					return iCard < jCard
				}
			}
			return false
		}
	})

	sum := 0
	for i, g := range games {
		sum += g.bid * (i + 1)
	}

	fmt.Println(sum)
	fmt.Println("Time:", time.Since(start))
}
