package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	numbers := utils.GetDay(2020, 1).ReadIntLines()

	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}

func part1(input []int) int {
	for i := range input {
		x := input[i]
		for j := i + 1; j < len(input); j++ {
			y := input[j]
			if x+y == 2020 {
				return x * y
			}
		}
	}
	return -1
}

func part2(input []int) int {
	for i := range input {
		x := input[i]
		for j := i + 1; j < len(input); j++ {
			y := input[j]
			for k := j + 1; k < len(input); k++ {
				z := input[k]
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return -1
}
