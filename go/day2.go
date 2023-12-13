package main

import (
	"bufio"
	"fmt"
	"os"
)

type colorset struct {
	blue  int
	red   int
	green int
}

type game struct {
	id   int
	sets []colorset
}

func parseColorset(line string) *colorset {
	return &colorset{
		blue:  0,
		red:   0,
		green: 0,
	}
}

func parseGame(line string) *game {
	return &game{
		id:   0,
		sets: []colorset{},
	}
}

func main() {
	file, _ := os.Open("../inputs/day2")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		game := parseGame(scanner.Text())
		fmt.Println(game)
	}

	file.Close()
}
