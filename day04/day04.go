package day04

import (
	"fmt"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {

	rows := make([]string, 0)
	utils.ReadFile("day04/input.txt", func(line string) {
		rows = append(rows, line)
	})

	fmt.Printf("part1: %d\n", part1(rows))
	fmt.Printf("part2: %d\n", part2(rows))
}

type Point struct {
	X, Y int
}

type Segment struct {
	S Point
	E Point
}

func part1(rows []string) int {
	count := 0

	ulimit := len(rows)
	counted := make(map[Segment]bool)

	for y, r := range rows {
		for x, c := range r {
			if c == 'S' || c == 'X' {
				if x+3 < ulimit { //cerco a destra
					if w := r[x : x+4]; w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x, y}, Point{x + 3, y}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
				if x-3 >= 0 { //leggo a sinistra
					if w := r[x-3 : x+1]; w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x - 3, y}, Point{x, y}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
				if y+3 < ulimit { // down
					if w := fmt.Sprintf("%s%s%s%s", string(rows[y][x]), string(rows[y+1][x]), string(rows[y+2][x]), string(rows[y+3][x])); w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x, y}, Point{x, y + 3}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
				if y-3 >= 0 { //up
					if w := fmt.Sprintf("%s%s%s%s", string(rows[y-3][x]), string(rows[y-2][x]), string(rows[y-1][x]), string(rows[y][x])); w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x, y - 3}, Point{x, y}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
				if x+3 < ulimit && y+3 < ulimit { //down right
					if w := fmt.Sprintf("%s%s%s%s", string(rows[y][x]), string(rows[y+1][x+1]), string(rows[y+2][x+2]), string(rows[y+3][x+3])); w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x, y}, Point{x + 3, y + 3}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
				if x-3 >= 0 && y+3 < ulimit { //down left
					if w := fmt.Sprintf("%s%s%s%s", string(rows[y][x]), string(rows[y+1][x-1]), string(rows[y+2][x-2]), string(rows[y+3][x-3])); w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x - 3, y + 3}, Point{x, y}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
				if x+3 < ulimit && y-3 >= 0 { //up right
					if w := fmt.Sprintf("%s%s%s%s", string(rows[y][x]), string(rows[y-1][x+1]), string(rows[y-2][x+2]), string(rows[y-3][x+3])); w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x, y}, Point{x + 3, y - 3}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
				if x-3 >= 0 && y-3 >= 0 { //up left
					if w := fmt.Sprintf("%s%s%s%s", string(rows[y][x]), string(rows[y-1][x-1]), string(rows[y-2][x-2]), string(rows[y-3][x-3])); w == "XMAS" || w == "SAMX" {
						s := Segment{Point{x - 3, y - 3}, Point{x, y}}
						if _, ok := counted[s]; !ok {
							counted[s] = true
							count = count + 1
						}
					}
				}
			}
		}
	}

	return count
}

func part2(rows []string) int {
	count := 0
	ulimit := len(rows)
	for y, r := range rows {
		for x, c := range r {
			if c == 'A' && x-1 >= 0 && x+1 < ulimit && y-1 >= 0 && y+1 < ulimit {

				if (rows[y-1][x-1] == 'M' &&
					rows[y-1][x+1] == 'M' &&
					rows[y+1][x-1] == 'S' &&
					rows[y+1][x+1] == 'S') ||

					(rows[y-1][x-1] == 'M' &&
						rows[y-1][x+1] == 'S' &&
						rows[y+1][x-1] == 'M' &&
						rows[y+1][x+1] == 'S') ||

					(rows[y-1][x-1] == 'S' &&
						rows[y-1][x+1] == 'M' &&
						rows[y+1][x-1] == 'S' &&
						rows[y+1][x+1] == 'M') ||

					(rows[y-1][x-1] == 'S' &&
						rows[y-1][x+1] == 'S' &&
						rows[y+1][x-1] == 'M' &&
						rows[y+1][x+1] == 'M') {
					count = count + 1

				}

			}
		}
	}
	return count
}
