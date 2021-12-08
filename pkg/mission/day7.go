package mission

import (
	"strconv"
	"strings"

	"github.com/yussiv/aoc21/util"
)

type Day7 struct {
	input     []string
	positions []int
	sums      []int
}

func (d *Day7) SetInput(input []string) {
	d.input = input
	d.parseInput()
	d.sums = make([]int, 2000)
	for i := 1; i < len(d.sums); i++ {
		d.sums[i] = d.sums[i-1] + i
	}
}

func (d *Day7) Task1() int {
	min := int(^uint(0) >> 1)
	max := 0
	for _, pos := range d.positions {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}
	sum := 0
	mmin := int(^uint(0) >> 1)
	for i := min; i <= max; i++ {
		sum = 0
		for _, pos := range d.positions {
			sum += util.Abs(i - pos)
		}
		if sum < mmin {
			mmin = sum
		}
	}
	return mmin
}

func (d *Day7) Task2() int {
	min := int(^uint(0) >> 1)
	max := 0
	for _, pos := range d.positions {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}
	sum := 0
	mmin := int(^uint(0) >> 1)
	for i := min; i <= max; i++ {
		sum = 0
		for _, pos := range d.positions {
			sum += d.sums[util.Abs(i-pos)]
		}
		if sum < mmin {
			mmin = sum
		}
	}
	return mmin
}

func (d *Day7) parseInput() {
	posStrs := strings.Split(d.input[0], ",")
	d.positions = make([]int, len(posStrs))
	for i, s := range posStrs {
		n, _ := strconv.Atoi(s)
		d.positions[i] = n
	}
	// d.positions = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
}
