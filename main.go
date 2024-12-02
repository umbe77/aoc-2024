package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2024/day01"
	"github.com/umbe77/aoc-2024/day02"
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
	}
}
