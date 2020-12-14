package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

const example2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func Test(t *testing.T) {
	require.Equal(t, 165, part1(example1))
	require.Equal(t, 208, part2(example2))
}
