package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBsp(t *testing.T) {
	x, y := bsp("FBFBBFFRLR")
	require.Equal(t, 44, x)
	require.Equal(t, 5, y)
	require.Equal(t, 357, x*8+y)
	x, y = bsp("BFFFBBFRRR")
	require.Equal(t, 70, x)
	require.Equal(t, 7, y)
	require.Equal(t, 567, x*8+y)
	x, y = bsp("FFFBBBFRRR")
	require.Equal(t, 14, x)
	require.Equal(t, 7, y)
	require.Equal(t, 119, x*8+y)
	x, y = bsp("BBFFBBFRLL")
	require.Equal(t, 102, x)
	require.Equal(t, 4, y)
	require.Equal(t, 820, x*8+y)

}
