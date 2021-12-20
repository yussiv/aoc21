package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/yussiv/aoc21/pkg/mission"
	"github.com/yussiv/aoc21/util"
)

type Day interface {
	Task1() int
	Task2() int
	SetInput(input []string)
}

func getSortedMapKeys(dayMap map[int]Day) []int {
	keys := make([]int, len(dayMap))
	i := 0
	for key := range dayMap {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	return keys
}

func runTasksForDay(day int, input []string) {
	done[day].SetInput(input)
	fmt.Printf("-- Day %d --\ntask 1: %d\ntask 2: %d\n\n",
		day,
		done[day].Task1(),
		done[day].Task2(),
	)
}

var done = map[int]Day{
	1:  &mission.Day1{},
	2:  &mission.Day2{},
	3:  &mission.Day3{},
	4:  &mission.Day4{},
	5:  &mission.Day5{},
	6:  &mission.Day6{},
	7:  &mission.Day7{},
	8:  &mission.Day8{},
	9:  &mission.Day9{},
	10: &mission.Day10{},
	11: &mission.Day11{},
	12: &mission.Day12{},
	13: &mission.Day13{},
	14: &mission.Day14{},
	15: &mission.Day15{},
	16: &mission.Day16{},
	17: &mission.Day17{},
}

func main() {
	var day int
	var inputPath string
	flag.IntVar(&day, "d", 0, "Day of tasks you want to run. If omitted, runs all implemented days.")
	flag.StringVar(&inputPath, "i", "", "Input file path. If omitted, will try to open './input/day<number of day>'.")
	flag.Parse()

	if day == 0 {
		keys := getSortedMapKeys(done)
		for _, i := range keys {
			input := util.GetInput("", i) // TODO: input directory flag
			runTasksForDay(i, input)
		}
	} else if done[day] == nil {
		fmt.Println("I'm afraid the elves are on their own that day.")
	} else {
		input := util.GetInput(inputPath, day)
		runTasksForDay(day, input)
	}
}
