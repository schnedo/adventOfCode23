package main

import (
	"fmt"
	"myadvent/internal"
	"regexp"
	"strconv"
)

func main() {
	for rawLine := range internal.ReadLines("day4") {
		card := parseCard(rawLine)
		fmt.Println(card)
	}
}

type Card struct {
	id             int
	winningNumbers [10]int
	myNumbers      [25]int
}

var cardPattern = regexp.MustCompile("Card +(\\d+): +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) \\| +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+) +(\\d+)")

func parseCard(line string) Card {
	allMatches := cardPattern.FindAllStringSubmatch(line, -1)
	matches := allMatches[0]
	id, _ := strconv.Atoi(matches[1])

	winningNumbers := [10]int{}
	for i, rawNum := range matches[2:12] {
		num, _ := strconv.Atoi(rawNum)
		winningNumbers[i] = num
	}

	myNumbers := [25]int{}
	for i, rawNum := range matches[12:] {
		num, _ := strconv.Atoi(rawNum)
		myNumbers[i] = num
	}

	return Card{
		id:             id,
		winningNumbers: winningNumbers,
		myNumbers:      myNumbers,
	}
}
