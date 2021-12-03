package util

import (
	"log"
	"strconv"
)

func StringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, row := range strs {
		value, err := strconv.Atoi(row)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = value
	}
	return ints
}

func BinaryStringsToIntegers(strs []string) []uint16 {
	result := make([]uint16, len(strs))
	n := len(strs[0])
	for i, str := range strs {
		for j, c := range str {
			if c == '1' {
				result[i] += (1 << (n - 1 - j))
			}
		}
	}
	return result
}
