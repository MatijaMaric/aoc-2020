package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 11).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	grid, _, _ := parseInput(input)
	return solve(grid, 1, 4)
}

func part2(input string) int {
	grid, width, height := parseInput(input)
	minD := utils.Min(width, height)

	return solve(grid, minD, 5)
}

func solve(grid [][]rune, minD int, maxOccupy int) int {
	changed := true
	for changed {
		changed = false
		newGrid := make([][]rune, len(grid))
		for i := 0; i < len(grid); i++ {
			newGrid[i] = make([]rune, len(grid[i]))
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == '.' {
					newGrid[i][j] = '.'
				} else {
					occupied := 0
					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							if x == 0 && y == 0 {
								continue
							}
							for mult := 1; mult <= minD; mult++ {
								nx := i + mult*x
								ny := j + mult*y
								if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[i]) {
									continue
								}
								if grid[nx][ny] != '.' {
									if grid[nx][ny] == '#' {
										occupied++
									}
									break
								}
							}
						}
					}
					if occupied == 0 && grid[i][j] == 'L' {
						newGrid[i][j] = '#'
						changed = true
					} else if occupied >= maxOccupy && grid[i][j] == '#' {
						newGrid[i][j] = 'L'
						changed = true
					} else {
						newGrid[i][j] = grid[i][j]
					}
				}
			}
		}
		grid = newGrid
	}

	occupied := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				occupied++
			}
		}
	}

	return occupied
}

func parseInput(input string) (grid [][]rune, width, height int) {
	lines := utils.SplitLines(input)
	height = len(lines)
	width = len(lines[0])
	grid = make([][]rune, height)
	for i, line := range lines {
		grid[i] = make([]rune, width)
		for j, c := range line {
			grid[i][j] = c
		}
	}
	return
}
