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
	<-linesChannel
	<-linesChannel
	seedMap := SeedMap{}

	mapLayer := []Range{}
	for rawLine := range linesChannel {
		if rawLine == "" {
			seedMap.rangeLayers = append(seedMap.rangeLayers, mapLayer)
			<-linesChannel
		} else {
			r := matchRange(rawLine)
			mapLayer = append(mapLayer, r)
		}
	}
	seedMap.rangeLayers = append(seedMap.rangeLayers, mapLayer)

	currentMin := math.MaxInt
	for seed := range matchSeedRanges(seedsline) {
		mappedSeed := seedMap.mapNum(seed)
		if mappedSeed < currentMin {
			currentMin = mappedSeed
		}
	}

	fmt.Println(currentMin)

}

type SeedMap struct {
	rangeLayers [][]Range
}

func (sm SeedMap) mapNum(num int) int {
	mappedNumber := num
	for _, layer := range sm.rangeLayers {
		for _, r := range layer {
			newNum, err := r.mapNum(mappedNumber)
			if err != nil {
				mappedNumber = newNum
				break
			}
		}
	}
	return mappedNumber
}

func matchSeeds(line string) []int {
	lineSplit := strings.Split(line, " ")
	return stringsToInts(lineSplit[1:])
}

func seedRanges(line string, c chan int) {
	lineSplit := strings.Split(line, " ")
	rawSeeds := lineSplit[1:]
	for i := range rawSeeds {
		if i%2 == 0 {
			rangeStart, _ := strconv.Atoi(rawSeeds[i])
			rangeLength, _ := strconv.Atoi(rawSeeds[i+1])
			for i := rangeStart; i <= rangeStart+rangeLength; i++ {
				c <- i
			}
		}
	}
}

func matchSeedRanges(line string) chan int {
	seedChannel := make(chan int)
	go seedRanges(line, seedChannel)
	return seedChannel
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
	if num < r.sourceStart || r.sourceEnd <= num {
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
