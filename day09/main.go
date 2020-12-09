package main

import (
	"fmt"
	"sort"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 9).ReadIntLines()

	first := part1(input, 25)
	fmt.Println(first)
	fmt.Println(part2(input, first))
}

func part1(input []int, preamble int) int {
OUTER:
	for i := preamble; i < len(input); i++ {
		for j := i - preamble; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if input[i] == input[j]+input[k] {
					continue OUTER
				}
			}
		}
		return input[i]
	}
	return -1
}

func part2(input []int, target int) int {
	first := 0
	last := 1
	sum := input[0] + input[1]
	for sum != target {
		if sum < target {
			last++
			sum += input[last]
		}
		if sum > target {
			sum -= input[first]
			first++
		}
	}
	nums := input[first : last+1]
	sort.Ints(nums)
	return nums[0] + nums[len(nums)-1]
}
