package mission

import (
	"fmt"
)

type Day20 struct {
	input         []string
	enhancement   []bool
	image         map[Coord]bool
	min           Coord
	max           Coord
	step          int
	infinityIsLit bool
}

func (d *Day20) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day20) Task1() int {
	d.simulateImageEnhancement(2)
	return len(d.image)
}

func (d *Day20) Task2() int {
	d.simulateImageEnhancement(50)
	return len(d.image)
}

func (d *Day20) simulateImageEnhancement(lastStep int) {
	for ; d.step < lastStep; d.step++ {
		newImage := make(map[Coord]bool)
		newMin := Coord{d.min.X, d.min.Y}
		newMax := Coord{d.max.X, d.max.Y}
		for x := d.min.X - 1; x <= d.max.X+1; x++ {
			for y := d.min.Y - 1; y <= d.max.Y+1; y++ {
				hasPixel := d.enhancement[d.calculateEnchancement(x, y)]
				if hasPixel {
					newImage[Coord{x, y}] = hasPixel
					if x < newMin.X {
						newMin.X = x
					} else if x > newMax.X {
						newMax.X = x
					}
					if y < newMin.Y {
						newMin.Y = y
					} else if y > newMax.Y {
						newMax.Y = y
					}
				}
			}
		}
		d.min = newMin
		d.max = newMax
		d.image = newImage
		// if the enhancement mapping turns 9 dark tiles to a lit tile and vice versa, the color of infinity changes at every step
		if d.infinityIsLit {
			d.infinityIsLit = d.enhancement[511] // binary 111111111
		} else {
			d.infinityIsLit = d.enhancement[0] // binary 000000000
		}
	}
}

var bitConversion = []int{256, 128, 64, 32, 16, 8, 4, 2, 1}

func (d Day20) calculateEnchancement(x, y int) int {
	y_off := y - 1
	x_off := x - 1
	enhancement := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if d.image[Coord{i, j}] || d.inLitInfinity(i, j) {
				enhancement += bitConversion[3*(j-y_off)+i-x_off]
			}
		}
	}
	return enhancement
}

func (d Day20) inLitInfinity(x, y int) bool {
	return (x < d.min.X || x > d.max.X || y < d.min.Y || y > d.max.Y) && d.infinityIsLit
}

func (d Day20) printImage() {
	for y := d.min.Y; y <= d.max.Y; y++ {
		for x := d.min.X; x <= d.max.X; x++ {
			if d.image[Coord{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func (d *Day20) parseInput() {
	d.enhancement = make([]bool, len(d.input[0]))
	for i, r := range d.input[0] {
		d.enhancement[i] = r == '#'
	}
	d.image = make(map[Coord]bool)
	d.min = Coord{0, 0}
	x_max := 0
	y_max := 0
	for i, row := range d.input[1:] {
		for j, r := range row {
			d.image[Coord{j, i}] = r == '#'
			if i > y_max {
				y_max = i
			}
			if j > x_max {
				x_max = j
			}
		}
	}
	d.max = Coord{x_max, y_max}
}
