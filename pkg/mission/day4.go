package mission

import (
	"strconv"
	"strings"
)

type bingoBoard struct {
	horizontalSums    [5]int
	verticalSums      [5]int
	numberCoordinates map[int]Coord
	hasBingo          bool
}

type Day4 struct {
	input                  []string
	boards                 []*bingoBoard
	drawnNumbers           []int
	boardsContainingNumber map[int][]*bingoBoard
	firstScore             int
	lastScore              int
}

func (d *Day4) SetInput(input []string) {
	d.input = input
	d.parseInput()
	d.calculateFirstAndLastBingoBoardScores()
}

func (d *Day4) Task1() int {
	return d.firstScore
}

func (d *Day4) Task2() int {
	return d.lastScore
}

func (d *Day4) calculateFirstAndLastBingoBoardScores() {
	var firstBoard *bingoBoard
	var firstNumber int
	var lastBoard *bingoBoard
	var lastNumber int

	for _, number := range d.drawnNumbers {
		boardsContainingNumber := d.boardsContainingNumber[number]
		for _, board := range boardsContainingNumber {
			if board.hasBingo {
				continue
			}
			x := board.numberCoordinates[number].X
			y := board.numberCoordinates[number].Y
			board.horizontalSums[x] -= number
			board.verticalSums[y] -= number
			if board.horizontalSums[x] == 0 || board.verticalSums[y] == 0 {
				// bingo!
				if firstBoard == nil {
					firstBoard = board
					firstNumber = number
				}
				board.hasBingo = true
				lastBoard = board
				lastNumber = number
			}
		}
	}
	firstBoardSum := d.calculateBoardSum(firstBoard)
	lastBoardSum := d.calculateBoardSum(lastBoard)

	d.firstScore = firstBoardSum * firstNumber
	d.lastScore = lastBoardSum * lastNumber
}

func (Day4) calculateBoardSum(board *bingoBoard) int {
	sum := 0
	for _, value := range board.horizontalSums {
		sum += value
	}
	return sum
}

func (d *Day4) parseInput() {
	d.drawnNumbers = d.parseDrawnNumbers(d.input[0])
	d.boardsContainingNumber = make(map[int][]*bingoBoard, 100)

	// input loader strips empty rows, so bingo board takes 5 lines instead of 6
	d.boards = make([]*bingoBoard, len(d.input)/5)
	for i := range d.boards {
		d.boards[i] = d.createBingoBoard(i)
	}
}

func (d *Day4) createBingoBoard(index int) *bingoBoard {
	inputOffset := 1 + index*5
	board := new(bingoBoard)
	board.numberCoordinates = make(map[int]Coord, 25)
	for i := inputOffset; i < inputOffset+5; i++ {
		x := i - inputOffset
		row := strings.Fields(d.input[i])
		for y, num := range row {
			intVal, _ := strconv.Atoi(num)
			board.horizontalSums[x] += intVal
			board.verticalSums[y] += intVal
			board.numberCoordinates[intVal] = Coord{X: x, Y: y}
			d.boardsContainingNumber[intVal] = append(d.boardsContainingNumber[intVal], board)
		}
	}
	return board
}

func (Day4) parseDrawnNumbers(str string) []int {
	numberStrings := strings.Split(strings.TrimSpace(str), ",")
	numbers := make([]int, len(numberStrings))
	for i, numberStr := range numberStrings {
		value, _ := strconv.Atoi(numberStr)
		numbers[i] = value
	}
	return numbers
}
