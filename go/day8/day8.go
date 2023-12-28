package main

import (
	"fmt"
	"myadvent/internal"
	"slices"
	"strings"
)

func main() {
	linesChannel := internal.ReadLines("day8")
	directionsLine := <-linesChannel
	directions := parseDirections(directionsLine)
	<-linesChannel

	startingNodes := nodeList{}
	network := Network{}
	for line := range linesChannel {
		baseNode, connectedNodes := parseNodeLine(line)
		network[baseNode] = connectedNodes
		if baseNode.endsWithA() {
			startingNodes = append(startingNodes, baseNode)
		}
	}

	pathDescriptions := make(map[Node]PathDescription)
	for _, startingNode := range startingNodes {
		pathDescriptions[startingNode] = startingNode.getPathDescription(network, directions)
	}
	intervals := []int{}
	for _, description := range pathDescriptions {
		intervals = append(intervals, description.loopLength)
	}

	nSteps := leastCommonMultiple(intervals)

	fmt.Println(nSteps)

}
func primesUpTo(n int) []int {
	primes := []int{}
outer:
	for i := 2; i <= n; i++ {
		for _, prime := range primes {
			if i%prime == 0 {
				continue outer
			}
		}
		primes = append(primes, i)
	}
	return primes
}

func leastCommonMultiple(nums []int) int {
	max := slices.Max(nums)
	primes := primesUpTo(max)

	primeFactors := map[int]int{}
	for _, num := range nums {
		for _, prime := range primes {
			if prime > num {
				break
			}
			nPrimeFits := 0
			for num%prime == 0 {
				nPrimeFits++
				num /= prime
			}
			if primeFactors[prime] < nPrimeFits {
				primeFactors[prime] = nPrimeFits
			}
		}
	}
	multiple := 1
	for prime, count := range primeFactors {
		multiple *= prime * count
	}
	return multiple
}

func primesOf(num int) []int {
	primes := []int{}
	return primes
}

type PathDescription struct {
	freeLength   int
	loopLength   int
	endingPoints []int
}

type nodeList []Node

func (nl nodeList) endAllWithZ() bool {
	for _, n := range nl {
		if !n.endsWithZ() {
			return false
		}
	}
	return true
}

func (n Node) endsWithA() bool {
	return strings.HasSuffix(string(n), "A")
}
func (n Node) endsWithZ() bool {
	return strings.HasSuffix(string(n), "Z")
}

type Node string
type Network = map[Node][2]Node

type Step struct {
	currentNode    Node
	directionIndex int
}

func (n Node) getPathDescription(nw Network, directions []Direction) PathDescription {
	loopLength := 0
	currentNode := n
	stepsTaken := []Step{}
	endingPoints := []int{}
	var nextStep Step

	for ; ; loopLength++ {
		if currentNode.endsWithZ() {
			endingPoints = append(endingPoints, loopLength)
		}
		directionIndex := loopLength % len(directions)
		direction := directions[directionIndex]
		nextStep = Step{
			currentNode:    currentNode,
			directionIndex: directionIndex,
		}
		if slices.Contains(stepsTaken, nextStep) {
			break
		}
		stepsTaken = append(stepsTaken, nextStep)
		connectedNodes := nw[currentNode]
		currentNode = connectedNodes[direction]
	}
	freeLength := slices.Index(stepsTaken, nextStep)
	return PathDescription{
		freeLength:   freeLength,
		loopLength:   loopLength - freeLength,
		endingPoints: endingPoints,
	}
}

func parseNodeLine(line string) (base Node, connectedNodes [2]Node) {
	lineSplit := strings.Split(line, " = ")
	rawConnectedNodes := lineSplit[1]
	rawConnectedNodes = rawConnectedNodes[1 : len(rawConnectedNodes)-1]
	rawConnectedNodesSplit := strings.Split(rawConnectedNodes, ", ")

	return Node(lineSplit[0]), [2]Node{Node(rawConnectedNodesSplit[0]), Node(rawConnectedNodesSplit[1])}
}

type Direction int

const (
	left Direction = iota
	right
)

func parseDirections(line string) []Direction {
	directions := []Direction{}
	for _, symbol := range strings.Split(line, "") {
		if symbol == "L" {
			directions = append(directions, left)
		} else {
			directions = append(directions, right)
		}
	}

	return directions
}
