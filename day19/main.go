package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 19).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type rule struct {
	this, that []int
	re         string
}

func part1(input string) int {
	return solve(input, false)
}

func part2(input string) int {
	return solve(input, true)
}

func solve(input string, part2 bool) int {
	parts := strings.Split(input, "\n\n")
	messages := utils.SplitLines(parts[1])

	rules := make(map[int]rule)

	a := -1
	b := -1

	for _, line := range utils.SplitLines(parts[0]) {
		p := strings.Split(line, ": ")
		idx := utils.ToInt(p[0])
		if p[1] == `"a"` {
			a = idx
			rules[a] = rule{nil, nil, "a"}
		} else if p[1] == `"b"` {
			b = idx
			rules[b] = rule{nil, nil, "b"}
		} else {
			p = strings.Split(p[1], " | ")
			if len(p) == 1 {
				rules[idx] = rule{parseList(p[0]), nil, ""}
			} else {
				rules[idx] = rule{parseList(p[0]), parseList(p[1]), ""}
			}
		}
	}

	found := mapset.NewSet()
	found.Add(a)
	found.Add(b)

	for !found.Contains(0) {
	outer:
		for k, v := range rules {
			if part2 {
				if k == 8 && found.Contains(42) {
					found.Add(8)
					rules[k] = rule{v.this, v.that, "(" + rules[42].re + ")+"}
					continue
				}
				if k == 11 && found.Contains(42) && found.Contains(31) {
					found.Add(11)
					// go pls implement recursive regex
					var ohboy strings.Builder
					ohboy.WriteString("(" + rules[42].re + rules[31].re)
					for i := 2; i < 20; i++ {
						idx := strconv.Itoa(i)
						ohboy.WriteString("|" + rules[42].re + "{" + idx + "}" + rules[31].re + "{" + idx + "}")
					}
					ohboy.WriteString(")")
					rules[k] = rule{v.this, v.that, ohboy.String()}
					continue
				}
			}
			if found.Contains(k) {
				continue
			}
			var re strings.Builder
			re.WriteString("(")
			for _, x := range v.this {
				if !found.Contains(x) {
					continue outer
				}
				re.WriteString(rules[x].re)
			}
			if len(v.that) > 0 {
				re.WriteString("|")
				for _, x := range v.that {
					if !found.Contains(x) {
						continue outer
					}
					re.WriteString(rules[x].re)
				}
			}
			re.WriteString(")")
			rules[k] = rule{v.this, v.that, re.String()}
			found.Add(k)
		}
	}

	re := regexp.MustCompile("^" + rules[0].re + "$")
	ans := 0
	for _, message := range messages {
		if re.MatchString(message) {
			ans++
		}
	}

	return ans
}

func parseList(input string) []int {
	list := strings.Split(input, " ")
	ans := make([]int, len(list))
	for i, x := range list {
		ans[i] = utils.ToInt(x)
	}
	return ans
}
