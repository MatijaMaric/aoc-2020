package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `2 * 3 + (4 * 5)`
const example2 = `5 + (8 * 3 + 9 + 3 * 4 * 3)`
const example3 = `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`
const example4 = `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`

func Test1(t *testing.T) {
	require.Equal(t, 26, solve1(example1))
	require.Equal(t, 437, solve1(example2))
	require.Equal(t, 12240, solve1(example3))
	require.Equal(t, 13632, solve1(example4))
}

func Test2(t *testing.T) {
	require.Equal(t, 46, solve2(example1))
	require.Equal(t, 1445, solve2(example2))
	require.Equal(t, 669060, solve2(example3))
	require.Equal(t, 23340, solve2(example4))
}
