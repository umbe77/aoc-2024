package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2024/utils"
)

const input_data = "41078 18 7 0 4785508 535256 8154 447"
const input_test = "125 17"

func Execute() {

	// fmt.Printf("part1: %d\n", part1(input_test))
	fmt.Printf("part1: %d\n", part1(input_data))
	fmt.Printf("part2: %d\n", part2(input_test))
	// fmt.Printf("part2: %d\n", part2(input_data))
}

func part1(input string) int {
	stones := strings.Fields(input)
	// fmt.Println(0, strings.Join(stones, " "))
	for range 25 {
		newStone := make([]string, 0)

		for _, s := range stones {
			if s == "0" {
				newStone = append(newStone, "1")
			} else if len(s)%2 == 0 {
				left := s[:(len(s) / 2)]
				rigth := s[len(s)/2:]
				newStone = append(newStone, left)
				newStone = append(newStone, "0")
				if utils.Atoi(rigth) != 0 {
					newStone[len(newStone)-1] = strconv.Itoa(utils.Atoi(rigth))
				}
			} else {
				newStone = append(newStone, strconv.Itoa(utils.Atoi(s)*2024))
			}
		}

		// fmt.Println(i+1, strings.Join(newStone, " "))
		stones = newStone
	}
	return len(stones)
}

func part2(input string) int {
	ss := strings.Fields(input)
	stones := []string{
		ss[0],
	}
	fmt.Println(0, strings.Join(stones, " "))
	for i := range 75 {
		newStone := make([]string, 0)

		for _, s := range stones {
			if s == "0" {
				newStone = append(newStone, "1")
			} else if len(s)%2 == 0 {
				left := s[:(len(s) / 2)]
				rigth := s[len(s)/2:]
				newStone = append(newStone, left)
				newStone = append(newStone, "0")
				if utils.Atoi(rigth) != 0 {
					newStone[len(newStone)-1] = strconv.Itoa(utils.Atoi(rigth))
				}
			} else {
				newStone = append(newStone, strconv.Itoa(utils.Atoi(s)*2024))
			}
		}

		fmt.Println(i+1, strings.Join(newStone, " "))
		stones = newStone
	}
	return len(stones)
}
