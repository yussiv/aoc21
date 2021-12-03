package mission

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct{ input []string }

func (d *Day2) SetInput(input []string) {
	d.input = input
}

func (d *Day2) Task1() int {
	directions, units := d.getDirections()
	x := 0
	z := 0
	for i, direction := range directions {
		switch direction {
		case "forward":
			x += units[i]
		case "up":
			z -= units[i]
		case "down":
			z += units[i]
		}
	}
	return x * z
}

func (d *Day2) Task2() int {
	directions, units := d.getDirections()
	x := 0
	z := 0
	aim := 0
	for i, direction := range directions {
		switch direction {
		case "forward":
			x += units[i]
			z += units[i] * aim
		case "up":
			aim -= units[i]
		case "down":
			aim += units[i]
		}
	}
	return x * z
}

func (d *Day2) getDirections() ([]string, []int) {
	directions := make([]string, len(d.input))
	units := make([]int, len(d.input))
	for i, v := range d.input {
		pair := strings.Split(v, " ")
		if len(pair) != 2 {
			fmt.Println("row has more or less than two items")
			directions[i] = ""
			units[i] = 0
			continue
		}
		directions[i] = pair[0]
		unit, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		units[i] = unit
	}
	return directions, units
}
