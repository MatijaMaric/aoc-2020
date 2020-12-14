package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 14).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	lines := utils.SplitLines(input)
	mem := make(map[int]int)

	var mask int
	var override int

	for _, line := range lines {
		s := strings.Split(line, " = ")
		if s[0] == "mask" {
			mask, override = parseMask(s[1])
		} else {
			addr := utils.ToInt(s[0][4 : len(s[0])-1])
			val := utils.ToInt(s[1])
			val = val&mask | override
			mem[addr] = val
		}
	}

	ans := 0
	for _, val := range mem {
		ans += val
	}

	return ans
}

func part2(input string) int {
	lines := utils.SplitLines(input)
	mem := make(map[int]int)

	var mask string

	for _, line := range lines {
		s := strings.Split(line, " = ")
		if s[0] == "mask" {
			mask = s[1]
		} else {
			x := utils.ToInt(s[0][4 : len(s[0])-1])
			val := utils.ToInt(s[1])
			for _, addr := range genAddrs(mask, x) {
				mem[addr] = val
			}
		}
	}

	ans := 0
	for _, val := range mem {
		ans += val
	}

	return ans
}

func parseMask(input string) (mask, override int) {
	mask, override = 0, 0
	for i, c := range input {
		if c == 'X' {
			mask += 1 << (len(input) - i - 1)
		}
		if c == '1' {
			override += 1 << (len(input) - i - 1)
		}
	}
	return
}

func genAddrs(input string, original int) []int {
	addrs := []int{}

	xs := strings.Count(input, "X")

	for i := 0; i < (1 << xs); i++ {
		addr := 0
		pos := xs - 1
		for j, c := range input {
			var bit int
			if c == 'X' {
				bit = (i >> pos) & 1
				pos--
			}
			if c == '1' {
				bit = 1
			}
			if c == '0' {
				bit = (original >> (len(input) - j - 1)) & 1
			}
			addr = (addr << 1) | bit
		}
		addrs = append(addrs, addr)
	}

	return addrs
}
