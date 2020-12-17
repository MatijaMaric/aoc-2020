package main

import (
	"fmt"

	"github.com/MatijaMaric/aoc-2020/utils"
)

func main() {
	input := utils.GetDay(2020, 17).ToString()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type v3 struct {
	x, y, z int
}

type v4 struct {
	x, y, z, w int
}

func part1(input string) int {
	cubes := make(map[v3]bool)

	lines := utils.SplitLines(input)

	min := v3{0, 0, 0}
	max := v3{len(lines[0]) - 1, len(lines) - 1, 0}

	for y, line := range lines {
		for x, c := range line {
			cubes[v3{x, y, 0}] = c == '#'
		}
	}

	for i := 0; i < 6; i++ {
		newcubes := copygrid(cubes)
		for x := min.x - 1; x <= max.x+1; x++ {
			for y := min.y - 1; y <= max.y+1; y++ {
				for z := min.z - 1; z <= max.z+1; z++ {
					neighbours := 0
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							for dz := -1; dz <= 1; dz++ {
								if dx == 0 && dy == 0 && dz == 0 {
									continue
								}
								if cubes[v3{x + dx, y + dy, z + dz}] {
									neighbours++
								}
							}
						}
					}
					if cubes[v3{x, y, z}] {
						if !(neighbours == 2 || neighbours == 3) {
							newcubes[v3{x, y, z}] = false
						} else {
							newcubes[v3{x, y, z}] = true
						}
					} else {
						if neighbours == 3 {
							newcubes[v3{x, y, z}] = true
						} else {
							newcubes[v3{x, y, z}] = false
						}
					}
				}
			}
		}
		min = v3{min.x - 1, min.y - 1, min.z - 1}
		max = v3{max.x + 1, max.y + 1, max.z + 1}
		cubes = newcubes
	}

	ans := 0
	for _, v := range cubes {
		if v {
			ans++
		}
	}

	return ans
}

func copygrid(grid map[v3]bool) map[v3]bool {
	ans := make(map[v3]bool)
	for k, v := range grid {
		ans[k] = v
	}
	return ans
}

func copygrid4(grid map[v4]bool) map[v4]bool {
	ans := make(map[v4]bool)
	for k, v := range grid {
		ans[k] = v
	}
	return ans
}

func part2(input string) int {
	cubes := make(map[v4]bool)

	lines := utils.SplitLines(input)

	min := v4{0, 0, 0, 0}
	max := v4{len(lines[0]) - 1, len(lines) - 1, 0, 0}

	for y, line := range lines {
		for x, c := range line {
			cubes[v4{x, y, 0, 0}] = c == '#'
		}
	}

	for i := 0; i < 6; i++ {
		newcubes := copygrid4(cubes)
		for x := min.x - 1; x <= max.x+1; x++ {
			for y := min.y - 1; y <= max.y+1; y++ {
				for z := min.z - 1; z <= max.z+1; z++ {
					for w := min.w - 1; w <= max.w+1; w++ {
						neighbours := 0
						for dx := -1; dx <= 1; dx++ {
							for dy := -1; dy <= 1; dy++ {
								for dz := -1; dz <= 1; dz++ {
									for dw := -1; dw <= 1; dw++ {
										if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
											continue
										}
										if cubes[v4{x + dx, y + dy, z + dz, w + dw}] {
											neighbours++
										}
									}
								}
							}
						}
						if cubes[v4{x, y, z, w}] {
							if !(neighbours == 2 || neighbours == 3) {
								newcubes[v4{x, y, z, w}] = false
							} else {
								newcubes[v4{x, y, z, w}] = true
							}
						} else {
							if neighbours == 3 {
								newcubes[v4{x, y, z, w}] = true
							} else {
								newcubes[v4{x, y, z, w}] = false
							}
						}
					}
				}
			}
		}
		min = v4{min.x - 1, min.y - 1, min.z - 1, min.w - 1}
		max = v4{max.x + 1, max.y + 1, max.z + 1, max.z + 1}
		cubes = newcubes
	}

	ans := 0
	for _, v := range cubes {
		if v {
			ans++
		}
	}

	return ans
}
