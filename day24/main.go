package main

import (
	"fmt"
	"regexp"

	"github.com/cheggaaa/pb/v3"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 24).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type v3 struct {
	x, y, z int
}

func part1(input string) int {
	grid := makeGrid(input)
	return countBlack(grid)
}

func part2(input string) int {
	grid := makeGrid(input)

	min, max := v3{utils.MaxInt, utils.MaxInt, utils.MaxInt}, v3{utils.MinInt, utils.MinInt, utils.MinInt}
	for k := range grid {
		min = minV3(min, k)
		max = maxV3(max, k)
	}
	min = min.add(v3{-1, -1, -1})
	max = max.add(v3{1, 1, 1})

	bar := pb.StartNew(100)
	for i := 0; i < 100; i++ {
		newGrid := make(map[v3]bool)

		for x := min.x; x <= max.x; x++ {
			for y := min.y; y <= max.y; y++ {
				for z := min.z; z <= max.z; z++ {
					pos := v3{x, y, z}
					if _, ok := grid[pos]; !ok {
						grid[pos] = false
					}
					cnt := 0
					for _, next := range pos.neighbors() {
						if grid[next] {
							cnt++
						}
					}
					if grid[pos] && (cnt == 0 || cnt > 2) {
						newGrid[pos] = false
						continue
					}
					if !grid[pos] && cnt == 2 {
						newGrid[pos] = true
						continue
					}
					newGrid[pos] = grid[pos]
				}
			}
		}

		grid = newGrid
		min = min.add(v3{-1, -1, -1})
		max = max.add(v3{1, 1, 1})
		bar.Increment()
	}
	bar.Finish()

	return countBlack(grid)
}

func countBlack(grid map[v3]bool) int {
	ans := 0
	for _, v := range grid {
		if v {
			ans++
		}
	}

	return ans
}

func makeGrid(input string) map[v3]bool {
	lines := utils.SplitLines(input)

	grid := make(map[v3]bool)

	for _, line := range lines {
		dirs := splitDirs(line)
		curr := v3{0, 0, 0}
		for _, dir := range dirs {
			switch dir {
			case "e":
				curr = v3{curr.x + 1, curr.y - 1, curr.z}
			case "se":
				curr = v3{curr.x, curr.y - 1, curr.z + 1}
			case "sw":
				curr = v3{curr.x - 1, curr.y, curr.z + 1}
			case "w":
				curr = v3{curr.x - 1, curr.y + 1, curr.z}
			case "nw":
				curr = v3{curr.x, curr.y + 1, curr.z - 1}
			case "ne":
				curr = v3{curr.x + 1, curr.y, curr.z - 1}
			}
		}
		if v, ok := grid[curr]; ok {
			grid[curr] = !v
		} else {
			grid[curr] = true
		}
	}
	return grid
}

func minV3(a, b v3) v3 {
	return v3{utils.Min(a.x, b.x), utils.Min(a.y, b.y), utils.Min(a.z, b.z)}
}

func maxV3(a, b v3) v3 {
	return v3{utils.Max(a.x, b.x), utils.Max(a.y, b.y), utils.Max(a.z, b.z)}
}

func (curr *v3) add(other v3) v3 {
	return v3{curr.x + other.x, curr.y + other.y, curr.z + other.z}
}

func (curr *v3) neighbors() []v3 {
	return []v3{
		curr.add(v3{1, -1, 0}),
		curr.add(v3{0, -1, 1}),
		curr.add(v3{-1, 0, 1}),
		curr.add(v3{-1, 1, 0}),
		curr.add(v3{0, 1, -1}),
		curr.add(v3{1, 0, -1}),
	}
}

func splitDirs(input string) []string {
	re := regexp.MustCompile(`(e|se|sw|w|nw|ne)`)

	return re.FindAllString(input, -1)
}
