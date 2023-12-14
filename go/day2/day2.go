package main

import (
	"fmt"
	"myadvent/internal"
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

func (c *colorset) isMoreThan(other *colorset) bool {
	return c.blue > other.blue || c.green > other.green || c.red > other.red
}

func (c *colorset) power() int {
	return c.blue * c.red * c.green
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

func (g *game) isPossibleWith(bag *colorset) bool {
	for _, set := range g.sets {
		if set.isMoreThan(bag) {
			return false
		}
	}
	return true
}

func enclosingSet(sets []colorset) *colorset {
	minBlue := 0
	minRed := 0
	minGreen := 0

	for _, set := range sets {
		if set.blue > minBlue {
			minBlue = set.blue
		}
		if set.red > minRed {
			minRed = set.red
		}
		if set.green > minGreen {
			minGreen = set.green
		}
	}

	return &colorset{
		green: minGreen,
		red:   minRed,
		blue:  minBlue,
	}
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
	lines := internal.ReadLines("day2")

	bag := &colorset{red: 12, green: 13, blue: 14}
	sumOfPossibles := 0
	sumOfEnclosingPowers := 0

	for line := range lines {
		game := parseGame(line)
		if game.isPossibleWith(bag) {
			sumOfPossibles += game.id
		}
		sumOfEnclosingPowers += enclosingSet(game.sets).power()
	}

	fmt.Println("sumOfPossibles: " + strconv.Itoa(sumOfPossibles))
	fmt.Println("sumOfEnclosingPowers: " + strconv.Itoa(sumOfEnclosingPowers))
}
