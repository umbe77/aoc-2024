package day06

import (
	"testing"
)

var input = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

func TestPart1(t *testing.T) {
	labMap := LabMap{
		Obstacles: make([]Point, 0),
	}

	lineNum := 0
	for _, line := range input {
		for i, c := range line {
			if c == '#' {
				labMap.Obstacles = append(labMap.Obstacles, Point{i, lineNum})
			}
			if c == '^' {
				labMap.GuardInit = Point{i, lineNum}
			}
		}
		lineNum = lineNum + 1
	}
	labMap.Columns = lineNum
	labMap.Rows = lineNum

	res := part1(labMap)
	if res != 41 {
		t.Errorf("Expected 41, Got %d", res)
	}
}

func TestPart2(t *testing.T) {
	labMap := LabMap{
		Obstacles: make([]Point, 0),
	}

	lineNum := 0
	for _, line := range input {
		for i, c := range line {
			if c == '#' {
				labMap.Obstacles = append(labMap.Obstacles, Point{i, lineNum})
			}
			if c == '^' {
				labMap.GuardInit = Point{i, lineNum}
			}
		}
		lineNum = lineNum + 1
	}
	labMap.Columns = lineNum
	labMap.Rows = lineNum

	res := part2(labMap)
	if res != 6 {
		t.Errorf("Expected 6, Got %d", res)
	}
}
