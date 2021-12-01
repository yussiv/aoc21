package util

import (
	"log"
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	str := strings.Trim(string(data), "\n\r ")
	lines := strings.Split(str, "\n")
	return lines
}
