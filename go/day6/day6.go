package main

import "fmt"

type Race struct {
	time     int
	distance int
}

func main() {
	races := []Race{
		{
			time:     48,
			distance: 261,
		},
		{
			time:     93,
			distance: 1192,
		},
		{
			time:     84,
			distance: 1019,
		},
		{
			time:     66,
			distance: 1063,
		},
	}

	nWaysToWinPerRace := []int{}
	for _, race := range races {
		nWaysToWin := 0
		for i := 1; i < race.time; i++ {
			if i*(race.time-i) > race.distance {
				nWaysToWin++
			}
		}
		nWaysToWinPerRace = append(nWaysToWinPerRace, nWaysToWin)
	}

	mul := 1
	for _, n := range nWaysToWinPerRace {
		mul *= n
	}
	fmt.Println(mul)

}
