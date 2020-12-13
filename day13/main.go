package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 13).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	time, buses := parseInput(input)

	for i := time; ; i++ {
		for _, bus := range buses {
			if bus == -1 {
				continue
			}
			if i%bus == 0 {
				return (i - time) * bus
			}
		}
	}
}

func part2(input string) int {
	_, buses := parseInput(input)

	time := 0
	step := 1

	for i, bus := range buses {
		if bus == -1 {
			continue
		}
		for j := time; ; j += step {
			if (j+i)%buses[i] == 0 {
				time = j
				break
			}
		}
		step *= buses[i]
	}

	return time
}

func parseInput(input string) (int, []int) {
	lines := utils.SplitLines(input)
	time := utils.ToInt(lines[0])
	var buses []int

	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			buses = append(buses, -1)
		} else {
			buses = append(buses, utils.ToInt(bus))
		}
	}
	return time, buses
}
