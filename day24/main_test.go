package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example1 = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestDirs(t *testing.T) {
	require.Equal(t, []string{"e", "se", "ne", "e"}, splitDirs("esenee"))
	require.Equal(t, []string{"nw", "w", "sw", "e", "e"}, splitDirs("nwwswee"))
}

func TestEx1(t *testing.T) {
	require.Equal(t, 10, part1(example1))
}

func TestEx2(t *testing.T) {
	require.Equal(t, 2208, part2(example1))
}
