package main

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 16).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	parts := strings.Split(input, "\n\n")

	valids := mapset.NewSet()
	for _, rule := range strings.Split(parts[0], "\n") {
		ranges := strings.Split(strings.Split(rule, ": ")[1], " or ")
		for _, r := range ranges {
			rg := strings.Split(r, "-")
			from := utils.ToInt(rg[0])
			to := utils.ToInt(rg[1])

			for i := from; i <= to; i++ {
				valids.Add(i)
			}
		}
	}

	err := 0

	for i, ticket := range strings.Split(parts[2], "\n") {
		if i == 0 {
			continue
		}
		for _, num := range strings.Split(ticket, ",") {
			numi := utils.ToInt(num)
			if !valids.Contains(numi) {
				err += numi
			}
		}
	}

	return err
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	myticket := strings.Split(strings.Split(parts[1], "\n")[1], ",")
	alltickets := strings.Split(parts[2], "\n")[1:]

	var fields []string
	allvalids := mapset.NewSet()
	valids := make(map[string]mapset.Set)

	for _, rule := range strings.Split(parts[0], "\n") {
		ruleparts := strings.Split(rule, ": ")
		fields = append(fields, ruleparts[0])
		ranges := strings.Split(ruleparts[1], " or ")
		valids[ruleparts[0]] = mapset.NewSet()
		for _, r := range ranges {
			rg := strings.Split(r, "-")
			from := utils.ToInt(rg[0])
			to := utils.ToInt(rg[1])

			for i := from; i <= to; i++ {
				allvalids.Add(i)
				valids[ruleparts[0]].Add(i)
			}
		}
	}

	var tickets [][]int

	for _, ticket := range alltickets {
		valid := true
		ticketi := []int{}
		for _, num := range strings.Split(ticket, ",") {
			numi := utils.ToInt(num)
			ticketi = append(ticketi, numi)
			if !allvalids.Contains(numi) {
				valid = false
			}
		}
		if valid {
			tickets = append(tickets, ticketi)
		}
	}

	could := make([]mapset.Set, len(tickets[0]))

	for i := 0; i < len(tickets[0]); i++ {
		could[i] = mapset.NewSet()
		for k := range valids {
			could[i].Add(k)
		}
		for _, ticket := range tickets {
			for k, v := range valids {
				if !v.Contains(ticket[i]) {
					could[i].Remove(k)
				}
			}
		}
	}

	fieldmap := make(map[string]int)
	used := mapset.NewSet()

	for {
		for i := 0; i < len(could); i++ {
			could[i] = could[i].Difference(used)
			if could[i].Cardinality() == 1 {
				fieldname := could[i].ToSlice()[0].(string)
				fieldmap[fieldname] = i
				used.Add(could[i].ToSlice()[0])
			}
		}
		if used.Cardinality() == len(fields) {
			break
		}
	}

	ans := 1
	for k, v := range fieldmap {
		if strings.Contains(k, "departure") {
			ans *= utils.ToInt(myticket[v])
		}
	}

	return ans
}
