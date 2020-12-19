package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 18).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	ans := 0

	for _, line := range utils.SplitLines(input) {
		ans += solve1(line)
	}

	return ans
}

func part2(input string) int {
	ans := 0

	for _, line := range utils.SplitLines(input) {
		ans += solve2(line)
	}

	return ans
}

func compute2(input string) int {

	re := regexp.MustCompile(`(\d+\s\+\s\d+)`)
	expr := input
	for re.MatchString(expr) {
		loc := re.FindStringIndex(expr)
		if loc == nil {
			break
		}
		slice := expr[loc[0]:loc[1]]
		expr = fmt.Sprintf("%s%d%s", expr[0:loc[0]], compute(slice), expr[loc[1]:len(expr)])
	}
	return compute(expr)
}

func compute(input string) int {
	parts := strings.Split(input, " ")
	ans := utils.ToInt(parts[0])

	for i := 1; i < len(parts); i += 2 {
		num := utils.ToInt(parts[i+1])
		if parts[i] == "+" {
			ans += num
		}
		if parts[i] == "*" {
			ans *= num
		}
	}
	return ans
}

func solve2(input string) int {
	return solve(input, true)
}

func solve1(input string) int {
	return solve(input, false)
}

func solve(input string, part2 bool) int {
	re := regexp.MustCompile(`(\([^()]+\))`)

	expr := input

	for {
		loc := re.FindStringIndex(expr)
		if loc == nil {
			if part2 {
				return compute2(expr)
			}
			return compute(expr)
		}
		slice := expr[loc[0]+1 : loc[1]-1]
		var newsub int
		if part2 {
			newsub = compute2(slice)
		} else {
			newsub = compute(slice)
		}
		expr = fmt.Sprintf("%s%d%s", expr[0:loc[0]], newsub, expr[loc[1]:len(expr)])
	}
}
