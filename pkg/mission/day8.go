package mission

import (
	"log"
	"math"
	"math/bits"
	"strings"
)

type Day8 struct {
	input    []string
	patterns [][]uint16
	statuses [][]uint16
}

func (d *Day8) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day8) Task1() int {
	sum := 0
	for _, status := range d.statuses {
		for _, digit := range status {
			n := rank(digit)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				sum++
			}
		}
	}
	return sum
}

func (d *Day8) Task2() int {
	sum := 0
	deleted := ^uint16(0)
	for i, status := range d.statuses {
		digits := make([]uint16, 10)
		fivers := make([]uint16, 0, 3)
		sixers := make([]uint16, 0, 3)
		pattern := d.patterns[i]
		for _, digit := range pattern {
			n := rank(digit)
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

		/*
			Segment mapping:

				 0000
				1    2
				1    2
				 3333
				4    5
				4    5
				 6666
		*/
		segments := make([]uint16, 7)
		segments[0] = digits[7] & ^digits[1]

		mask47 := digits[4] | digits[7]
		for i, sixer := range sixers {
			if rank(sixer & ^mask47) == 1 {
				segments[6] = sixer & ^mask47
				digits[9] = sixer
				sixers[i] = deleted
				break
			}
		}
		d.assertNotZero(digits, 9)

		mask106 := digits[1] | segments[0] | segments[6]
		for i, fiver := range fivers {
			if rank(fiver & ^mask106) == 1 {
				segments[3] = fiver & ^mask106
				digits[3] = fiver
				fivers[i] = deleted
				break
			}
		}
		d.assertNotZero(digits, 3)

		for i, sixer := range sixers {
			if rank(sixer&segments[3]) == 0 {
				digits[0] = sixer
				sixers[i] = deleted
				break
			}
		}
		d.assertNotZero(digits, 0)

		for i, fiver := range fivers {
			if rank(fiver&^digits[9]) == 0 {
				digits[5] = fiver
				fivers[i] = deleted
				break
			}
		}
		d.assertNotZero(digits, 5)

		for _, sixer := range sixers {
			if sixer != deleted {
				digits[6] = sixer
				break
			}
		}
		d.assertNotZero(digits, 6)

		for _, fiver := range fivers {
			if fiver != deleted {
				digits[2] = fiver
				break
			}
		}
		d.assertNotZero(digits, 2)

		digitMap := make(map[uint16]int, 10)
		for i, digit := range digits {
			digitMap[digit] = i
		}
		result := 0
		for i, code := range status {
			result += int(math.Pow(10, float64(3-i))) * digitMap[code]
		}
		sum += result
	}
	return sum
}

func rank(x uint16) int {
	return bits.OnesCount16(x)
}

func (Day8) assertNotZero(digits []uint16, num int) {
	if digits[num] == 0 {
		log.Fatalf("didn't find %d\n", num)
	}
}

func (d *Day8) parseInput() {
	d.patterns = make([][]uint16, len(d.input))
	d.statuses = make([][]uint16, len(d.input))
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

func (Day8) encode(fields []string) []uint16 {
	result := make([]uint16, len(fields))
	for i, field := range fields {
		value := uint16(0)
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
