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
	aoc.buffer = GetInput(year, day)
	return aoc
}

// ToString reads buffer as string
func (aoc *AdventOfCode) ToString() string {
	return strings.TrimSpace(string(aoc.buffer))
}

// SplitLines splits string into lines
func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}

// ReadLines splits buffer into lines of string
func (aoc *AdventOfCode) ReadLines() []string {
	return SplitLines(aoc.ToString())
}

// ReadBoolGrid reads grid marked by # and . as boolean
func (aoc *AdventOfCode) ReadBoolGrid() [][]bool {
	input := aoc.ReadLines()
	grid := make([][]bool, len(input))
	for i, line := range input {
		grid[i] = make([]bool, len(line))
		for j, c := range line {
			grid[i][j] = c == '#'
		}
	}
	return grid
}

// ParseIntLines parses lines of string to int
func ParseIntLines(input string) []int {
	lines := strings.Fields(input)
	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i] = ToInt(line)
	}
	return numbers
}

// ReadIntLines converts lines of string to int
func (aoc *AdventOfCode) ReadIntLines() []int {
	return ParseIntLines(aoc.ToString())
}

// ParseIntList converts comma seperated list to int array
func ParseIntList(input string) []int {
	isComma := func(c rune) bool {
		return c == ','
	}
	fields := strings.FieldsFunc(input, isComma)
	numbers := make([]int, len(fields))
	for i, num := range fields {
		numbers[i] = ToInt(num)
	}
	return numbers
}

// ReadIntList converts comma seperated list in the buffer to int array
func (aoc *AdventOfCode) ReadIntList() []int {
	return ParseIntList(aoc.ToString())
}
