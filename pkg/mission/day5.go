package mission

import (
	"strconv"
	"strings"

	"github.com/yussiv/aoc21/util"
)

type Day5 struct {
	input           []string
	verticalLines   []*Line
	horizontalLines []*Line
	diagonalLines   []*Line
	ocean           [1000][1000]uint8
}

func (d *Day5) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day5) Task1() int {
	d.addLines(d.horizontalLines)
	d.addLines(d.verticalLines)
	return d.countIntersections()
}

func (d *Day5) Task2() int {
	d.addLines(d.diagonalLines)
	return d.countIntersections()
}

func (d *Day5) countIntersections() int {
	count := 0
	for _, row := range d.ocean {
		for _, value := range row {
			if value > 1 {
				count++
			}
		}
	}
	return count
}

func (d *Day5) addLines(lines []*Line) {
	for _, line := range lines {
		y_dir := line.End.Y - line.Start.Y
		if y_dir != 0 {
			y_dir = y_dir / util.Abs(y_dir)
		}
		x_dir := line.End.X - line.Start.X
		if x_dir != 0 {
			x_dir = x_dir / util.Abs(x_dir)
		}
		i_x := line.Start.X
		i_y := line.Start.Y
		for (line.End.X-i_x)*x_dir >= 0 && (line.End.Y-i_y)*y_dir >= 0 {
			d.ocean[i_y][i_x] += 1
			i_x += x_dir
			i_y += y_dir
		}
	}
}

func (d *Day5) parseInput() {
	for _, row := range d.input {
		strCoords := strings.Split(row, " -> ")
		strStart := strings.Split(strCoords[0], ",")
		x1, _ := strconv.Atoi(strStart[0])
		y1, _ := strconv.Atoi(strStart[1])
		strEnd := strings.Split(strCoords[1], ",")
		x2, _ := strconv.Atoi(strEnd[0])
		y2, _ := strconv.Atoi(strEnd[1])
		line := Line{
			Start: Coord{X: x1, Y: y1},
			End:   Coord{X: x2, Y: y2},
		}
		if line.Start.X == line.End.X {
			d.horizontalLines = append(d.horizontalLines, &line)
		} else if line.Start.Y == line.End.Y {
			d.verticalLines = append(d.verticalLines, &line)
		} else {
			d.diagonalLines = append(d.diagonalLines, &line)
		}
	}
}
