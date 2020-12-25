package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 25).ReadIntLines()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []int) int {
	encA, encB := 1, 1
	crack := 1
	for {
		if crack == input[0] {
			return encB
		}
		if crack == input[1] {
			return encA
		}
		crack = (crack * 7) % 20201227
		encA = (encA * input[0]) % 20201227
		encB = (encB * input[1]) % 20201227
	}
}

func part2(input []int) int {
	return -1
}
