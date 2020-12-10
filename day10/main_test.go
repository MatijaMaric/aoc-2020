package main

import (
	"testing"

	"github.com/MatijaMaric/aoc-2020/utils"

	"github.com/stretchr/testify/require"
)

const example1 = `16
10
15
5
1
11
7
19
6
12
4`

const example2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func Test(t *testing.T) {
	input1 := utils.ParseIntLines(example1)
	input2 := utils.ParseIntLines(example2)

	require.Equal(t, 35, part1(input1))
	require.Equal(t, 220, part1(input2))

	require.Equal(t, 8, part2(input1))
	require.Equal(t, 19208, part2(input2))
}
