package main

import (
	"testing"

	"github.com/MatijaMaric/aoc-2020/utils"

	"github.com/stretchr/testify/require"
)

const example = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func Test(t *testing.T) {
	input := utils.ParseIntLines(example)

	require.Equal(t, 127, part1(input, 5))
	require.Equal(t, 62, part2(input, 127))
}
