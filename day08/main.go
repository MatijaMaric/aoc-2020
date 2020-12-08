package main

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set"

	"github.com/MatijaMaric/aoc-2020/utils"
)

type Instruction struct {
	op  string
	arg int
}

func main() {
	input := utils.GetDay(2020, 8).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	instructions := parseInstructions(input)

	_, acc := solve(instructions)

	return acc
}

func part2(input string) int {
	instructions := parseInstructions(input)

	for i := 0; i < len(instructions); i++ {
		prev := instructions[i].op
		if instructions[i].op == "jmp" {
			instructions[i].op = "nop"
		} else if instructions[i].op == "nop" {
			instructions[i].op = "jmp"
		} else {
			continue
		}

		if halt, acc := solve(instructions); halt {
			return acc
		}

		instructions[i].op = prev
	}

	return -1
}

func parseInstructions(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}
	return instructions
}

func parseInstruction(instruction string) Instruction {
	split := strings.Split(instruction, " ")
	return Instruction{split[0], utils.ToInt(split[1])}
}

func solve(instructions []Instruction) (halt bool, acc int) {
	executed := mapset.NewSet()

	acc = 0
	pc := 0
	for !executed.Contains(pc) {
		if pc >= len(instructions) {
			return true, acc
		}
		executed.Add(pc)
		instruction := instructions[pc]

		switch instruction.op {
		case "nop":
			pc++
			break
		case "jmp":
			pc += instruction.arg
			break
		case "acc":
			pc++
			acc += instruction.arg
		}
	}
	return false, acc
}
