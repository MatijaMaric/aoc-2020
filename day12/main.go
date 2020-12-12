package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 12).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	lines := utils.SplitLines(input)
	x, y := 0, 0
	dx, dy := 1, 0

	for _, line := range lines {
		instr := line[0]
		arg := utils.ToInt(line[1:])

		switch instr {
		case 'N':
			y -= arg
			break
		case 'S':
			y += arg
			break
		case 'W':
			x -= arg
			break
		case 'E':
			x += arg
			break
		case 'F':
			x += dx * arg
			y += dy * arg
			break
		case 'L':
			dx, dy = rot(dx, dy, -1, arg)
			break
		case 'R':
			dx, dy = rot(dx, dy, 1, arg)
			break
		}
	}

	return utils.Abs(x) + utils.Abs(y)
}

func part2(input string) int {
	lines := utils.SplitLines(input)
	x, y := 0, 0
	wx, wy := 10, -1

	for _, line := range lines {
		instr := line[0]
		arg := utils.ToInt(line[1:])
		switch instr {
		case 'N':
			wy -= arg
			break
		case 'S':
			wy += arg
			break
		case 'W':
			wx -= arg
			break
		case 'E':
			wx += arg
			break
		case 'F':
			x += wx * arg
			y += wy * arg
			break
		case 'L':
			wx, wy = rot(wx, wy, -1, arg)
			break
		case 'R':
			wx, wy = rot(wx, wy, 1, arg)
			break
		}
	}

	return utils.Abs(x) + utils.Abs(y)
}

func rot(dx, dy, dir, deg int) (int, int) {
	for i := 0; i < deg/90; i++ {
		if dir == -1 {
			dx, dy = dy, -dx
		} else {
			dx, dy = -dy, dx
		}
	}
	return dx, dy
}
