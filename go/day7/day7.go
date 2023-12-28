package main

import (
	"fmt"
	"myadvent/internal"
	"slices"
	"strconv"
	"strings"
)

func main() {
	hands := []Hand{}
	for line := range internal.ReadLines("day7") {
		hand := parseHand(line)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, compareHands)

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	fmt.Println(totalWinnings)
}

type Card int16

const (
	two Card = iota
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

type Type int8

const (
	highCard Type = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type Hand struct {
	cards [5]Card
	bid   int
	ttype Type
}

func compareHands(a, b Hand) int {
	if a.ttype < b.ttype {
		return -1
	}
	if a.ttype > b.ttype {
		return 1
	}
	for i, aCard := range a.cards {
		bCard := b.cards[i]
		if aCard < bCard {
			return -1
		}
		if aCard > bCard {
			return 1
		}
	}
	panic("Is same hand even possible?")
}

func getTypeOf(cards [5]Card) Type {
	counts := make(map[Card]int8)
	for _, card := range cards {
		counts[card]++
	}
	switch len(counts) {
	case 1:
		return fiveOfAKind
	case 5:
		return highCard
	case 4:
		return onePair
	case 3:
		for _, count := range counts {
			if count == 3 {
				return threeOfAKind
			}
			if count == 2 {
				return twoPair
			}
		}
		panic("type of hand with 3 different cards seems fishy")
	case 2:
		for _, count := range counts {
			if count == 4 || count == 1 {
				return fourOfAKind
			} else {
				return fullHouse
			}
		}
	}
	panic("impossible count of cards received")
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
		ttype: getTypeOf(cards),
	}
}
