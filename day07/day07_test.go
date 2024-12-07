package day07

import (
	"strings"
	"testing"

	"github.com/umbe77/aoc-2024/utils"
)

var input = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func TestPart1(t *testing.T) {
	eqs := make([]Equation, 0)
	for _, line := range input {
		eq := Equation{Operands: make([]int, 0)}
		s := strings.Split(line, ":")
		eq.Result = utils.Atoi(s[0])
		for _, v := range strings.Fields(s[1]) {
			eq.Operands = append(eq.Operands, utils.Atoi(v))
		}
		eqs = append(eqs, eq)
	}

	res := part1(eqs)
	if res != 3749 {
		t.Errorf("Expected 3749, Got %d", res)
	}
}

func TestPart2(t *testing.T) {
	eqs := make([]Equation, 0)
	for _, line := range input {
		eq := Equation{Operands: make([]int, 0)}
		s := strings.Split(line, ":")
		eq.Result = utils.Atoi(s[0])
		for _, v := range strings.Fields(s[1]) {
			eq.Operands = append(eq.Operands, utils.Atoi(v))
		}
		eqs = append(eqs, eq)
	}

	res := part2(eqs)
	if res != 11387 {
		t.Errorf("Expected 11387, Got %d", res)
	}
}
