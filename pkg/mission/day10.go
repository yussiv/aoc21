package mission

import (
	"sort"

	"github.com/yussiv/aoc21/util"
)

type Day10 struct {
	input            []string
	errorScore       int
	completionScores []int
}

func (d *Day10) SetInput(input []string) {
	d.input = input
	d.calculateScores()
}

func (d *Day10) Task1() int {
	return d.errorScore
}

func (d *Day10) Task2() int {
	return d.completionScores[len(d.completionScores)/2]
}

func (d *Day10) calculateScores() {
	errorPoints := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	completionPoints := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	openingRune := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	closingRune := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	stack := util.NewRuneStack()
	d.completionScores = make([]int, 0, len(d.input))
	d.errorScore = 0

	for _, line := range d.input {
		stack.Reset()
		isNotCorrupted := true
		completionScore := 0
	Loop:
		for _, char := range line {
			switch char {
			case '[', '<', '(', '{':
				stack.Push(char)
			default:
				if stack.Peek() == openingRune[char] {
					stack.Pop()
				} else {
					d.errorScore += errorPoints[char]
					isNotCorrupted = false
					break Loop
				}
			}
		}
		if isNotCorrupted {
			for !stack.IsEmpty() {
				char := stack.Pop()
				completionScore = completionScore*5 + completionPoints[closingRune[char]]
			}
			d.completionScores = append(d.completionScores, completionScore)
		}
	}
	sort.Ints(sort.IntSlice(d.completionScores))
}
