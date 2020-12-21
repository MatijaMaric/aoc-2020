package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 21).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	lines := utils.SplitLines(input)
	allergens := make(map[string]mapset.Set)

	re := regexp.MustCompile(`([^\()]+) \(contains ([^\)]+)\)`)

	var allings []string

	for _, line := range lines {
		grp := re.FindStringSubmatch(line)

		ings := strings.Split(grp[1], " ")
		alls := strings.Split(grp[2], ", ")
		allings = append(allings, ings...)

		ingset := mapset.NewSet()

		for _, ing := range ings {
			ingset.Add(ing)
		}

		for _, all := range alls {
			if _, ok := allergens[all]; !ok {
				allergens[all] = ingset
			} else {
				allergens[all] = allergens[all].Intersect(ingset)
			}
		}
	}

	found := mapset.NewSet()
	foundings := make(map[string]string)

	for {
		for k := range allergens {
			allergens[k] = allergens[k].Difference(found)
			if allergens[k].Cardinality() == 1 {
				ing := allergens[k].Pop()
				found.Add(ing)
				foundings[k] = ing.(string)
			}
		}

		if found.Cardinality() == len(allergens) {
			break
		}
	}

	ans := 0
	for _, ing := range allings {
		if !found.Contains(ing) {
			ans++
		}
	}

	keys := make([]string, len(allergens))
	idx := 0
	for key := range allergens {
		keys[idx] = key
		idx++
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Print(foundings[key])
		fmt.Print(",")
	}
	fmt.Println()

	return ans
}

func part2(input string) int {
	return -1
}
