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
