package mission

import (
	"container/heap"
	"strconv"
	"strings"
)

type xyRisk struct {
	Pos  Coord
	Risk int
}

type riskHeap []xyRisk

func (h riskHeap) Len() int           { return len(h) }
func (h riskHeap) Less(i, j int) bool { return h[i].Risk < h[j].Risk }
func (h riskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *riskHeap) Push(x interface{}) {
	*h = append(*h, x.(xyRisk))
}

func (h *riskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Day15 struct {
	input  []string
	risk   [][]int
	width  int
	height int
}

func (d *Day15) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day15) Task1() int {
	return d.bendItLikeDijkstra(d.width, d.height)
}

func (d *Day15) Task2() int {
	return d.bendItLikeDijkstra(d.width*5, d.height*5)
}

func (d Day15) bendItLikeDijkstra(width, height int) int {
	h := &riskHeap{}
	heap.Init(h)
	for _, n := range d.getNeighbors(Coord{0, 0}, width, height) {
		heap.Push(h, n)
	}
	visited := make(map[Coord]bool)
	end := Coord{width - 1, height - 1}

	for h.Len() > 0 {
		xyr := heap.Pop(h).(xyRisk)
		pos := xyr.Pos
		if visited[pos] {
			continue
		}
		visited[pos] = true
		if pos == end {
			return xyr.Risk
		} else {
			for _, n := range d.getNeighbors(pos, width, height) {
				if !visited[n.Pos] {
					n.Risk += xyr.Risk
					heap.Push(h, n)
				}
			}
		}
	}
	return -1
}

func (d Day15) getNeighbors(pos Coord, width, height int) []xyRisk {
	x, y := pos.X, pos.Y
	neighbors := make([]xyRisk, 0, 4)
	for _, p := range []Coord{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
		if p.X >= 0 && p.Y >= 0 && p.X < width && p.Y < height {
			neighbors = append(neighbors, xyRisk{p, d.getRisk(p)})
		}
	}
	return neighbors
}

func (d Day15) getRisk(pos Coord) int {
	xOffset := pos.X / d.width
	yOffset := pos.Y / d.height
	x := pos.X % d.width
	y := pos.Y % d.height
	return (d.risk[y][x]+xOffset+yOffset-1)%9 + 1
}

func (d *Day15) parseInput() {
	d.risk = make([][]int, len(d.input))
	for y, row := range d.input {
		riskStrs := strings.Split(row, "")
		d.risk[y] = make([]int, len(riskStrs))
		for x, riskStr := range riskStrs {
			risk, _ := strconv.Atoi(riskStr)
			d.risk[y][x] = risk
		}
	}
	d.height = len(d.risk)
	d.width = len(d.risk[0])
}
