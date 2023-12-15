package main

import (
	"fmt"
	"myadvent/internal"
	"regexp"
	"strconv"
)

func main() {

	var lastLine *schematicLine
	var notYetAdjacents []number
	sum := 0
	gears := []symbol{}

	for rawLine := range internal.ReadLines("day3") {
		line := parseLine(rawLine)
		adjacents, notAdjacents := checkAdjacent(line.numbers, line.symbols)
		for _, num := range adjacents {
			sum += num.value
		}
		adjacents, _ = checkAdjacent(notYetAdjacents, line.symbols)
		for _, num := range adjacents {
			sum += num.value
		}
		notYetAdjacents = notAdjacents
		if lastLine != nil {
			adjacents, _ = checkAdjacent(line.numbers, lastLine.symbols)
			gears = append(gears, filterGears(lastLine.symbols)...)
		}
		for _, num := range adjacents {
			sum += num.value
		}
		lastLine = &line
	}
	fmt.Println(sum)
	gears = append(gears, filterGears(lastLine.symbols)...)

	sumOfGearRatios := 0

	for _, gear := range gears {
		sumOfGearRatios += gear.adjacentNums[0] * gear.adjacentNums[1]
	}

	fmt.Println(sumOfGearRatios)
}

func filterGears(symbols []symbol) []symbol {
	gears := []symbol{}
	for _, sym := range symbols {
		if sym.value == "*" && len(sym.adjacentNums) == 2 {
			gears = append(gears, sym)
		}
	}
	return gears
}

func checkAdjacent(numbers []number, symbols []symbol) (adjacents, notAdjacents []number) {
	symbolsStartIndex := 0
numLoop:
	for _, num := range numbers {
		for i := symbolsStartIndex; i < len(symbols); i++ {
			if num.positionRangeStart > symbols[i].position {
				symbolsStartIndex = i + 1
			}
			if areAdjacent(num, symbols[i]) {
				adjacents = append(adjacents, num)
				symbols[i].adjacentNums = append(symbols[i].adjacentNums, num.value)
				continue numLoop
			}
		}
		notAdjacents = append(notAdjacents, num)
	}
	return
}

func areAdjacent(num number, sym symbol) bool {
	return sym.position >= num.positionRangeStart-1 && sym.position <= num.positionRangeEnd
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
	value        string
	position     int
	adjacentNums []int
}

func parseLine(line string) schematicLine {
	matchedRanges := numberOrSymbolPattern.FindAllStringIndex(line, -1)

	numbers, symbols := []number{}, []symbol{}
	for _, matchedRange := range matchedRanges {
		value := line[matchedRange[0]:matchedRange[1]]
		num, err := strconv.Atoi(value)
		if err == nil {
			numbers = append(numbers, number{value: num, positionRangeStart: matchedRange[0], positionRangeEnd: matchedRange[1]})
		} else {
			symbols = append(symbols, symbol{position: matchedRange[0], value: value})
		}
	}

	return schematicLine{numbers: numbers, symbols: symbols}
}
