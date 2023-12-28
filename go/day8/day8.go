package main

import (
	"fmt"
	"myadvent/internal"
	"strings"
)

func main() {
	linesChannel := internal.ReadLines("day8")
	directionsLine := <-linesChannel
	directions := parseDirections(directionsLine)
	<-linesChannel

	network := Network{}
	for line := range linesChannel {
		baseNode, connectedNodes := parseNodeLine(line)
		network[baseNode] = connectedNodes
	}

	nSteps := 0
	currentNode := node("AAA")
	for ; currentNode != node("ZZZ"); nSteps++ {
		direction := directions[nSteps%len(directions)]
		currentNode = network[currentNode][direction]
	}
	fmt.Println(nSteps)
}

type node string
type Network = map[node][2]node

func parseNodeLine(line string) (base node, connectedNodes [2]node) {
	lineSplit := strings.Split(line, " = ")
	rawConnectedNodes := lineSplit[1]
	rawConnectedNodes = rawConnectedNodes[1 : len(rawConnectedNodes)-1]
	rawConnectedNodesSplit := strings.Split(rawConnectedNodes, ", ")

	return node(lineSplit[0]), [2]node{node(rawConnectedNodesSplit[0]), node(rawConnectedNodesSplit[1])}
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
