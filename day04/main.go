package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 4).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	passports := strings.Split(input, "\n\n")
	valid := 0
	for _, passport := range passports {
		pairs := strings.Fields(passport)
		var fields = map[string]string{}
		for _, pair := range pairs {
			field := strings.Split(pair, ":")
			fields[field[0]] = field[1]
		}
		delete(fields, "cid")
		if len(fields) == 7 {
			valid++
		}
	}
	return valid
}

func byr(field string) bool {
	val := utils.ToInt(field)
	if val < 1920 || val > 2002 {
		return false
	}
	return true
}

func iyr(field string) bool {
	val := utils.ToInt(field)
	if val < 2010 || val > 2020 {
		return false
	}
	return true
}

func eyr(field string) bool {
	val := utils.ToInt(field)
	if val < 2020 || val > 2030 {
		return false
	}
	return true
}

func hgt(field string) bool {
	hgt := regexp.MustCompile(`(\d{2,3}(cm)|(in))`)
	if !hgt.MatchString(field) {
		return false
	}
	unit := field[len(field)-2:]
	val := utils.ToInt(field[0 : len(field)-2])
	if (unit == "cm" && (val < 150 || val > 193)) || (unit == "in" && (val < 59 || val > 76)) {
		return false
	}
	return true
}

func hcl(field string) bool {
	hcl := regexp.MustCompile(`(#[\dabcdef]{6})`)

	if !hcl.MatchString(field) {
		return false
	}
	return true
}

func ecl(field string) bool {
	if !(field == "amb" || field == "blu" || field == "brn" || field == "gry" || field == "grn" || field == "hzl" || field == "oth") {
		return false
	}
	return true
}

func pid(field string) bool {
	pid := regexp.MustCompile(`(\d+)`)

	if !pid.MatchString(field) || len(field) != 9 {
		return false
	}
	return true
}

func part2(input string) int {
	passports := strings.Split(input, "\n\n")
	valid := 0

	for _, passport := range passports {
		pairs := strings.Fields(passport)
		var fields = map[string]string{}
		isValid := true
		for _, pair := range pairs {
			field := strings.Split(pair, ":")
			fields[field[0]] = field[1]
			switch field[0] {
			case "byr":
				isValid = isValid && byr(field[1])
				break
			case "iyr":
				isValid = isValid && iyr(field[1])
				break
			case "eyr":
				isValid = isValid && eyr(field[1])
				break
			case "hgt":
				isValid = isValid && hgt(field[1])
				break
			case "hcl":
				isValid = isValid && hcl(field[1])

				break
			case "ecl":
				isValid = isValid && ecl(field[1])
				break
			case "pid":
				isValid = isValid && pid(field[1])
				break
			}
		}
		delete(fields, "cid")
		if isValid && len(fields) == 7 {
			valid++
			fmt.Println("valid", pairs)
		} else {
			fmt.Println("invalid", pairs)
		}

	}
	return valid
}
