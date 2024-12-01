package day01

import (
	"fmt"
	"sort"
	"strings"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {
	left := make([]int, 0)
	right := make([]int, 0)

	utils.ReadFile("day01/input.txt", func(line string) {
		l, r := splitLine(line)
		left = append(left, l)
		right = append(right, r)
	})
	fmt.Printf("part1: %d\n", part1(left, right))
	fmt.Printf("part2: %d\n", part2(left, right))
}

func splitLine(line string) (left, right int) {
	lineChrs := strings.Fields(line)
	return utils.Atoi(lineChrs[0]), utils.Atoi(lineChrs[1])
}

func part1(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		sum = sum + utils.Abs(left[i] - right[i])
	}
	return sum
}

func part2(left, right []int) int {
	res := 0

	uniques := make(map[int]int)

	for _, r := range right {
		if _, ok := uniques[r]; !ok {
			uniques[r] = 0
		}
		uniques[r] = uniques[r] + 1
	}

	for _, l := range left{
		val := 0
		ok := false
		if val, ok = uniques[l]; !ok {
			val = 0
		}
		res = res + (l * val)
	}

	return res
}
