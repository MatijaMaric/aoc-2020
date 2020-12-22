package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 22).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	players := strings.Split(input, "\n\n")
	player1 := utils.SplitLines(players[0])[1:]
	player2 := utils.SplitLines(players[1])[1:]

	deck1 := listToInt(player1)
	deck2 := listToInt(player2)

	_, score := play(deck1, deck2)

	return score
}

func copyArr(arr []int) []int {
	newarr := make([]int, len(arr))
	for i, x := range arr {
		newarr[i] = x
	}
	return newarr
}

func play(player1, player2 []int) (winner, score int) {
	return play2(player1, player2, false)
}

func play2(player1, player2 []int, recurse bool) (winner, score int) {
	seen1 := make(map[string]bool)
	seen2 := make(map[string]bool)

	deck1, deck2 := copyArr(player1), copyArr(player2)
	for {

		if seen1[a2s(deck1)] || seen2[a2s(deck2)] {
			return 1, 0
		}

		seen1[a2s(deck1)] = true
		seen2[a2s(deck2)] = true

		if len(deck1) == 0 || len(deck2) == 0 {
			break
		}
		card1 := deck1[0]
		card2 := deck2[0]

		if recurse && card1 < len(deck1) && card2 < len(deck2) {
			winner, _ := play2(deck1[1:card1+1], deck2[1:card2+1], true)

			if winner == 1 {
				deck1 = append(deck1[1:], card1, card2)
				deck2 = deck2[1:]
			} else {
				deck2 = append(deck2[1:], card2, card1)
				deck1 = deck1[1:]
			}
			continue
		}

		if card1 > card2 {
			deck1 = append(deck1[1:], card1, card2)
			deck2 = deck2[1:]
		} else if card2 > card1 {
			deck2 = append(deck2[1:], card2, card1)
			deck1 = deck1[1:]
		}

	}
	var deck []int
	if len(deck1) > len(deck2) {
		deck = deck1
		winner = 1
	} else {
		deck = deck2
		winner = 2
	}

	score = 0
	for i := len(deck) - 1; i >= 0; i-- {
		score += deck[i] * (len(deck) - i)
	}

	return
}

func a2s(list []int) string {
	var ans strings.Builder
	for _, x := range list {
		ans.WriteString(strconv.Itoa(x) + ",")
	}
	return ans.String()
}

func listToInt(list []string) []int {
	ans := make([]int, len(list))
	for i, x := range list {
		ans[i] = utils.ToInt(x)
	}
	return ans
}

func part2(input string) int {
	players := strings.Split(input, "\n\n")
	player1 := utils.SplitLines(players[0])[1:]
	player2 := utils.SplitLines(players[1])[1:]

	deck1 := listToInt(player1)
	deck2 := listToInt(player2)

	_, score := play2(deck1, deck2, true)

	return score
}
