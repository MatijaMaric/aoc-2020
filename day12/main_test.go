package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `F10
N3
F7
R90
F11`

func Test(t *testing.T) {
	input1 := example1

	require.Equal(t, 25, part1(input1))
	require.Equal(t, 286, part2(input1))
}
