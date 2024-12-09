package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2024/day01"
	"github.com/umbe77/aoc-2024/day02"
	"github.com/umbe77/aoc-2024/day03"
	"github.com/umbe77/aoc-2024/day04"
	"github.com/umbe77/aoc-2024/day05"
	"github.com/umbe77/aoc-2024/day06"
	"github.com/umbe77/aoc-2024/day07"
	"github.com/umbe77/aoc-2024/day08"
	"github.com/umbe77/aoc-2024/day09"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "01", "day in format dd")
	flag.Parse()

	fmt.Printf("Day %s\n", day)
	switch day {
	case "01":
		day01.Execute()
		break
	case "02":
		day02.Execute()
		break
	case "03":
		day03.Execute()
		break
	case "04":
		day04.Execute()
		break
	case "05":
		day05.Execute()
		break
	case "06":
		day06.Execute()
		break
	case "07":
		day07.Execute()
		break
	case "08":
		day08.Execute()
		break
	case "09":
		day09.Execute()
		break
	}
}
