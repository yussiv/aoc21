package mission

import "github.com/yussiv/aoc21/util"

type Day3 struct{ input []string }

func (d *Day3) SetInput(input []string) {
	d.input = input
}

func (d *Day3) Task1() int {
	gamma, epsilon := d.calculateGammaAndEpsilon()
	return gamma * epsilon
}

func (d *Day3) Task2() int {
	o2 := d.getRating(true)
	co2 := d.getRating(false)
	return o2 * co2
}

func (d *Day3) getRating(isO2 bool) int {
	codes := util.BinaryStringsToIntegers(d.input)
	nCodes := uint16(len(codes))
	bitOffset := len(d.input[0]) - 1
	nOnes := uint16(0)
	for _, code := range codes {
		nOnes += uint16(code >> bitOffset & 1)
	}
	targetBit := uint8(0)
	if nOnes >= nCodes-nOnes && isO2 || nOnes <= nCodes-nOnes && !isO2 {
		targetBit = 1
	}
	nOnes = 0
	var j uint16 = 0
	var i uint16 = 0
	for {
		if nCodes == 1 {
			break
		}
		if bit := uint8(codes[i] >> bitOffset & 1); bit == targetBit {
			codes[j] = codes[i]
			if bitOffset > 0 {
				nextBit := uint16(codes[i] >> (bitOffset - 1) & 1)
				nOnes += nextBit
			}
			j++
		}
		if i == nCodes-1 {
			if nOnes >= j-nOnes && isO2 || nOnes < j-nOnes && !isO2 {
				targetBit = 1
			} else {
				targetBit = 0
			}
			nOnes = 0
			nCodes = j
			i = 0
			j = 0
			bitOffset--
		} else {
			i++
		}
	}
	return int(codes[0])
}

func (d *Day3) calculateGammaAndEpsilon() (gamma int, epsilon int) {
	nDigits := len(d.input[0])
	counts := make([]int, nDigits)
	for _, str := range d.input {
		for i, c := range str {
			if c == '1' {
				counts[i] += 1
			}
		}
	}
	half := len(d.input) >> 1
	for i, n := range counts {
		if n >= half {
			gamma += (1 << (nDigits - 1 - i))
		} else {
			epsilon += (1 << (nDigits - 1 - i))
		}
	}
	return
}
