package day08

import (
	"testing"
)

var input = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

var input2 = []string{
	"T.........",
	"...T......",
	".T........",
	"..........",
	"..........",
	"..........",
	"..........",
	"..........",
	"..........",
	"..........",
}

func TestPart1(t *testing.T) {
	grid := ParseGrid(input)
	res := part1(grid)
	if res != 14 {
		t.Errorf("Expected 14, Got %d", res)
	}
}

func TestPart2(t *testing.T) {
	grid := ParseGrid(input)
	res := part2(grid)
	if res != 34 {
		t.Errorf("Expected 34, Got %d", res)
	}
}

func TestPart2_1(t *testing.T) {
	grid := ParseGrid(input2)
	res := part2(grid)
	if res != 9 {
		t.Errorf("Expected 9, Got %d", res)
	}
}
