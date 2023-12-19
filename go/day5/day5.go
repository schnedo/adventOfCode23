package main

import (
	"fmt"
	"myadvent/internal"
	"strconv"
	"strings"
)

func main() {
	linesChannel := internal.ReadLines("day5")
	seedsline := <-linesChannel
	seedRanges := matchSeedRanges(seedsline)
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

	currentMin := 0
	for ; !seedRanges.includes(seedMap.inverseMapNum(currentMin)); currentMin++ {
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

func (sm SeedMap) inverseMapNum(num int) int {
	mappedNumber := num
	for i := len(sm.rangeLayers) - 1; i >= 0; i-- {
		for _, r := range sm.rangeLayers[i] {
			newNum, err := r.inverseMapNum(mappedNumber)
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

type SeedRanges struct {
	ranges []SeedRange
}

func (sr SeedRanges) includes(num int) bool {
	for _, r := range sr.ranges {
		if r.includes(num) {
			return true
		}
	}
	return false
}

type SeedRange struct {
	start int
	end   int
}

func (sr SeedRange) includes(num int) bool {
	return sr.start <= num && num < sr.end
}

func matchSeedRanges(line string) SeedRanges {
	lineSplit := strings.Split(line, " ")
	rawSeeds := lineSplit[1:]
	seedRanges := []SeedRange{}

	for i := range rawSeeds {
		if i%2 == 0 {
			start, _ := strconv.Atoi(rawSeeds[i])
			length, _ := strconv.Atoi(rawSeeds[i+1])
			seedRanges = append(seedRanges, SeedRange{start: start, end: start + length})
		}
	}
	return SeedRanges{ranges: seedRanges}
}

type Range struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
}

type RangeMapError struct {
	num   int
	start int
	end   int
}

func (r RangeMapError) Error() string {
	return "Could not map " + strconv.Itoa(r.num) + ": Not in range " + strconv.Itoa(r.start) + " - " + strconv.Itoa(r.end)
}

func (r Range) mapNum(num int) (int, error) {
	if num < r.sourceStart || r.sourceEnd <= num {
		return 0, RangeMapError{
			num:   num,
			start: r.sourceStart,
			end:   r.sourceEnd,
		}
	}
	return num - r.sourceStart + r.destinationStart, nil
}

func (r Range) inverseMapNum(num int) (int, error) {
	if num < r.destinationStart || r.destinationEnd <= num {
		return 0, RangeMapError{
			num:   num,
			start: r.destinationStart,
			end:   r.destinationEnd,
		}
	}
	return num - r.destinationStart + r.sourceStart, nil
}

func matchRange(line string) Range {
	rawNums := strings.Split(line, " ")
	nums := stringsToInts(rawNums)
	return Range{
		destinationStart: nums[0],
		destinationEnd:   nums[0] + nums[2],
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
