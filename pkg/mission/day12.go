package mission

import (
	"strings"
	"unicode"

	"github.com/yussiv/aoc21/util"
)

type Day12 struct {
	input      []string
	neighbors  map[string][]string
	routeCount map[string]int
}

func (d *Day12) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day12) Task1() int {
	d.routeCount = make(map[string]int)
	v := make(map[string]bool)
	d.traverseCave("start", v, false)
	return d.routeCount["end"]
}

func (d *Day12) Task2() int {
	d.routeCount = make(map[string]int)
	v := make(map[string]bool)
	d.traverseCave("start", v, true)
	return d.routeCount["end"]
}

func (d *Day12) traverseCave(cave string, visited map[string]bool, hasExtraTime bool) {
	nextVisited := make(map[string]bool)
	for k, v := range visited {
		nextVisited[k] = v
	}
	if unicode.IsLower(util.RuneAt(cave, 0)) {
		nextVisited[cave] = true
	}
	for _, neighbor := range d.neighbors[cave] {
		if neighbor != "start" {
			d.routeCount[neighbor]++
			if neighbor != "end" {
				if !visited[neighbor] {
					d.traverseCave(neighbor, nextVisited, hasExtraTime)
				} else if hasExtraTime {
					d.traverseCave(neighbor, nextVisited, false)
				}
			}
		}
	}
}

func (d *Day12) parseInput() {
	d.neighbors = make(map[string][]string)
	for _, row := range d.input {
		pair := strings.Split(row, "-")
		if _, ok := d.neighbors[pair[0]]; !ok {
			d.neighbors[pair[0]] = make([]string, 0)
		}
		if _, ok := d.neighbors[pair[1]]; !ok {
			d.neighbors[pair[1]] = make([]string, 0)
		}
		d.neighbors[pair[0]] = append(d.neighbors[pair[0]], pair[1])
		d.neighbors[pair[1]] = append(d.neighbors[pair[1]], pair[0])
	}
}
