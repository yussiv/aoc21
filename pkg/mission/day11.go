package mission

import (
	"strconv"
	"strings"
)

type Day11 struct {
	input    []string
	energies [][]int
	width    int
	height   int
}

func (d *Day11) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day11) Task1() int {
	flashCount := 0
	for cycle := 0; cycle < 100; cycle++ {
		flashCount += d.simulateCycle()
	}
	return flashCount
}

func (d *Day11) Task2() int {
	cycle := 100 // start where first part ended
	for {
		flashCount := d.simulateCycle()
		cycle++
		if flashCount == d.height*d.width {
			break
		}
	}
	return cycle
}

func (d *Day11) simulateCycle() int {
	flashCount := 0
	flashed := make([][]bool, d.height)
	for i := range flashed {
		flashed[i] = make([]bool, d.width)
	}
	willFlash := make([]Coord, 0)
	for y, row := range d.energies {
		for x, val := range row {
			d.energies[y][x] = val + 1
			if d.energies[y][x] > 9 {
				willFlash = append(willFlash, Coord{x, y})
			}
		}
	}

	for n := 0; n < len(willFlash); n++ {
		x, y := willFlash[n].X, willFlash[n].Y
		if flashed[y][x] {
			continue
		}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if x+i >= 0 && x+i < d.width &&
					y+j >= 0 && y+j < d.height &&
					!flashed[y+j][x+i] {
					d.energies[y+j][x+i]++
					if d.energies[y+j][x+i] > 9 {
						willFlash = append(willFlash, Coord{x + i, y + j})
					}
				}
			}
		}
		flashed[y][x] = true
		flashCount++
		d.energies[y][x] = 0
	}
	return flashCount
}

func (d *Day11) parseInput() {
	d.energies = make([][]int, len(d.input))
	for i, row := range d.input {
		strVals := strings.Split(row, "")
		d.energies[i] = make([]int, len(strVals))
		for j, s := range strVals {
			v, _ := strconv.Atoi(s)
			d.energies[i][j] = v
		}
	}
	d.height = len(d.energies)
	d.width = len(d.energies[0])
}
