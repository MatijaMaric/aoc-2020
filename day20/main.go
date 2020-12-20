package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/MatijaMaric/aoc-2020/utils"
)

const monster = "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "

func main() {
	input := utils.GetDay(2020, 20).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	tiles := parseTiles(input)
	edges := findEdges(tiles)

	ans := 1
	for _, corner := range findCorners(edges) {
		ans *= corner
	}

	return ans
}

func part2(input string) int {
	tiles := parseTiles(input)
	var edges map[int][]int
	var corners []int
	type node struct {
		from, to, dfrom, dto int
	}
	var g map[int]map[int]node

	var tilesize int
	for _, v := range tiles {
		tilesize = len(v)
		break
	}

	calcEdges := func() {
		edges = findEdges(tiles)
		corners = findCorners(edges)
		g = make(map[int]map[int]node)
		for from := range edges {
			g[from] = make(map[int]node)
			for to := range edges {
				if from == to {
					continue
				}
				for da, ea := range edges[from] {
					for db, eb := range edges[to] {
						if ea == eb {
							g[from][to] = node{from, to, da, db}
						}
					}
				}
			}
		}
	}

	calcEdges()

	gridsize := int(math.Sqrt(float64(len(tiles))))

	tilegrid := make([][]int, gridsize)
	for i := 0; i < gridsize; i++ {
		tilegrid[i] = make([]int, gridsize)
	}

	tilegrid[0][0] = corners[0]
	rotate := func(id int) {
		tiles[id] = transform(tiles[id], 1)
		calcEdges()
	}
	flip := func(id int) {
		tiles[id] = transform(tiles[id], 4)
		calcEdges()
	}
	for {
		from := tilegrid[0][0]
		hasbot := false
		hasright := false
		for _, v := range g[from] {
			if v.dfrom%4 == 1 {
				hasright = true
			}
			if v.dfrom%4 == 2 {
				hasbot = true
			}
		}
		if hasright && hasbot {
			break
		} else {
			rotate(from)
		}
	}

	for i := 1; i < gridsize; i++ {
		from := tilegrid[0][i-1]
		for to, v := range g[from] {
			if v.dfrom%4 == 1 {
				tilegrid[0][i] = to
				if g[from][to].dfrom > 3 && g[from][to].dto > 3 {
					flip(to)
				}
				for {
					if g[from][to].dfrom%4 == 1 && g[from][to].dto == 3 {
						break
					}
					rotate(to)
				}
				break
			}
		}
	}

	for i := 1; i < gridsize; i++ {
		for j := 0; j < gridsize; j++ {
			from := tilegrid[i-1][j]
			for to, v := range g[from] {
				if v.dfrom%4 == 2 {
					tilegrid[i][j] = to
					if g[from][to].dfrom > 3 && g[from][to].dto > 3 {
						flip(to)
					}
					for {
						if g[from][to].dfrom%4 == 2 && g[from][to].dto == 0 {
							break
						}
						rotate(to)
					}
				}
			}
		}
	}

	size := gridsize * (tilesize - 2)
	grid := make([][]bool, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]bool, size)
	}

	for tx := 0; tx < gridsize; tx++ {
		for ty := 0; ty < gridsize; ty++ {
			id := tilegrid[tx][ty]
			for x := 1; x < tilesize-1; x++ {
				for y := 1; y < tilesize-1; y++ {
					grid[tx*(tilesize-2)+x-1][ty*(tilesize-2)+y-1] = tiles[id][x][y]
				}
			}
		}
	}

	return countMonster(grid)
}

func countMonster(grid [][]bool) int {
	monsterMask := utils.ParseBoolGrid(monster)
	mh := len(monsterMask)
	mw := len(monsterMask[0])

	for d := 0; d < 8; d++ {
		cnt := 0
		transformed := transform(grid, d)

		monsters := 0

		occupied := make([][]bool, len(grid))
		for i := 0; i < len(grid); i++ {
			occupied[i] = make([]bool, len(grid))
			for j := 0; j < len(grid); j++ {
				occupied[i][j] = transformed[i][j]
			}
		}

		for gx := 0; gx < len(grid)-mh; gx++ {
			for gy := 0; gy < len(grid)-mw; gy++ {
				found := true

				for mx := 0; mx < mh; mx++ {
					for my := 0; my < mw; my++ {
						if !transformed[gx+mx][gy+my] && monsterMask[mx][my] {
							found = false
						}
					}
				}

				if found {
					monsters++
					for mx := 0; mx < mh; mx++ {
						for my := 0; my < mw; my++ {
							if monsterMask[mx][my] {
								occupied[gx+mx][gy+my] = false
							}
						}
					}
				}
			}
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid); j++ {
				if occupied[i][j] {
					cnt++
				}
			}
		}
		if monsters > 0 {
			return cnt
		}
	}

	return -1
}

func printGrid(grid [][]bool) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printOverlaidGrid(grid [][]bool, occupied [][]bool) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] {
				if occupied[i][j] {
					fmt.Print("#")
				} else {
					fmt.Print("O")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseTiles(input string) map[int][][]bool {
	tiles := make(map[int][][]bool)

	re := regexp.MustCompile(`Tile (\d+):`)

	for _, part := range strings.Split(input, "\n\n") {
		id := utils.ToInt(re.FindStringSubmatch(part)[1])
		tiles[id] = utils.ParseBoolGrid(strings.SplitN(part, "\n", 2)[1])
	}
	return tiles
}

func findEdges(tiles map[int][][]bool) map[int][]int {
	edges := make(map[int][]int)
	for id, tile := range tiles {
		edges[id] = getEdges(tile)
	}

	return edges
}

func countEdges(edges map[int][]int) map[int]int {
	count := make(map[int]int)

	for _, tile := range edges {
		for _, edge := range tile {
			if _, ok := count[edge]; !ok {
				count[edge] = 1
			} else {
				count[edge]++
			}
			if count[edge] > 2 {
				panic("i'm too lazy to solve this if there are more than two matching edges")
			}
		}
	}

	return count
}

func findCorners(edges map[int][]int) []int {
	count := countEdges(edges)

	var ans []int
	for id, tile := range edges {
		c := 0
		for _, edge := range tile {
			if count[edge] == 1 {
				c++
			}
		}
		if c > 2 {
			ans = append(ans, id)
		}
	}
	return ans
}

func getEdges(tile [][]bool) []int {
	ans := make([]int, 8)
	size := len(tile) - 1
	for i := 0; i < len(tile); i++ {
		ans[0] += ternary(tile[0][i], 1<<(size-i), 0)
		ans[1] += ternary(tile[i][size], 1<<(size-i), 0)
		ans[2] += ternary(tile[size][size-i], 1<<(size-i), 0)
		ans[3] += ternary(tile[size-i][0], 1<<(size-i), 0)
		ans[4] += ternary(tile[0][i], 1<<i, 0)
		ans[5] += ternary(tile[i][size], 1<<i, 0)
		ans[6] += ternary(tile[size][size-i], 1<<i, 0)
		ans[7] += ternary(tile[size-i][0], 1<<i, 0)
	}
	return ans
}

func ternary(cond bool, valtru, valfal int) int {
	if cond {
		return valtru
	}
	return valfal
}

func toNumber(input string) int {
	ans := 0
	for i, c := range input {
		if c == '#' {
			ans += 1 << (len(input) - 1 - i)
		}
	}
	return ans
}

func reverse(input string) string {
	var ans strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		ans.WriteByte(input[i])
	}
	return ans.String()
}

func transform(tile [][]bool, d int) [][]bool {
	ans := make([][]bool, len(tile))
	for i := 0; i < len(tile); i++ {
		ans[i] = make([]bool, len(tile[i]))
	}

	for i := 0; i < len(tile); i++ {
		for j := 0; j < len(tile); j++ {
			ans[i][j] = transformPos(tile, i, j, d)
		}
	}

	return ans
}

func transformPos(grid [][]bool, i, j, d int) bool {
	size := len(grid) - 1
	if d == 0 {
		return grid[i][j]
	}
	if d == 1 {
		return grid[size-j][i]
	}
	if d == 2 {
		return grid[size-i][size-j]
	}
	if d == 3 {
		return grid[j][size-i]
	}
	if d == 4 {
		return grid[i][size-j]
	}
	if d == 5 {
		return grid[size-j][size-i]
	}
	if d == 6 {
		return grid[size-i][j]
	}
	if d == 7 {
		return grid[j][i]
	}
	panic("invalid transform")
}
