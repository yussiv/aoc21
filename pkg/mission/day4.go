package mission

import (
	"fmt"
	"strconv"
	"strings"
)

type bingoBoard struct {
	horizontalSums [5]int
	verticalSums   [5]int
	numberPosition map[int][2]int
	hasBingo       bool
}

type Day4 struct {
	input          []string
	boards         []bingoBoard
	drawnNumbers   []int
	lastDrawnIndex int
	numberMap      map[int][]int
}

func (d *Day4) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day4) Task1() int {
	for i, num := range d.drawnNumbers {
		d.lastDrawnIndex = i
		boardsWithNumber := d.numberMap[num]
		for _, boardIndex := range boardsWithNumber {
			board := &d.boards[boardIndex]
			position := board.numberPosition[num]
			board.horizontalSums[position[0]] -= num
			board.verticalSums[position[1]] -= num
			if board.horizontalSums[position[0]] == 0 || board.verticalSums[position[1]] == 0 {
				fmt.Println("BINGO!")
				board.hasBingo = true
				leftoverSum := 0
				for _, value := range board.horizontalSums {
					leftoverSum += value
				}
				return num * leftoverSum
			}
		}
	}
	return 0
}

func (d *Day4) Task2() int {
	var lastBoard *bingoBoard
	var lastNumber int
	// continue where task 1 left of, so we don't need to build the data structure again
	for i := d.lastDrawnIndex + 1; i < len(d.drawnNumbers); i++ {
		num := d.drawnNumbers[i]
		boardsWithNumber := d.numberMap[num]
		for _, boardIndex := range boardsWithNumber {
			board := &d.boards[boardIndex]
			if board.hasBingo {
				continue
			}
			position := board.numberPosition[num]
			board.horizontalSums[position[0]] -= num
			board.verticalSums[position[1]] -= num
			if board.horizontalSums[position[0]] == 0 || board.verticalSums[position[1]] == 0 {
				// bingo!
				board.hasBingo = true
				lastBoard = board
				lastNumber = num
			}
		}
	}
	leftoverSum := 0
	for _, value := range lastBoard.horizontalSums {
		leftoverSum += value
	}
	return lastNumber * leftoverSum
}

func (d *Day4) parseInput() {
	d.drawnNumbers = d.parseDrawnNumbers(d.input[0])
	// keep track of which boards contain which bingo number from range 0-99
	d.numberMap = make(map[int][]int, 100)

	// input loader strips empty rows, so bingo board takes 5 lines instead of 6
	d.boards = make([]bingoBoard, len(d.input)/5)
	for i := range d.boards {
		d.boards[i] = *d.createBingoBoard(i)
	}
}

func (d *Day4) createBingoBoard(index int) *bingoBoard {
	offset := 1 + index*5
	board := new(bingoBoard)
	board.numberPosition = make(map[int][2]int, 25)
	for i := offset; i < offset+5; i++ {
		x := i - offset
		row := strings.Fields(d.input[i])
		for y, num := range row {
			intVal, _ := strconv.Atoi(num)
			board.horizontalSums[x] += intVal
			board.verticalSums[y] += intVal
			board.numberPosition[intVal] = [2]int{x, y}
			// store information that this board contains the current value
			d.numberMap[intVal] = append(d.numberMap[intVal], index)
		}
	}
	return board
}

func (Day4) parseDrawnNumbers(str string) []int {
	splitStr := strings.Split(strings.TrimSpace(str), ",")
	nums := make([]int, len(splitStr))
	for i, numStr := range splitStr {
		value, _ := strconv.Atoi(numStr)
		nums[i] = value
	}
	return nums
}
