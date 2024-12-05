package day05

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day05/input.txt", func(line string) {
		input = append(input, line)
	})
	rules := Parse(input)
	fmt.Printf("part1: %d\n", part1(rules))
	fmt.Printf("part2: %d\n", part2(rules))
}

type PageRule struct {
	B int
	E int
}

type Rules struct {
	RB map[int][]int
	RE map[int][]int
	P  []map[int]int
	Pi [][]int
	Po []int
}

func Parse(input []string) Rules {

	rules := Rules{
		RB: make(map[int][]int),
		RE: make(map[int][]int),
		P:  make([]map[int]int, 0),
		Pi: make([][]int, 0),
	}

	readRules := true

	for _, line := range input {
		if line == "" {
			readRules = false
			continue
		}
		if readRules {
			pages := strings.Split(line, "|")
			page0 := utils.Atoi(pages[0])
			page1 := utils.Atoi(pages[1])
			if _, ok := rules.RB[page0]; !ok {
				rules.RB[page0] = make([]int, 0)
			}
			rules.RB[page0] = append(rules.RB[page0], page1)
			if _, ok := rules.RE[page1]; !ok {
				rules.RE[page1] = make([]int, 0)
			}
			rules.RE[page1] = append(rules.RE[page1], page0)
			continue
		}
		prints := strings.Split(line, ",")
		pages := make(map[int]int)
		pindex := make([]int, 0)
		for i, p := range prints {
			page := utils.Atoi(p)
			pages[page] = i
			pindex = append(pindex, page)
		}
		rules.P = append(rules.P, pages)
		rules.Pi = append(rules.Pi, pindex)
	}

	return rules

}

func part1(rules Rules) int {
	sum := 0

	// fmt.Println(rules.RB)
	// fmt.Println(rules.RE)
	// [29:4 47:1 53:3 61:2 75:0]
	for i, p := range rules.P {
		valid := true
		for k, v := range p {
			// fmt.Println(k, v)
			if _, ok := rules.RB[k]; ok {
				for _, b := range rules.RB[k] {
					if e, bOk := p[b]; bOk && e < v {
						// if i == 0 {
						// 	fmt.Println("begin", "p", p, "b", b, "e", e, "v", v, rules.RB[k])
						// }
						valid = false
						break
					}
				}
			}
			// if _, ok := rules.RE[k]; ok {
			// 	for _, e := range rules.RE[k] {
			// 		if b, eOk := p[e]; eOk && b < v {
			// 			if i == 0 {
			// 				fmt.Println("end", i, p, b, e, v)
			// 			}
			// 			valid = false
			// 			break
			// 		}
			// 	}
			// }
		}
		// fmt.Println(p, valid)
		if valid {
			idx := int(math.Floor(float64(len(rules.P[i])) / 2))
			sum = sum + rules.Pi[i][idx]
		}
	}

	return sum
}

func part2(rules Rules) int {
	sum := 0
	notValid := make([][]int, 0)

	for i, p := range rules.P {
		valid := true
		for k, v := range p {
			// fmt.Println(k, v)
			if _, ok := rules.RB[k]; ok {
				for _, b := range rules.RB[k] {
					if e, bOk := p[b]; bOk && e < v {
						// if i == 0 {
						// 	fmt.Println("begin", "p", p, "b", b, "e", e, "v", v, rules.RB[k])
						// }
						valid = false
						break
					}
				}
			}
		}
		if !valid {
			notValid = append(notValid, rules.Pi[i])
		}
	}

	for _, prints := range notValid {
		// fmt.Println(prints)
		slices.SortFunc(prints, func(a, b int) int {
			if _, ok := rules.RB[a]; ok {
				for _, v := range rules.RB[a] {
					if v == b {
						return -1
					}
				}
			}
			if _, ok := rules.RE[a]; ok {
				for _, v := range rules.RE[a] {
					if v == b {
						return 1
					}
				}
			}
			return 0
		})
		
		idx := int(math.Floor(float64(len(prints)) / 2))
		sum = sum + prints[idx]
		// fmt.Println(idx, prints[idx], prints)
		// fmt.Println("---------")

	}

	return sum
}
