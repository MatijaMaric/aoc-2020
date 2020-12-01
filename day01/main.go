package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	numbers := utils.GetDay(2020, 1).ReadIntLines()

	var part1, part2 int

	for i, x := range numbers {
		for j, y := range numbers {
			if i == j {
				continue
			}
			for k, z := range numbers {
				if j == k {
					continue
				}
				if x+y+z == 2020 {
					part2 = x * y * z
				}
			}
			if x+y == 2020 {
				part1 = x * y
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
