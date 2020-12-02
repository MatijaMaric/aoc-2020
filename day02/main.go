package main

import (
	"fmt"
	"regexp"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 2).ReadLines()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	ans := 0
	for _, line := range input {
		re := regexp.MustCompile(`(\d+)\-(\d+) (\w): (\w+)`)
		matches := re.FindStringSubmatch(line)
		min := utils.ToInt(matches[1])
		max := utils.ToInt(matches[2])
		letter := matches[3][0]
		password := matches[4]

		count := 0
		for _, x := range password {
			if rune(x) == rune(letter) {
				count++
			}
		}
		if count >= min && count <= max {
			ans++
		}
	}

	return ans
}

func part2(input []string) int {
	ans := 0
	for _, line := range input {
		re := regexp.MustCompile(`(\d+)\-(\d+) (\w): (\w+)`)
		matches := re.FindStringSubmatch(line)
		first := utils.ToInt(matches[1])
		second := utils.ToInt(matches[2])
		letter := matches[3][0]
		password := matches[4]

		count := 0
		if password[first-1] == letter {
			count++
		}
		if password[second-1] == letter {
			count++
		}

		if count == 1 {
			ans++
		}
	}

	return ans
}
