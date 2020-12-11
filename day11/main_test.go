package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func Test(t *testing.T) {
	input1 := example1

	require.Equal(t, 37, part1(input1))
	require.Equal(t, 26, part2(input1))
}
