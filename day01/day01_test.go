package day01

import (
	"testing"
)

var input = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

func TestPart1(t *testing.T) {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, row := range input {
		l, r := splitLine(row)
		left = append(left, l)
		right = append(right, r)
	}
	res := part1(left, right)
	if res != 11 {
		t.Errorf("Expected 11, got: %d", res)
	}
}

func TestPart2(t *testing.T) {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, row := range input {
		l, r := splitLine(row)
		left = append(left, l)
		right = append(right, r)
	}
	res := part2(left, right)
	if res != 31 {
		t.Errorf("Expected 31, got: %d", res)
	}
}
