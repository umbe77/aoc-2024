package day03

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {

	lines := make([][]Mul, 0)
	utils.ReadFile("day03/input.txt", func(line string) {
		l := ScanInput(line)
		lines = append(lines, l)
	})

	input := utils.ReadAllFile("day03/input.txt")
	lines2 := ScanInput2(input)

	fmt.Printf("part1: %d\n", part1(lines))
	fmt.Printf("part2: %d\n", part2(lines2))
}

type Mul struct {
	Op1 int
	Op2 int
}

func ReadDo(r *bufio.Reader) (bool, bool) {
	enable := false
	ok := false

	for {
		m := utils.ReadStringAndQuote(r)
		if m == "do" {
			c, _, err := r.ReadRune()
			if err != nil {
				break
			}
			if c == '(' {
				c, _, err = r.ReadRune()
				if err != nil {
					break
				}
				if c == ')' {
					enable = true
					ok = true
					break
				}
			}
		}
		if m == "don't" {
			c, _, err := r.ReadRune()
			if err != nil {
				break
			}
			if c == '(' {
				c, _, err = r.ReadRune()
				if err != nil {
					break
				}
				if c == ')' {
					enable = false
					ok = true
					break
				}
			}
		}
		break
	}
	return enable, ok
}

func ReadMul(r *bufio.Reader) (Mul, bool) {
	var (
		mul Mul
		ok  bool
	)
	ok = false
	for {
		m := utils.ReadString(r)
		if m == "mul" {
			c, _, err := r.ReadRune()
			if err != nil {
				break
			}
			if c == '(' {
				c, _, err = r.ReadRune()
				if err != nil {
					break
				}
				var op1, op2 int
				if utils.IsDigit(c) {
					op1 = utils.ReadInt(r)
				}
				c, _, err = r.ReadRune()
				if err != nil {
					break
				}
				if c == ',' {
					c, _, err = r.ReadRune()
					if err != nil {
						break
					}
					if utils.IsDigit(c) {
						op2 = utils.ReadInt(r)
					}
					c, _, err = r.ReadRune()
					if err != nil {
						break
					}
					if c == ')' {
						mul = Mul{op1, op2}
						ok = true
						break
					}
				}
			}
		}
		break
	}
	return mul, ok
}

func ScanInput(input string) []Mul {

	result := make([]Mul, 0)
	r := bufio.NewReader(strings.NewReader(input))

	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		if c == 'm' {
			if mul, ok := ReadMul(r); ok {
				result = append(result, mul)
			}

		}
	}
	return result
}

func ScanInput2(input string) []Mul {

	result := make([]Mul, 0)
	r := bufio.NewReader(strings.NewReader(input))

	enable := true

	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		if c == 'd' {
			if enbl, ok := ReadDo(r); ok {
				enable = enbl
			}
		}
		if c == 'm' && enable {
			if mul, ok := ReadMul(r); ok {
				result = append(result, mul)
			}

		}
	}
	return result
}

func part1(lines [][]Mul) int {
	sum := 0
	for _, l := range lines {
		for _, m := range l {
			sum = sum + (m.Op1 * m.Op2)
		}
	}

	return sum
}

func part2(lines []Mul) int {
	sum := 0
	for _, m := range lines {
		sum = sum + (m.Op1 * m.Op2)
	}

	return sum
}
