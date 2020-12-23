package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = "389125467"

func Test1(t *testing.T) {
	require.Equal(t, "92658374", part1(example1, 10))
}

func Test2(t *testing.T) {
	require.Equal(t, "67384529", part1(example1, 100))
}

func Test3(t *testing.T) {
	require.Equal(t, 149245887792, part2(example1))
}
