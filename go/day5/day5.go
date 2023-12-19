package main

import (
	"fmt"
	"math"
	"myadvent/internal"
	"strconv"
	"strings"
)

func main() {
	linesChannel := internal.ReadLines("day5")
	seedsline := <-linesChannel
	transformer := intsTransformer{}
	transformer.current = matchSeedRanges(seedsline)

	for rawLine := range linesChannel {
		if rawLine == "" {
			transformer.finishStep()
			<-linesChannel
		} else {
			rang := matchRange(rawLine)
			transformer.mapRange(rang)
		}
	}
	transformer.finishStep()

	fmt.Println(min(transformer.current))

}

func min(nums []int) int {
	m := math.MaxInt
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}

type intsTransformer struct {
	next    []int
	current []int
}

func (it *intsTransformer) finishStep() {
	it.next = append(it.next, it.current...)
	it.current = it.next
	it.next = []int{}
}

func (it *intsTransformer) mapRange(r Range) {
	notMapped := []int{}
	for _, num := range it.current {
		newNum, err := r.mapNum(num)
		if err != nil {
			notMapped = append(notMapped, num)
		} else {
			it.next = append(it.next, newNum)
		}
	}
	it.current = notMapped
}

func matchSeeds(line string) []int {
	lineSplit := strings.Split(line, " ")
	return stringsToInts(lineSplit[1:])
}

func matchSeedRanges(line string) []int {
	lineSplit := strings.Split(line, " ")
	seeds := []int{}
	rawSeeds := lineSplit[1:]
	for i := range rawSeeds {
		if i%2 == 0 {
			rangeStart, _ := strconv.Atoi(rawSeeds[i])
			rangeLength, _ := strconv.Atoi(rawSeeds[i+1])
			for i := rangeStart; i <= rangeStart+rangeLength; i++ {
				seeds = append(seeds, i)
			}
		}
	}
	return seeds
}

type Range struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
}

type RangeMapError struct {
	num         int
	sourceStart int
	sourceEnd   int
}

func (r RangeMapError) Error() string {
	return "Could not map " + strconv.Itoa(r.num) + ": Not in range " + strconv.Itoa(r.sourceStart) + " - " + strconv.Itoa(r.sourceEnd)
}

func (r Range) mapNum(num int) (int, error) {
	if num < r.sourceStart || r.sourceEnd < num {
		return 0, RangeMapError{
			num:         num,
			sourceStart: r.sourceStart,
			sourceEnd:   r.sourceEnd,
		}
	}
	return num - r.sourceStart + r.destinationStart, nil
}

func matchRange(line string) Range {
	rawNums := strings.Split(line, " ")
	nums := stringsToInts(rawNums)
	return Range{
		destinationStart: nums[0],
		sourceStart:      nums[1],
		sourceEnd:        nums[1] + nums[2],
	}
}

func stringsToInts(strings []string) []int {
	ints := []int{}
	for _, rawNum := range strings {
		num, _ := strconv.Atoi(rawNum)
		ints = append(ints, num)
	}
	return ints
}
