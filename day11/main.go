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
	lines := utils.SplitLines(input)
	height := len(lines)
	width := len(lines[0])
	grid := make([][]rune, height)
	for i, line := range lines {
		grid[i] = make([]rune, width)
		for j, c := range line {
			grid[i][j] = c
		}
	}

	changed := true
	for changed {
		changed = false
		newGrid := make([][]rune, height)
		for i := 0; i < height; i++ {
			newGrid[i] = make([]rune, width)
			for j := 0; j < width; j++ {
				if grid[i][j] == '.' {
					newGrid[i][j] = '.'
				} else {
					occupied := 0
					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							if x == 0 && y == 0 {
								continue
							}
							if i+x < 0 || j+y < 0 || i+x >= height || j+y >= width {
								continue
							}
							if grid[i+x][j+y] == '#' {
								occupied++
							}
						}
					}
					if occupied == 0 && grid[i][j] == 'L' {
						newGrid[i][j] = '#'
						changed = true
					} else if occupied >= 4 && grid[i][j] == '#' {
						newGrid[i][j] = 'L'
						changed = true
					} else {
						newGrid[i][j] = grid[i][j]
					}
				}
			}
		}
		grid = newGrid
		// printGrid(grid)

		// fmt.Println()
	}

	occupied := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '#' {
				occupied++
			}
		}
	}

	return occupied
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
}

func part2(input string) int {
	lines := utils.SplitLines(input)
	height := len(lines)
	width := len(lines[0])
	grid := make([][]rune, height)
	for i, line := range lines {
		grid[i] = make([]rune, width)
		for j, c := range line {
			grid[i][j] = c
		}
	}
	// printGrid(grid)

	minD := utils.Min(width, height)

	changed := true
	for changed {
		changed = false
		newGrid := make([][]rune, height)
		for i := 0; i < height; i++ {
			newGrid[i] = make([]rune, width)
			for j := 0; j < width; j++ {
				if grid[i][j] == '.' {
					newGrid[i][j] = '.'
				} else {
					occupied := 0
					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							if x == 0 && y == 0 {
								continue
							}
							for mult := 1; mult < minD; mult++ {
								nx := i + mult*x
								ny := j + mult*y
								if nx < 0 || ny < 0 || nx >= height || ny >= width {
									continue
								}
								if grid[nx][ny] == '.' {
									continue
								} else if grid[nx][ny] == '#' {
									occupied++
									break
								} else {
									break
								}
							}
						}
					}
					if occupied == 0 && grid[i][j] == 'L' {
						newGrid[i][j] = '#'
						changed = true
					} else if occupied >= 5 && grid[i][j] == '#' {
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
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '#' {
				occupied++
			}
		}
	}

	return occupied
}
