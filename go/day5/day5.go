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

	// seedRanges := matchSeedRangesPartOne(seedsline)
	seedRanges := matchSeedRanges(seedsline)

	<-linesChannel
	<-linesChannel
	seedMap := SeedMap{}

	mapLayer := []Range{}
	for rawLine := range linesChannel {
		if rawLine == "" {
			seedMap.rangeLayers = append(seedMap.rangeLayers, mapLayer)
			mapLayer = []Range{}
			<-linesChannel
		} else {
			r := matchRange(rawLine)
			mapLayer = append(mapLayer, r)
		}
	}
	seedMap.rangeLayers = append(seedMap.rangeLayers, mapLayer)

	mappedRanges := seedMap.applyTo(seedRanges)

	currentMinLocation := math.MaxInt
	for _, r := range mappedRanges {
		if r.start < currentMinLocation {
			currentMinLocation = r.start
		}
	}

	fmt.Println(currentMinLocation)

}

type SeedMap struct {
	rangeLayers [][]Range
}

func (sm SeedMap) applyTo(ranges []SeedRange) []SeedRange {
	var mapped []SeedRange
	notYetMapped := ranges
	for _, layer := range sm.rangeLayers {
		mapped = []SeedRange{}
		for _, r := range layer {
			singleMapped, singleUnmapped := r.applyToRanges(notYetMapped)
			mapped = append(mapped, singleMapped...)
			notYetMapped = singleUnmapped
		}
		notYetMapped = append(notYetMapped, mapped...)
	}
	return notYetMapped
}

func (sm SeedMap) inverseMapNum(num int) int {
	mappedNumber := num
	for i := len(sm.rangeLayers) - 1; i >= 0; i-- {
		for _, r := range sm.rangeLayers[i] {
			newNum, err := r.inverseMapNum(mappedNumber)
			if err == nil {
				mappedNumber = newNum
				break
			}
		}
	}
	return mappedNumber
}

func (sm SeedMap) mapNum(num int) int {
	mappedNumber := num
	for _, layer := range sm.rangeLayers {
		for _, r := range layer {
			newNum, err := r.mapNum(mappedNumber)
			if err == nil {
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

type SeedRange struct {
	start int
	end   int
}

func (sr SeedRange) includes(num int) bool {
	return sr.start <= num && num < sr.end
}

func matchSeedRangesPartOne(line string) []SeedRange {
	lineSplit := strings.Split(line, " ")
	rawSeeds := lineSplit[1:]
	seedRanges := []SeedRange{}

	for _, rawSeed := range rawSeeds {
		start, _ := strconv.Atoi(rawSeed)
		seedRanges = append(seedRanges, SeedRange{
			start: start,
			end:   start + 1,
		})
	}

	return seedRanges
}

func matchSeedRanges(line string) []SeedRange {
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
	return seedRanges
}

type Range struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
}

type RangeMapError struct {
}

func (r Range) applyToRanges(seedRanges []SeedRange) (mapped, unmapped []SeedRange) {
	for _, seedRange := range seedRanges {
		singleMapped, singleUnmapped := r.applyTo(seedRange)
		mapped = append(mapped, singleMapped...)
		unmapped = append(unmapped, singleUnmapped...)
	}
	return
}

func (r Range) applyTo(s SeedRange) (mapped, unmapped []SeedRange) {
	if r.sourceEnd <= s.start || r.sourceStart >= s.end {
		unmapped = append(unmapped, s)
		return
	}
	var mappedStart, mappedEnd int
	if r.sourceStart < s.start {
		m, _ := r.mapNum(s.start)
		mappedStart = m
	} else {
		mappedStart = r.destinationStart
		if r.sourceStart > s.start {
			unmapped = append(unmapped, SeedRange{
				start: s.start,
				end:   r.sourceStart,
			})
		}
	}
	if r.sourceEnd > s.end {
		m, _ := r.mapNum(s.end)
		mappedEnd = m
	} else {
		mappedEnd = r.destinationEnd
		if r.sourceEnd < s.end {
			unmapped = append(unmapped, SeedRange{
				start: r.sourceEnd,
				end:   s.end,
			})
		}
	}

	mapped = append(mapped, SeedRange{
		start: mappedStart,
		end:   mappedEnd,
	})
	return
}

func (r RangeMapError) Error() string {
	return "Could not map range"
}

func (r Range) inverseMapNum(num int) (int, error) {
	if num < r.destinationStart || r.destinationEnd <= num {
		return 0, RangeMapError{}
	}
	return num - r.destinationStart + r.sourceStart, nil
}

func (r Range) mapNum(num int) (int, error) {
	if num < r.sourceStart || r.sourceEnd <= num {
		return 0, RangeMapError{}
	}
	return num - r.sourceStart + r.destinationStart, nil
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
