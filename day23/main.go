package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"

	"github.com/MatijaMaric/aoc-2020/utils"
	mapset "github.com/deckarep/golang-set"
)

func main() {
	input := utils.GetDay(2020, 23).ToString()

	fmt.Println(part1(input, 100))
	fmt.Println(part2(input))
}

func part1(input string, moves int) string {
	max := len(input)
	cups := ring.New(len(input))
	for _, c := range input {
		cups.Value = int(c - '0')
		cups = cups.Next()
	}

	shuffle(cups, max, moves)

	for cups.Value != 1 {
		cups = cups.Next()
	}
	cups = cups.Next()

	var ans strings.Builder
	for cups.Value != 1 {
		ans.WriteString(strconv.Itoa(cups.Value.(int)))
		cups = cups.Next()
	}

	return ans.String()
}

func part2(input string) int {
	max := 1_000_000
	moves := 10_000_000

	cups := ring.New(max)
	for _, c := range input {
		cups.Value = int(c - '0')
		cups = cups.Next()
	}
	for i := len(input) + 1; i <= max; i++ {
		cups.Value = i
		cups = cups.Next()
	}

	cups = shuffle(cups, max, moves)

	for cups.Value != 1 {
		cups = cups.Next()
	}
	cups = cups.Next()
	ans := cups.Value.(int)
	cups = cups.Next()
	ans *= cups.Value.(int)

	return ans
}

func shuffle(cups *ring.Ring, max, moves int) *ring.Ring {
	nodes := make(map[int]*ring.Ring)
	for i := 0; i < max; i++ {
		nodes[cups.Value.(int)] = cups
		cups = cups.Next()
	}
	bar := pb.StartNew(moves)
	for i := 0; i < moves; i++ {
		bar.Increment()
		pickup := cups.Unlink(3)
		picked := mapset.NewSet()
		for j := 0; j < 3; j++ {
			picked.Add(pickup.Value)
			pickup = pickup.Next()
		}
		dest := cups.Value.(int)
		for {
			dest--
			if dest == 0 {
				dest = max
			}
			if !picked.Contains(dest) {
				break
			}
		}

		current := cups
		cups = nodes[dest]
		cups = cups.Link(pickup)
		cups = current.Next()
	}
	bar.Finish()
	return cups
}

func printRing(ring *ring.Ring) {
	ring.Do(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println()
}
