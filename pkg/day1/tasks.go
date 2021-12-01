package day1

import (
	"github.com/yussiv/aoc21/util"
)

func getInput() []int {
	rows := util.ReadLines("./input/day1")
	return util.StringsToInts(rows)
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
	windowed := make([]int, len(measurements)-1)
	sum := measurements[0] + measurements[1]
	for i := range measurements[2:] {
		sum += measurements[i+2]
		windowed[i] = sum
		sum -= measurements[i]
	}
	return windowed
}

func Task1() int {
	nums := getInput()
	return countIncreases(nums)
}

func Task2() int {
	nums := getInput()
	windowed := windowedMeasurements(nums)
	return countIncreases(windowed)
}
