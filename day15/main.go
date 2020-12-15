package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 15).ReadIntList()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []int) int {
	return solve(input, 2020)
}

func part2(input []int) int {
	return solve(input, 30000000)
}

func solve(input []int, iter int) int {
	type number struct {
		a int
		b int
	}

	numbers := make(map[int]number)

	for i, num := range input {
		if _, ok := numbers[num]; !ok {
			numbers[num] = number{i, -1}
		} else {
			numbers[num] = number{i, numbers[num].a}
		}
	}

	last := input[len(input)-1]

	for i := len(input); i < iter; i++ {
		if numbers[last].b == -1 {
			last = 0
		} else {
			last = numbers[last].a - numbers[last].b
		}
		if _, ok := numbers[last]; !ok {
			numbers[last] = number{i, -1}
		} else {
			numbers[last] = number{i, numbers[last].a}
		}
	}

	return last
}
