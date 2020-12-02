package main

import (
	"testing"

	"github.com/MatijaMaric/aoc-2020/utils"
	"github.com/stretchr/testify/require"
)

const example = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestPart1(t *testing.T) {
	input := utils.SplitLines(example)

	result := part1(input)

	require.Equal(t, 2, result)
}

func TestPart2(t *testing.T) {
	input := utils.SplitLines(example)

	result := part2(input)

	require.Equal(t, 1, result)
}
