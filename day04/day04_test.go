package day04

import (
	"testing"
)

var input = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestPart1(t *testing.T) {
	count := part1(input)
	if count != 18 {
		t.Errorf("Expected 18, Got %d", count)
	}
}

func TestPart2(t *testing.T) {
	count := part2(input)
	if count != 9 {
		t.Errorf("Expected 9, Got %d", count)
	}
}
