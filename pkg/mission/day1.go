package mission

import (
	"github.com/yussiv/aoc21/util"
)

type Day1 struct{}

func (Day1) Task1(input []string) int {
	nums := getInput(input)
	return countIncreases(nums)
}

func (Day1) Task2(input []string) int {
	nums := getInput(input)
	windowed := windowedMeasurements(nums)
	return countIncreases(windowed)
}

func getInput(input []string) []int {
	return util.StringsToInts(input)
}

func countIncreases(nums []int) int {
	count := 0
	for i := range nums[1:] {
		if nums[i] < nums[i+1] {
			count++
		}
	}
	return count
}

func windowedMeasurements(measurements []int) []int {
	windowed := make([]int, len(measurements)-2)
	sum := measurements[0] + measurements[1]
	for i := range measurements[2:] {
		sum += measurements[i+2]
		windowed[i] = sum
		sum -= measurements[i]
	}
	return windowed
}
