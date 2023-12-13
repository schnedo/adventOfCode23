package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type colorset struct {
	blue  int
	red   int
	green int
}

func (c *colorset) String() string {
	return "{red: " + strconv.Itoa(c.red) + ", blue: " + strconv.Itoa(c.blue) + ", green: " + strconv.Itoa(c.green) + "}"
}

type game struct {
	id   int
	sets []colorset
}

func (g *game) String() string {
	colorsets := ""
	for i, set := range g.sets {
		colorsets = colorsets + set.String()
		if i != len(g.sets)-1 {
			colorsets = colorsets + " "
		}
	}
	return "{" + strconv.Itoa(g.id) + " [" + colorsets + "]}"
}

func parseColorset(line string) colorset {
	lineSplit := strings.Split(line, ",")

	colorMap := map[string]int{}

	for _, setString := range lineSplit {
		colorsetSplit := strings.Split(setString, " ")
		nColor, _ := strconv.Atoi(colorsetSplit[len(colorsetSplit)-2])
		colorMap[colorsetSplit[len(colorsetSplit)-1]] = nColor
	}

	return colorset{
		blue:  colorMap["blue"],
		red:   colorMap["red"],
		green: colorMap["green"],
	}
}

func parseGame(line string) *game {
	lineSplit := strings.Split(line, ":")
	gameSplit := strings.Split(lineSplit[0], " ")

	gameId, _ := strconv.Atoi(gameSplit[1])

	colorsetSplit := strings.Split(lineSplit[1], ";")

	sets := []colorset{}
	for _, line := range colorsetSplit {
		sets = append(sets, parseColorset(line))
	}

	return &game{
		id:   gameId,
		sets: sets,
	}
}

func main() {
	file, _ := os.Open("../inputs/day2")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		game := parseGame(scanner.Text())
		fmt.Println(game)
	}

	file.Close()
}
