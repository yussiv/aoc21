package mission

import (
	"sort"
	"strconv"
	"strings"
)

type Day9 struct {
	input          []string
	floor          [][]int
	surveyed       [][]bool
	width          int
	height         int
	lowestPointSum int
	basins         []int
}

func (d *Day9) SetInput(input []string) {
	d.input = input
	d.parseInput()
	d.scanSeaFloor()
}

func (d *Day9) Task1() int {
	return d.lowestPointSum
}

func (d *Day9) Task2() int {
	return d.basins[0] * d.basins[1] * d.basins[2]
}

func (d *Day9) scanSeaFloor() {
	d.lowestPointSum = 0
	d.surveyed = make([][]bool, d.height)
	for i := range d.floor {
		d.surveyed[i] = make([]bool, d.width)
	}
	lowestPoints := make([]Coord, 0, d.height)
	for y, line := range d.floor {
		for x, val := range line {
			isLowest := d.isLower(val, x-1, y) &&
				d.isLower(val, x+1, y) &&
				d.isLower(val, x, y-1) &&
				d.isLower(val, x, y+1)
			if isLowest {
				lowestPoints = append(lowestPoints, Coord{X: x, Y: y})
				d.lowestPointSum += d.floor[y][x] + 1
			}
		}
	}
	d.basins = make([]int, len(lowestPoints))
	for i := range lowestPoints {
		d.basins[i] = d.surveyBasin(lowestPoints[i].X, lowestPoints[i].Y)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(d.basins)))
}

func (d *Day9) isLower(value, x, y int) bool {
	return x < 0 || y < 0 || x >= d.width || y >= d.height || d.floor[y][x] > value
}

func (d *Day9) surveyBasin(x, y int) int {
	if x < 0 || x >= d.width ||
		y < 0 || y >= d.height ||
		d.surveyed[y][x] || d.floor[y][x] == 9 {
		return 0
	}
	d.surveyed[y][x] = true
	sum := 1
	sum += d.surveyBasin(x-1, y)
	sum += d.surveyBasin(x+1, y)
	sum += d.surveyBasin(x, y-1)
	sum += d.surveyBasin(x, y+1)
	return sum
}

func (d *Day9) parseInput() {
	d.floor = make([][]int, len(d.input))
	for i, line := range d.input {
		strVals := strings.Split(strings.TrimSpace(line), "")
		d.floor[i] = make([]int, len(strVals))
		for j, s := range strVals {
			v, _ := strconv.Atoi(s)
			d.floor[i][j] = v
		}
	}
	d.height = len(d.floor)
	d.width = len(d.floor[0])
}
