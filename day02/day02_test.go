package day02

import (
	"strings"
	"testing"

	"github.com/umbe77/aoc-2024/utils"
)

var input = []string{
	"7 6 4 2 1",
	"1 2 7 8 9",
	"9 7 6 2 1",
	"1 3 2 4 5",
	"8 6 4 4 1",
	"1 3 6 7 9",
}

func TestPart1(t *testing.T) {
	
	var reports = make([]Report, 0)

	for _, line := range input {
		levels := strings.Fields(line)
		report := Report{
			Levels: make([]int, len(levels)),
		}
		for i, l := range levels {
			report.Levels[i] = utils.Atoi(l)
		}
		reports = append(reports, report)
	}

	res := part1(reports)
	if res != 2 {
		t.Errorf("Expected 2, Got %d", res)
	}
}

func TestPart2(t *testing.T) {
	
	var reports = make([]Report, 0)

	for _, line := range input {
		levels := strings.Fields(line)
		report := Report{
			Levels: make([]int, len(levels)),
		}
		for i, l := range levels {
			report.Levels[i] = utils.Atoi(l)
		}
		reports = append(reports, report)
	}

	res := part2(reports)
	if res != 4 {
		t.Errorf("Expected 4, Got %d", res)
	}
}
