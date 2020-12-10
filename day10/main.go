package main

import (
	"fmt"
	"sort"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 10).ReadIntLines()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []int) int {
	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	ones := 0
	threes := 0

	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		if diff == 1 {
			ones++
		}
		if diff == 3 {
			threes++
		}
	}

	return ones * threes
}

func part2(input []int) int {
	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	combs := make([]int, len(input))
	combs[len(combs)-1] = 1

	for i := len(input) - 2; i >= 0; i-- {
		for j := i + 1; j < len(input) && input[j]-input[i] <= 3; j++ {
			combs[i] += combs[j]
		}
	}

	return combs[0]
}
