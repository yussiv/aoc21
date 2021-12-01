package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetInput(path string, dayNumber int) []string {
	return readLines(resolveInputFile(path, dayNumber))
}

func readLines(input *os.File) []string {
	lines := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n\r ")
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func resolveFallbackInputFile(fallbackDayNumber int) *os.File {
	path := fmt.Sprintf("./input/day%d", fallbackDayNumber)
	fmt.Printf("Using input file %s\n", path)
	fallbackPath, pathErr := filepath.Abs(path)
	if pathErr != nil {
		panic(pathErr)
	}
	file, openErr := os.Open(fallbackPath)
	if openErr != nil {
		panic(openErr)
	}
	return file
}

func resolveInputFile(path string, fallbackDayNumber int) *os.File {
	if path == "" {
		return resolveFallbackInputFile(fallbackDayNumber)
	}

	absPath, pathErr := filepath.Abs(path)
	if pathErr != nil {
		panic(pathErr)
	}

	file, openErr := os.Open(absPath)
	if openErr != nil {
		fmt.Printf("Could not open %s\n", absPath)
		return resolveFallbackInputFile(fallbackDayNumber)
	}
	return file
}
