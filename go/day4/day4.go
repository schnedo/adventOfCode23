package main

import (
	"fmt"
	"myadvent/internal"
	"regexp"
	"strconv"
)

func main() {
	sum := 0
	copies := Copies{nCopies: []int{}}
	nCards := 0

	for rawLine := range internal.ReadLines("day4") {
		card := parseCard(rawLine)
		sum += card.points()

		nCopies := copies.pop()
		nCards += nCopies + 1
		for i := 0; i <= nCopies; i++ {
			copies.addNCopies(card.nMyWinningNumbers())
		}

	}
	fmt.Println(sum)
	fmt.Println(nCards)
}

type Copies struct {
	nCopies []int
}

func (c *Copies) addNCopies(count int) {
	for i := len(c.nCopies); i < count; i++ {
		c.nCopies = append(c.nCopies, 0)
	}
	for i := 0; i < count; i++ {
		c.nCopies[i]++
	}
}

func (c *Copies) pop() int {
	if len(c.nCopies) == 0 {
		return 0
	}
	popped := c.nCopies[0]
	c.nCopies = c.nCopies[1:]
	return popped
}

type Card struct {
	id             int
	winningNumbers [10]int
	myNumbers      [25]int
}

func (card Card) points() int {
	nMyWinningNumbers := card.nMyWinningNumbers()
	if nMyWinningNumbers == 0 {
		return 0
	}
	points := 1
	for i := 1; i < nMyWinningNumbers; i++ {
		points *= 2
	}
	return points
}

func (card Card) myWinningNumbers() []int {
	myWinningNumbers := []int{}
	for _, num := range card.myNumbers {
		if card.isWinningNumber(num) {
			myWinningNumbers = append(myWinningNumbers, num)
		}
	}
	return myWinningNumbers
}

func (card Card) nMyWinningNumbers() int {
	return len(card.myWinningNumbers())
}

func (card Card) isWinningNumber(num int) bool {
	for _, winning := range card.winningNumbers {
		if num == winning {
			return true
		}
	}
	return false
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
