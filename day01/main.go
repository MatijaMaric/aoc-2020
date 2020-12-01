package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	numbers := utils.GetDay(2020, 1).ReadIntLines()

	var part1, part2 int

	for i := range numbers {
		x := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			y := numbers[j]
			for k := j + 1; k < len(numbers); k++ {
				z := numbers[k]
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
