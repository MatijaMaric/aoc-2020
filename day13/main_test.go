package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `939
7,13,x,x,59,x,31,19`

const example2 = `69
17,x,13,19`

const example3 = `69
67,7,59,61`

const example4 = `69
67,x,7,59,61`

const example5 = `69
67,7,x,59,61`

const example6 = `69
1789,37,47,1889`

func Test(t *testing.T) {
	require.Equal(t, 295, part1(example1))
	require.Equal(t, 1068781, part2(example1))
	require.Equal(t, 3417, part2(example2))
	require.Equal(t, 754018, part2(example3))
	require.Equal(t, 779210, part2(example4))
	require.Equal(t, 1261476, part2(example5))
	require.Equal(t, 1202161486, part2(example6))
}
