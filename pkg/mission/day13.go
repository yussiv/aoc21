package mission

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day13 struct {
	input []string
	dots  map[Coord]bool
	folds []Coord
}

func (d *Day13) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day13) Task1() int {
	sum := 0
	fold := d.folds[0]
	for k := range d.dots {
		if fold.X < 0 {
			if k.Y < fold.Y {
				sum++
			} else if _, ok := d.dots[Coord{k.X, fold.Y*2 - k.Y}]; !ok {
				sum++
			}
		} else {
			if k.X < fold.X {
				sum++
			} else if _, ok := d.dots[Coord{fold.X*2 - k.X, k.Y}]; !ok {
				sum++
			}
		}
	}
	return sum
}

func (d *Day13) Task2() int {
	for _, fold := range d.folds {
		for k := range d.dots {
			if fold.X < 0 { // fold along Y
				if k.Y > fold.Y {
					if _, ok := d.dots[Coord{k.X, fold.Y*2 - k.Y}]; !ok {
						d.dots[Coord{k.X, fold.Y*2 - k.Y}] = true
					}
					delete(d.dots, k)
				}
			} else { // fold along X
				if k.X > fold.X {
					if _, ok := d.dots[Coord{fold.X*2 - k.X, k.Y}]; !ok {
						d.dots[Coord{fold.X*2 - k.X, k.Y}] = true
					}
					delete(d.dots, k)
				}
			}
		}
	}
	max_x := 0
	max_y := 0
	for k := range d.dots {
		if max_x < k.X {
			max_x = k.X
		}
		if max_y < k.Y {
			max_y = k.Y
		}
	}
	canvas := make([][]rune, max_y+1)
	for i := range canvas {
		canvas[i] = make([]rune, max_x+1)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}
	for k := range d.dots {
		canvas[k.Y][k.X] = '#'
	}
	for _, row := range canvas {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	return 0
}

func (d *Day13) parseInput() {
	d.dots = make(map[Coord]bool)
	d.folds = make([]Coord, 0)
	for _, row := range d.input {
		pair := strings.Split(row, ",")
		if len(pair) == 2 {
			x, _ := strconv.Atoi(pair[0])
			y, _ := strconv.Atoi(pair[1])
			d.dots[Coord{x, y}] = true
		} else {
			re := regexp.MustCompile(`(x|y)=(\d+)`)
			match := re.FindAllStringSubmatch(row, -1)
			if match != nil {
				val, _ := strconv.Atoi(match[0][2])
				if match[0][1] == "x" {
					d.folds = append(d.folds, Coord{val, -1})
				} else {
					d.folds = append(d.folds, Coord{-1, val})
				}
			}
		}
	}
}
