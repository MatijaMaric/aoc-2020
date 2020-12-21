package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func Test1(t *testing.T) {
	require.Equal(t, 5, part1(example1))
	require.Equal(t, 273, part2(example1))
}
