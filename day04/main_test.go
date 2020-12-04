package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestPart1(t *testing.T) {
	input := example

	result := part1(input)

	require.Equal(t, 2, result)
}

func TestByr(t *testing.T) {
	require.False(t, byr("1900"))
	require.False(t, byr("2040"))
	require.True(t, byr("1920"))
	require.True(t, byr("2002"))
	require.True(t, byr("1994"))
}

func TestIyr(t *testing.T) {
	require.False(t, iyr("1900"))
	require.False(t, iyr("2040"))
	require.True(t, iyr("2010"))
	require.True(t, iyr("2020"))
	require.True(t, iyr("2015"))
}

func TestEyr(t *testing.T) {
	require.False(t, eyr("1900"))
	require.False(t, eyr("2040"))
	require.True(t, eyr("2020"))
	require.True(t, eyr("2030"))
	require.True(t, eyr("2022"))
}

func TestHgt(t *testing.T) {
	require.False(t, hgt("190in"))
	require.False(t, hgt("54cm"))
	require.False(t, hgt("190"))
	require.True(t, hgt("59in"))
	require.True(t, hgt("76in"))
	require.True(t, hgt("60in"))
	require.True(t, hgt("150cm"))
	require.True(t, hgt("193cm"))
}

func TestHcl(t *testing.T) {
	require.True(t, hcl("#123abc"))
	require.False(t, hcl("#123abz"))
	require.False(t, hcl("123abc"))
}

func TestEcl(t *testing.T) {
	require.True(t, ecl("amb"))
	require.True(t, ecl("blu"))
	require.True(t, ecl("brn"))
	require.True(t, ecl("gry"))
	require.True(t, ecl("grn"))
	require.True(t, ecl("hzl"))
	require.True(t, ecl("oth"))
	require.False(t, ecl("red"))
}

func TestPid(t *testing.T) {
	require.True(t, pid("000000001"))
	require.False(t, pid("0123456789"))
}
