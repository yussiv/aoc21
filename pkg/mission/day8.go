package mission

import (
	"log"
	"math/bits"
	"strings"

	"github.com/yussiv/aoc21/util"
)

type Day8 struct {
	input    []string
	patterns [][]uint8
	statuses [][]uint8
}

func (d *Day8) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day8) Task1() int {
	sum := 0
	for _, status := range d.statuses {
		for _, digit := range status {
			n := rank8(digit)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				sum++
			}
		}
	}
	return sum
}

func (d *Day8) Task2() int {
	sum := 0
	digits := make([]uint8, 10)
	for i, status := range d.statuses {
		fivers := make([]uint8, 0, 3)
		sixers := make([]uint8, 0, 3)
		pattern := d.patterns[i]
		for _, digit := range pattern {
			n := rank8(digit)
			switch n {
			case 2:
				digits[1] = digit
			case 3:
				digits[7] = digit
			case 4:
				digits[4] = digit
			case 5:
				fivers = append(fivers, digit)
			case 6:
				sixers = append(sixers, digit)
			case 7:
				digits[8] = digit
			}
		}
		for _, sixer := range sixers {
			if rank8(sixer|digits[1]) == 7 {
				digits[6] = sixer
			} else if rank8(sixer|digits[4]) == 7 {
				digits[0] = sixer
			} else {
				digits[9] = sixer
			}
		}
		for _, fiver := range fivers {
			if rank8(fiver|digits[1]) == 5 {
				digits[3] = fiver
			} else if rank8(fiver|digits[9]) == 7 {
				digits[2] = fiver
			} else {
				digits[5] = fiver
			}
		}

		digitMap := make(map[uint8]int, 10)
		for i, digit := range digits {
			digitMap[digit] = i
		}
		result := 0
		for i, code := range status {
			result += util.IntPow(10, 3-i) * digitMap[code]
		}
		sum += result
	}
	return sum
}

func rank8(x uint8) int {
	return bits.OnesCount8(x)
}

func (d *Day8) parseInput() {
	d.patterns = make([][]uint8, len(d.input))
	d.statuses = make([][]uint8, len(d.input))
	for i, line := range d.input {
		splitLine := strings.Split(line, " | ")
		fields := strings.Fields(splitLine[0])
		d.patterns[i] = d.encode(fields)
		fields = strings.Fields(splitLine[1])
		d.statuses[i] = d.encode(fields)
		if len(d.patterns[i]) != 10 || len(d.statuses[i]) != 4 {
			log.Fatalln("input did not contain the correct amount of fields")
		}
	}
}

func (Day8) encode(fields []string) []uint8 {
	result := make([]uint8, len(fields))
	for i, field := range fields {
		value := uint8(0)
		for _, c := range field {
			switch c {
			case 'a':
				value += 1
			case 'b':
				value += 2
			case 'c':
				value += 4
			case 'd':
				value += 8
			case 'e':
				value += 16
			case 'f':
				value += 32
			case 'g':
				value += 64
			}
		}
		result[i] = value
	}
	return result
}
