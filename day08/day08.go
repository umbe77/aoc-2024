package day08

import (
	"fmt"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {
	lines := make([]string, 0)

	utils.ReadFile("day08/input.txt", func(line string) {
		lines = append(lines, line)
	})

	grid := ParseGrid(lines)

	fmt.Printf("part1: %d\n", part1(grid))
	fmt.Printf("part2: %d\n", part2(grid))
}

type Point struct {
	X, Y int
}

func ParseGrid(lines []string) Grid {
	grid := Grid{
		Antennas: make(map[rune][]Point),
		Rows:     len(lines),
		Cols:     len(lines[0]),
	}
	for r, row := range lines {
		for c, col := range row {
			if col != '.' {
				if _, ok := grid.Antennas[col]; !ok {
					grid.Antennas[col] = make([]Point, 0)
				}
				grid.Antennas[col] = append(grid.Antennas[col], Point{c, r})
			}
		}
	}
	return grid
}

type Grid struct {
	Antennas   map[rune][]Point
	Rows, Cols int
}

func part1(grid Grid) int {
	locations := make([]Point, 0)
	uniqueLocations := make(map[Point]bool)
	for _, v := range grid.Antennas {
		if len(v) > 1 {
			for _, current := range v {
				for _, point := range v {
					if current != point {
						dx := utils.Abs(current.X - point.X)
						dy := utils.Abs(current.Y - point.Y)

						var direction Point
						if current.X > point.X && current.Y < point.Y { // top right
							direction = Point{1, -1}
						} else if current.X > point.X && current.Y > point.Y { // down right
							direction = Point{1, 1}
						} else if current.X < point.X && current.Y > point.Y { // down left
							direction = Point{-1, 1}
						} else if current.X < point.X && current.Y < point.Y { // top left
							direction = Point{-1, -1}
						}

						loc := Point{
							X: current.X + dx*direction.X,
							Y: current.Y + dy*direction.Y,
						}
						if 0 <= loc.X && loc.X < grid.Cols && 0 <= loc.Y && loc.Y < grid.Rows {
							if _, ok := uniqueLocations[loc]; !ok {
								uniqueLocations[loc] = true
								locations = append(locations, loc)
							}
						}
					}
				}
			}
		}
	}
	return len(locations)
}

func part2(grid Grid) int {
	locations := make([]Point, 0)
	uniqueLocations := make(map[Point]bool)
	for _, v := range grid.Antennas {
		if len(v) > 1 {
			for _, current := range v {
				for _, point := range v {
					if current != point {
						dx := utils.Abs(current.X - point.X)
						dy := utils.Abs(current.Y - point.Y)

						var direction Point
						if current.X > point.X && current.Y < point.Y { // top right
							direction = Point{1, -1}
						} else if current.X > point.X && current.Y > point.Y { // down right
							direction = Point{1, 1}
						} else if current.X < point.X && current.Y > point.Y { // down left
							direction = Point{-1, 1}
						} else if current.X < point.X && current.Y < point.Y { // top left
							direction = Point{-1, -1}
						}

						loc := Point{current.X, current.Y}
						nearesrPoint := make([]Point, 0)
						i := 0
						for {
							loc = Point{
								X: loc.X + dx*direction.X,
								Y: loc.Y + dy*direction.Y,
							}
							if 0 > loc.X || loc.X >= grid.Cols || 0 > loc.Y || loc.Y >= grid.Rows {
								break
							}
							if _, ok := uniqueLocations[loc]; !ok {
								uniqueLocations[loc] = true
								locations = append(locations, loc)
								if i <= 1 {
									nearesrPoint = append(nearesrPoint, loc)
								}
								i++
							}
						}
						if len(nearesrPoint) == 1 {
							if _, ok := uniqueLocations[current]; !ok {
								uniqueLocations[current] = true
								locations = append(locations, current)
							}
						}
						if len(nearesrPoint) == 2 {
							if _, ok := uniqueLocations[point]; !ok {
								uniqueLocations[point] = true
							locations = append(locations, point)
							}
						}
					}
				}
			}
		}
	}

	return len(locations)
}
