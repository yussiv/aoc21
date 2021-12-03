package mission

import (
	"github.com/yussiv/aoc21/util"
)

type Day1 struct{ input []int }

func (d *Day1) SetInput(input []string) {
	d.input = util.StringsToInts(input)
}

func (d *Day1) Task1() int {
	return d.countIncreases(d.input)
}

func (d *Day1) Task2() int {
	windowed := d.windowedMeasurements(d.input)
	return d.countIncreases(windowed)
}

func (Day1) countIncreases(nums []int) int {
	count := 0
	for i := range nums[1:] {
		if nums[i] < nums[i+1] {
			count++
		}
	}
	return count
}

func (Day1) windowedMeasurements(measurements []int) []int {
	windowed := make([]int, len(measurements)-2)
	sum := measurements[0] + measurements[1]
	for i := range measurements[2:] {
		sum += measurements[i+2]
		windowed[i] = sum
		sum -= measurements[i]
	}
	return windowed
}
