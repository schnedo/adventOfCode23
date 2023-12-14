package main

import (
	"fmt"
	"myadvent/internal"
	"regexp"
	"strconv"
)

func main() {

	for rawLine := range internal.ReadLines("day3") {
		line := parseLine(rawLine)
		fmt.Println(line)
	}
}

var numberOrSymbolPattern = regexp.MustCompile("\\d+|[^.]")

type schematicLine struct {
	numbers []number
	symbols []symbol
}

type number struct {
	value              int
	positionRangeStart int
	positionRangeEnd   int
}

type symbol struct {
	position int
}

func parseLine(line string) schematicLine {
	matchedRanges := numberOrSymbolPattern.FindAllStringIndex(line, -1)

	numbers, symbols := []number{}, []symbol{}
	for _, matchedRange := range matchedRanges {
		num, err := strconv.Atoi(line[matchedRange[0]:matchedRange[1]])
		if err == nil {
			numbers = append(numbers, number{value: num, positionRangeStart: matchedRange[0], positionRangeEnd: matchedRange[1]})
		} else {
			symbols = append(symbols, symbol{position: matchedRange[0]})
		}
	}

	return schematicLine{numbers: numbers, symbols: symbols}
}
