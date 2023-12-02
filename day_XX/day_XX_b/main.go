package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(f) // Scanner splits on '\n' by default
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
}
