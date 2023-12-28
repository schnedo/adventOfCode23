package main

import (
	"fmt"
	"myadvent/internal"
	"strconv"
	"strings"
)

func main() {

	sum := 0
	for line := range internal.ReadLines("day9") {
		sequence := parseSequence(line)
		nextValue := extrapolate(sequence)
		sum += nextValue
	}
	fmt.Println(sum)

}

func extrapolate(sequence []int) int {
	sequences := [][]int{sequence}
	for currentSequence := sequence; !isAllZeroes(currentSequence); currentSequence = sequences[len(sequences)-1] {
		sequences = append(sequences, calcSequenceOfDiffs(currentSequence))
	}

	lastSequenceIndex := len(sequences) - 1

	for i := lastSequenceIndex - 1; i >= 0; i-- {
		previousSequence := sequences[i+1]
		currentSequence := sequences[i]
		nextValue := lastOf(currentSequence) + lastOf(previousSequence)
		currentSequence = append(currentSequence, nextValue)
		sequences[i] = currentSequence
	}

	return lastOf(sequences[0])
}

func lastOf(sequence []int) int {
	return sequence[len(sequence)-1]
}

func isAllZeroes(in []int) bool {
	for _, num := range in {
		if num != 0 {
			return false
		}
	}
	return true
}

func parseSequence(line string) []int {
	sequence := []int{}
	rawSequence := strings.Split(line, " ")
	for _, rawNum := range rawSequence {
		num, _ := strconv.Atoi(rawNum)
		sequence = append(sequence, num)
	}
	return sequence
}

func calcSequenceOfDiffs(in []int) []int {
	sequenceOfDiffs := []int{}
	for i := 0; i < len(in)-1; i++ {
		sequenceOfDiffs = append(sequenceOfDiffs, in[i+1]-in[i])
	}
	return sequenceOfDiffs
}
