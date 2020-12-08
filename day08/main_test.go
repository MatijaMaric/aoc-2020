package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func Test(t *testing.T) {
	require.Equal(t, 5, part1(example1))
	require.Equal(t, 8, part2(example1))
}
