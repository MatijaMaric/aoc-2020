package main

import (
	"testing"

	"github.com/MatijaMaric/aoc-2020/utils"
	"github.com/stretchr/testify/require"
)

const example = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestPart1(t *testing.T) {
	input := utils.SplitLines(example)

	result := part1(input)

	require.Equal(t, 7, result)
}

func TestPart2(t *testing.T) {
	input := utils.SplitLines(example)

	result := part2(input)

	require.Equal(t, 336, result)
}
