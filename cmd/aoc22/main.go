package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/yussiv/aoc21/pkg/mission"
	"github.com/yussiv/aoc21/util"
)

type Day interface {
	Task1(input []string) int
	Task2(input []string) int
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

func printTasksForDay(day int, input []string) {
	fmt.Printf("-- Day %d --\ntask 1: %d\ntask 2: %d\n",
		day,
		done[day].Task1(input),
		done[day].Task2(input),
	)
}

var done = map[int]Day{
	1: mission.Day1{},
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
			printTasksForDay(i, input)
		}
	} else if done[day] == nil {
		fmt.Println("I'm afraid the elves are on their own that day.")
	} else {
		input := util.GetInput(inputPath, day)
		printTasksForDay(day, input)
	}
}
