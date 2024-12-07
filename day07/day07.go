package day07

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {
	eqs := make([]Equation, 0)
	utils.ReadFile("day07/input.txt", func(line string) {
		eq := Equation{Operands: make([]int, 0)}
		s := strings.Split(line, ":")
		eq.Result = utils.Atoi(s[0])
		for _, v := range strings.Fields(s[1]) {
			eq.Operands = append(eq.Operands, utils.Atoi(v))
		}
		eqs = append(eqs, eq)
	})

	fmt.Printf("part1: %d\n", part1(eqs))
	fmt.Printf("part2: %d\n", part2(eqs))
}

type Equation struct {
	Result   int
	Operands []int
}

func Calculate(operands []int, concat bool) []int {
	var results = make([]int, 0)
	results = append(results, operands[0])
	for i := 1; i < len(operands); i++ {
		currentResults := make([]int, 0)
		for _, r := range results {
			currentResults = append(currentResults, r+operands[i])
			currentResults = append(currentResults, r*operands[i])
			if concat {
				currentResults = append(currentResults, utils.Atoi(fmt.Sprintf("%d%d", r, operands[i])))
			}
		}
		results = currentResults
	}
	return results
}

func part1(eqs []Equation) int {
	sum := 0
	for _, eq := range eqs {
		for _, r := range Calculate(eq.Operands, false) {
			if r == eq.Result {
				sum = sum + r
				break
			}
		}
	}
	return sum
}

func part2(eqs []Equation) int {
	sum := 0
	for _, eq := range eqs {
		for _, r := range Calculate(eq.Operands, true) {
			if r == eq.Result {
				sum = sum + r
				break
			}
		}
	}
	return sum
}
