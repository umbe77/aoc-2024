package day02

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {

	var reports = make([]Report, 0)

	utils.ReadFile("day02/input.txt", func(line string) {
		levels := strings.Fields(line)
		report := Report{
			Levels: make([]int, len(levels)),
		}
		for i, l := range levels {
			report.Levels[i] = utils.Atoi(l)
		}
		reports = append(reports, report)
	})
	fmt.Printf("part1: %d\n", part1(reports))
	fmt.Printf("part2: %d\n", part2(reports))
}

type Report struct {
	Levels []int
}

func isSafeReport(report Report) bool {
	direction := 1
	for i := 1; i < len(report.Levels); i++ {
		var currentDir = report.Levels[i] - report.Levels[i-1]
		if utils.Abs(currentDir) > 3 || currentDir == 0 {
			return false
		}
		if i == 1 {
			direction = currentDir
		}

		if direction*currentDir < 0 {
			return false
		}
	}
	return true
}

func isSafeLevels(levels []int) bool {
	direction := 1
	for i := 1; i < len(levels); i++ {
		var currentDir = levels[i] - levels[i-1]
		if utils.Abs(currentDir) > 3 || currentDir == 0 {
			return false
		}
		if i == 1 {
			direction = currentDir
		}

		if direction*currentDir < 0 {
			return false
		}
	}
	return true
}

func removeNthLevel(level int, levels []int) []int {
	res := make([]int, 0)
	for i, v := range levels {
		if i != level {
			res = append(res, v)
		}
	}
	return res
}

func part1(reports []Report) int {
	reportsSafe := 0
	for _, r := range reports {
		if isSafeReport(r) {
			reportsSafe = reportsSafe + 1
		}
	}
	return reportsSafe
}

func part2(reports []Report) int {
	reportsSafe := 0
	for _, r := range reports {
		isSafe := false
		for i := range r.Levels {
			if i == 0 && isSafeLevels(r.Levels) {
				isSafe = true
				break
			}
			levels := removeNthLevel(i, r.Levels)
			if isSafeLevels(levels) {
				isSafe = true
				break
			}
		}
		if isSafe {
			reportsSafe = reportsSafe + 1
		}
	}
	return reportsSafe
}
