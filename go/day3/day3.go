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
	for rawLine := range internal.ReadLines("day3") {
		fmt.Println("----------------")
		fmt.Println("raw:")
		fmt.Println(rawLine)
		line := parseLine(rawLine)
		adjacents, notAdjacents := checkAdjacent(line.numbers, line.symbols)
		fmt.Println("adjacents:")
		for _, num := range adjacents {
			fmt.Print(strconv.Itoa(num.value) + ",")
			sum += num.value
		}
		adjacents, _ = checkAdjacent(notYetAdjacents, line.symbols)
		for _, num := range adjacents {
			fmt.Print(strconv.Itoa(num.value) + ",")
			sum += num.value
		}
		notYetAdjacents = notAdjacents
		if lastLine != nil {
			adjacents, _ = checkAdjacent(line.numbers, lastLine.symbols)
		}
		for _, num := range adjacents {
			fmt.Print(strconv.Itoa(num.value) + ",")
			sum += num.value
		}
		lastLine = &line
		fmt.Println()
	}
	fmt.Println(sum)
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
