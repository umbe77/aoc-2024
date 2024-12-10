package day10

import (
	"fmt"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {

	input := make([]string, 0)
	// utils.ReadFile("day10/input_test_1.txt", func(line string) {
	utils.ReadFile("day10/input.txt", func(line string) {
		input = append(input, line)
	})

	grid := ParseGrid(input)

	fmt.Printf("part1: %d\n", part1(grid))
	fmt.Printf("part2: %d\n", part2(grid))
}

func ParseGrid(input []string) Grid {
	grid := Grid{
		Points:       make(map[Point]int),
		TailingHeads: make([]Point, 0),
	}

	for y, row := range input {
		for x, v := range row {
			pointValue := utils.Atoi(string(v))
			grid.Points[Point{x, y}] = pointValue
			if pointValue == 0 {
				grid.TailingHeads = append(grid.TailingHeads, Point{x, y})
			}
		}
	}

	return grid
}

type Grid struct {
	Points       map[Point]int
	TailingHeads []Point
}

type Point struct {
	X, Y int
}

func part1(grid Grid) int {
	totalScore := 0
	for _, p := range grid.TailingHeads {
		score := make(map[Point]bool)

		stack := utils.New[Point]()
		stack.Push(p)
		for !stack.IsEmpty() {
			currentPoint := stack.Pop()
			currentValue := grid.Points[currentPoint]
			if currentValue == 9 {
				// fmt.Println(currentValue, currentPoint)
				score[currentPoint] = true
				continue
			}
			for i := range 4 {
				var checkPoint Point
				switch i {
				case 0:
					checkPoint = Point{currentPoint.X, currentPoint.Y - 1}
					break
				case 1:
					checkPoint = Point{currentPoint.X + 1, currentPoint.Y}
					break
				case 2:
					checkPoint = Point{currentPoint.X, currentPoint.Y + 1}
					break
				case 3:
					checkPoint = Point{currentPoint.X - 1, currentPoint.Y}
					break
				}
				// fmt.Println("tocheck", currentValue, currentPoint, grid.Points[checkPoint], checkPoint)
				if checkValue, ok := grid.Points[checkPoint]; ok && checkValue == currentValue+1 {
					stack.Push(checkPoint)
					// fmt.Println(checkPoint, currentPoint, currentValue, stack.Len(), stack)
				}
			}
		}

		totalScore = totalScore + len(score)
	}
	return totalScore
}

func part2(grid Grid) int {
	totalScore := 0
	for _, p := range grid.TailingHeads {
		score := 0

		stack := utils.New[Point]()
		stack.Push(p)
		for !stack.IsEmpty() {
			currentPoint := stack.Pop()
			currentValue := grid.Points[currentPoint]
			if currentValue == 9 {
				// fmt.Println(currentValue, currentPoint)
				score = score + 1
				continue
			}
			for i := range 4 {
				var checkPoint Point
				switch i {
				case 0:
					checkPoint = Point{currentPoint.X, currentPoint.Y - 1}
					break
				case 1:
					checkPoint = Point{currentPoint.X + 1, currentPoint.Y}
					break
				case 2:
					checkPoint = Point{currentPoint.X, currentPoint.Y + 1}
					break
				case 3:
					checkPoint = Point{currentPoint.X - 1, currentPoint.Y}
					break
				}
				// fmt.Println("tocheck", currentValue, currentPoint, grid.Points[checkPoint], checkPoint)
				if checkValue, ok := grid.Points[checkPoint]; ok && checkValue == currentValue+1 {
					stack.Push(checkPoint)
					// fmt.Println(checkPoint, currentPoint, currentValue, stack.Len(), stack)
				}
			}
		}

		totalScore = totalScore + score
	}
	return totalScore
}
