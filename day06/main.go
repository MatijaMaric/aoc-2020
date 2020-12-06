package main

import (
	"fmt"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
	mapset "github.com/deckarep/golang-set"
)

func main() {
	input := utils.GetDay(2020, 6).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	groups := strings.Split(input, "\n\n")
	ans := 0
	for _, group := range groups {
		people := strings.Split(group, "\n")
		groupanswers := mapset.NewSet()
		for _, person := range people {
			personanswers := mapset.NewSet()
			for _, answer := range person {
				personanswers.Add(answer)
			}
			groupanswers = groupanswers.Union(personanswers)
		}
		ans += groupanswers.Cardinality()
	}
	return ans
}

func part2(input string) int {
	groups := strings.Split(input, "\n\n")
	ans := 0
	for _, group := range groups {
		people := strings.Split(group, "\n")
		groupanswers := mapset.NewSet()
		for _, answer := range people[0] {
			groupanswers.Add(answer)
		}
		for _, person := range people {
			personanswers := mapset.NewSet()
			for _, answer := range person {
				personanswers.Add(answer)
			}
			groupanswers = groupanswers.Intersect(personanswers)
		}
		ans += groupanswers.Cardinality()
	}
	return ans
}
