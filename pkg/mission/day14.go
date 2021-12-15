package mission

import (
	"strings"
)

type monomer rune
type dimer [2]monomer

type Day14 struct {
	input    []string
	template []monomer
	rules    map[dimer]monomer
	memo     map[dimer]map[int]map[monomer]int
}

func (d *Day14) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day14) Task1() int {
	return d.countPolymerComponents(10)
}

func (d *Day14) Task2() int {
	return d.countPolymerComponents(40)
}

func (d *Day14) countPolymerComponents(depth int) int {
	counts := make(map[monomer]int)
	for _, r := range d.template {
		counts[r]++
	}
	for j := 0; j < len(d.template)-1; j++ {
		pair := dimer{d.template[j], d.template[j+1]}
		for k, v := range d.countTwomerInsertions(pair, depth) {
			counts[k] += v
		}
	}
	max := 0
	min := int(^uint(0) >> 1)
	for _, v := range counts {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return max - min
}

func (d *Day14) countTwomerInsertions(pair dimer, depth int) map[monomer]int {
	insert := d.rules[pair]
	if depth == 0 {
		return map[monomer]int{}
	}
	if depth == 1 {
		return map[monomer]int{insert: 1}
	}
	pair1 := dimer{pair[0], insert}
	pair2 := dimer{insert, pair[1]}

	nextDepth := depth - 1
	counts := make(map[monomer]int)
	if _, ok := d.memo[pair1][nextDepth]; !ok {
		d.memo[pair1][nextDepth] = d.countTwomerInsertions(pair1, nextDepth)
	}
	if _, ok := d.memo[pair2][nextDepth]; !ok {
		d.memo[pair2][nextDepth] = d.countTwomerInsertions(pair2, nextDepth)
	}
	for k, v := range d.memo[pair1][nextDepth] {
		counts[k] = v
	}
	for k, v := range d.memo[pair2][nextDepth] {
		counts[k] += v
	}
	counts[insert]++
	return counts
}

func (d *Day14) parseInput() {
	d.template = []monomer(d.input[0])
	d.rules = make(map[dimer]monomer)
	d.memo = make(map[dimer]map[int]map[monomer]int)
	for _, row := range d.input[1:] {
		pair := strings.Split(row, " -> ")
		monomers := []monomer(pair[0])
		id := dimer{monomers[0], monomers[1]}
		d.rules[id] = []monomer(pair[1])[0]
		d.memo[id] = make(map[int]map[monomer]int, 0)
	}
}
