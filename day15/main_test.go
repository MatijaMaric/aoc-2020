package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require.Equal(t, 436, part1([]int{0, 3, 6}))
	require.Equal(t, 1, part1([]int{1, 3, 2}))
	require.Equal(t, 10, part1([]int{2, 1, 3}))
	require.Equal(t, 27, part1([]int{1, 2, 3}))
	require.Equal(t, 78, part1([]int{2, 3, 1}))
	require.Equal(t, 438, part1([]int{3, 2, 1}))
	require.Equal(t, 1836, part1([]int{3, 1, 2}))

	require.Equal(t, 175594, part2([]int{0, 3, 6}))
	require.Equal(t, 2578, part2([]int{1, 3, 2}))
	require.Equal(t, 3544142, part2([]int{2, 1, 3}))
	require.Equal(t, 261214, part2([]int{1, 2, 3}))
	require.Equal(t, 6895259, part2([]int{2, 3, 1}))
	require.Equal(t, 18, part2([]int{3, 2, 1}))
	require.Equal(t, 362, part2([]int{3, 1, 2}))
}
