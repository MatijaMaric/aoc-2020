package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func Test1(t *testing.T) {
	require.Equal(t, 306, part1(example1))
}

func Test2(t *testing.T) {
	require.Equal(t, 291, part2(example1))
}
