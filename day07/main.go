package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 7).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type content struct {
	color string
	count int
}

func part1(input string) int {
	bags := strings.Split(input, "\n")

	colors := make(map[string][]string)
	contains := make(map[string]bool)

	for _, line := range bags {
		split := strings.Split(line[0:len(line)-1], " bags contain ")
		bag := split[0]
		others := strings.Split(split[1], ", ")
		colors[bag] = make([]string, len(others))
		contains[bag] = false
		for i, other := range others {
			colors[bag][i] = other[strings.Index(other, " ")+1 : strings.Index(other, " bag")]
			if colors[bag][i] == "shiny gold" {
				contains[bag] = true
			}
		}
	}

	changed := true
	for changed {
		changed = false
		for color, others := range colors {
			if contains[color] {
				continue
			}
			for _, other := range others {
				if contains[other] {
					contains[color] = true
					changed = true
				}
			}
		}
	}

	ans := 0

	for _, yes := range contains {
		if yes {
			ans++
		}
	}

	return ans
}

func part2(input string) int {
	bags := strings.Split(input, "\n")

	colors := make(map[string][]content)

	for _, line := range bags {
		split := strings.Split(line[0:len(line)-1], " bags contain ")
		bag := split[0]
		others := strings.Split(split[1], ", ")
		colors[bag] = make([]content, len(others))
		for i, other := range others {
			if other == "no other bags" {
				continue
			}
			colors[bag][i] = content{
				other[strings.Index(other, " ")+1 : strings.Index(other, " bag")],
				utils.ToInt(other[0:strings.Index(other, " ")]),
			}
		}
	}

	return countBags(colors, "shiny gold") - 1
}

func countBags(colors map[string][]content, color string) int {
	ans := 1

	if content, ok := colors[color]; ok {
		for _, other := range content {
			ans += countBags(colors, other.color) * other.count
		}
	}

	return ans
}
