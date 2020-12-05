package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 5).ReadLines()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	max := 0

	for _, pass := range input {
		x, y := bsp(pass)
		max = utils.Max(max, x*8+y)
	}

	return max
}

func part2(input []string) int {
	var seats [128][8]bool

	for _, pass := range input {
		x, y := bsp(pass)
		seats[x][y] = true
	}

	for x := 1; x < 127; x++ {
		for y := 0; y < 8; y++ {
			if !seats[x][y] {
				return x*8 + y
			}
		}
	}
	return -1
}

func bsp(input string) (x, y int) {
	rows := input[0:7]
	cols := input[7:10]

	l, r := 0, 127
	var m int
	for _, a := range rows {
		m = (l + r) / 2
		if a == 'F' {
			r = m
		} else {
			l = m + 1
		}
		if l == r {
			x = l
		}
	}

	l, r = 0, 7
	for _, a := range cols {
		m = (l + r) / 2
		if a == 'L' {
			r = m
		} else {
			l = m + 1
		}
		if l == r {
			y = l
		}
	}

	return
}
