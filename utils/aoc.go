package utils

import (
	"strings"
)

// AdventOfCode helper for loading input data
type AdventOfCode struct {
	buffer []byte
}

// GetDay fetches input data for problem
func GetDay(year int, day int) *AdventOfCode {
	aoc := new(AdventOfCode)
	aoc.buffer = Fetch(year, day)
	return aoc
}

// ToString reads buffer as string
func (aoc *AdventOfCode) ToString() string {
	return strings.TrimSpace(string(aoc.buffer))
}

// ReadLines splits buffer into lines of string
func (aoc *AdventOfCode) ReadLines() []string {
	return strings.Split(aoc.ToString(), "\n")
}

// ReadIntLines converts lines of string to int
func (aoc *AdventOfCode) ReadIntLines() []int {
	lines := strings.Fields(aoc.ToString())
	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i] = ToInt(line)
	}
	return numbers
}

// ReadIntList converts comma seperated list in the buffer to int array
func (aoc *AdventOfCode) ReadIntList() []int {
	isComma := func(c rune) bool {
		return c == ','
	}
	fields := strings.FieldsFunc(aoc.ToString(), isComma)
	numbers := make([]int, len(fields))
	for i, num := range fields {
		numbers[i] = ToInt(num)
	}
	return numbers
}
