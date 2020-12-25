package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEx1(t *testing.T) {
	require.Equal(t, 14897079, part1([]int{5764801, 17807724}))
}
