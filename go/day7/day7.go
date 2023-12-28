package main

import (
	"fmt"
	"myadvent/internal"
	"strconv"
	"strings"
)

func main() {
	hands := []Hand{}
	for line := range internal.ReadLines("day7") {
		hand := parseHand(line)
		hands = append(hands, hand)
	}

	fmt.Println(hands)
}

type Card int8

const (
	two Card = 2 + iota
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

type Hand struct {
	cards [5]Card
	bid   int
}

func parseCard(label string) Card {
	switch label {
	case "2":
		return two
	case "3":
		return three
	case "4":
		return four
	case "5":
		return five
	case "6":
		return six
	case "7":
		return seven
	case "8":
		return eight
	case "9":
		return nine
	case "T":
		return ten
	case "J":
		return jack
	case "Q":
		return queen
	case "K":
		return king
	case "A":
		return ace
	}
	panic("Could not parse card")
}

func parseHand(line string) Hand {
	linesplit := strings.Split(line, " ")
	bid, _ := strconv.Atoi(linesplit[1])
	labels := strings.Split(linesplit[0], "")
	cards := [5]Card{}

	for i, label := range labels {
		cards[i] = parseCard(label)
	}

	return Hand{
		cards: cards,
		bid:   bid,
	}
}
