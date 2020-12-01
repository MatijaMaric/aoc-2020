package main

import (
	"testing"

	"github.com/MatijaMaric/aoc-2020/utils"
	"github.com/stretchr/testify/require"
)

const example = `1721
979
366
299
675
1456`

func TestPart1(t *testing.T) {
	input := utils.ParseIntLines(example)

	result := part1(input)

	require.Equal(t, 514579, result)
}

func TestPart2(t *testing.T) {
	input := utils.ParseIntLines(example)

	result := part2(input)

	require.Equal(t, 241861950, result)
}
