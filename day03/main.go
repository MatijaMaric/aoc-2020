package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 3).ReadBoolGrid()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input [][]bool) int {
	return count(input, 1, 3)
}

func part2(input [][]bool) int {
	return count(input, 1, 1) * count(input, 1, 3) * count(input, 1, 5) * count(input, 1, 7) * count(input, 2, 1)
}

func count(grid [][]bool, dy, dx int) int {
	x := 0
	ans := 0
	for y := 0; y < len(grid); y += dy {
		if grid[y][x] {
			ans++
		}
		x = (x + dx) % len(grid[y])
	}
	return ans
}
