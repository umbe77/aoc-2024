package day12

import (
	"fmt"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {

	input := make([]string, 0)
	utils.ReadFile("day12/input_test.txt", func(line string) {
		// utils.ReadFile("day12/input.txt", func(line string) {
		input = append(input, line)
	})

	// fmt.Printf("part1: %d\n", part1(Parse(input), getRegionPart1))
	fmt.Printf("part2: %d\n", part1(Parse(input), getRegionPart2))
}

func Parse(input []string) Grid {
	grid := Grid{
		Points: make(map[Point]string),
		Limit:  len(input),
	}
	for y, row := range input {
		for x, c := range row {
			grid.Points[Point{x, y}] = string(c)
		}
	}
	return grid
}

type Grid struct {
	Points map[Point]string
	Limit  int
}
type Point struct {
	X, Y int
}

type Region struct {
	Area      map[Point]bool
	Perimeter int
}

func getRegionPart2(currentPoint Point, grid Grid, visitedPoints map[Point]bool) Region {
	region := Region{
		Area:      make(map[Point]bool),
		Perimeter: 0,
	}

	perimeter := make(map[Point]bool)

	stack := utils.New[Point]()
	stack.Push(currentPoint)
	region.Area[currentPoint] = true

	for !stack.IsEmpty() {
		p := stack.Pop()
		if _, visited := visitedPoints[p]; !visited {
			visitedPoints[p] = true
			for i := range 4 {
				var checkPoint Point
				var side Point
				switch i {
				case 0:
					checkPoint = Point{p.X, p.Y - 1}
					side = checkPoint
					break
				case 1:
					checkPoint = Point{p.X + 1, p.Y}
					side = checkPoint
					break
				case 2:
					checkPoint = Point{p.X, p.Y + 1}
					side = checkPoint
					break
				case 3:
					checkPoint = Point{p.X - 1, p.Y}
					side = checkPoint
					break
				}

				checkValue := ""
				if c, ok := grid.Points[checkPoint]; ok {
					checkValue = c
				}

				if checkValue == grid.Points[currentPoint] {
					region.Area[checkPoint] = true
					region.Perimeter = region.Perimeter - 1
					stack.Push(checkPoint)
				} else {
					perimeter[side] = true
				}

			}
		}
	}

	region.Perimeter = len(perimeter)

	// fmt.Println(stack.Len(), region)

	return region
}

func getRegionPart1(currentPoint Point, grid Grid, visitedPoints map[Point]bool) Region {
	region := Region{
		Area:      make(map[Point]bool),
		Perimeter: 0,
	}

	stack := utils.New[Point]()
	stack.Push(currentPoint)
	region.Area[currentPoint] = true

	for !stack.IsEmpty() {
		p := stack.Pop()
		if _, visited := visitedPoints[p]; !visited {
			visitedPoints[p] = true
			for i := range 4 {
				var checkPoint Point
				switch i {
				case 0:
					checkPoint = Point{p.X, p.Y - 1}
					break
				case 1:
					checkPoint = Point{p.X + 1, p.Y}
					break
				case 2:
					checkPoint = Point{p.X, p.Y + 1}
					break
				case 3:
					checkPoint = Point{p.X - 1, p.Y}
					break
				}
				region.Perimeter = region.Perimeter + 1
				if checkValue, ok := grid.Points[checkPoint]; ok {
					if checkValue == grid.Points[currentPoint] {
						region.Area[checkPoint] = true
						region.Perimeter = region.Perimeter - 1
						stack.Push(checkPoint)
					}
				}
			}
		}
	}

	// fmt.Println(stack.Len(), region)

	return region
}

func part1(grid Grid, getRegion func(currentPoint Point, grid Grid, visitedPoints map[Point]bool) Region) int {
	visitedPoints := make(map[Point]bool)

	// regions := make(map[string]Region)

	price := 0
	for y := range grid.Limit {
		for x := range grid.Limit {
			currentPoint := Point{x, y}
			if _, ok := visitedPoints[currentPoint]; !ok {
				region := getRegion(currentPoint, grid, visitedPoints)
				// fmt.Println(grid.Points[currentPoint], len(region.Area), region.Perimeter)
				// regions[grid.Points[currentPoint]] = region
				price = price + len(region.Area)*region.Perimeter
			}
		}
	}

	return price
}
