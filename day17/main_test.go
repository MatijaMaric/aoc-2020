package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `.#.
..#
###`

func Test(t *testing.T) {
	require.Equal(t, 112, part1(example1))
}
