package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `abc

a
b
c

ab
ac

a
a
a
a

b`

func Test(t *testing.T) {
	require.Equal(t, 11, part1(example))
	require.Equal(t, 6, part2(example))
}
