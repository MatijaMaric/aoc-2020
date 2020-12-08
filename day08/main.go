package main

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 8).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	instructions := strings.Split(input, "\n")
	acc := 0
	pc := 0

	executed := mapset.NewSet()

	for !executed.Contains(pc) {
		if pc >= len(instructions) {
			break
		}
		executed.Add(pc)
		op, arg := parseInstruction(instructions[pc])

		if op == "nop" {
			pc++
		}
		if op == "jmp" {
			pc += arg
		}
		if op == "acc" {
			acc += arg
			pc++
		}
	}

	return acc
}

func part2(input string) int {
	instructions := strings.Split(input, "\n")

	for i := 0; i < len(instructions); i++ {
		if op, _ := parseInstruction(instructions[i]); op == "acc" {
			continue
		}
		executed := mapset.NewSet()

		acc := 0
		pc := 0
		for !executed.Contains(pc) {
			if pc >= len(instructions) {
				return acc
			}
			executed.Add(pc)
			op, arg := parseInstruction(instructions[pc])

			if op == "nop" {
				if pc == i {
					pc += arg
				} else {
					pc++
				}
			}
			if op == "jmp" {
				if pc == i {
					pc++
				} else {
					pc += arg
				}
			}
			if op == "acc" {
				acc += arg
				pc++
			}
		}
	}

	return -1
}

func parseInstruction(instruction string) (op string, arg int) {
	split := strings.Split(instruction, " ")
	op = split[0]
	arg = utils.ToInt(split[1])
	return
}
