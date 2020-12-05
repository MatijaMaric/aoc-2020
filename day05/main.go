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
		x, y := decode(pass)
		max = utils.Max(max, x*8+y)
	}

	return max
}

func part2(input []string) int {
	var seats [128][8]bool

	for _, pass := range input {
		x, y := decode(pass)
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

func decode(input string) (x, y int) {
	rows := input[0:7]
	cols := input[7:10]

	x = bsp(rows, 'F')
	y = bsp(cols, 'L')

	return
}

func bsp(input string, char rune) int {
	l, r := 0, (1<<len(input))-1
	for _, a := range input {
		m := (l + r) / 2
		if a == char {
			r = m
		} else {
			l = m + 1
		}
		if l == r {
			return l
		}
	}
	return -1
}
