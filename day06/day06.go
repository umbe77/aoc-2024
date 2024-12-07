package day06

import (
	"fmt"
	"maps"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {
	labMap := LabMap{
		Obstacles: make([]Point, 0),
	}

	lineNum := 0
	utils.ReadFile("day06/input.txt", func(line string) {
		for i, c := range line {
			if c == '#' {
				labMap.Obstacles = append(labMap.Obstacles, Point{i, lineNum})
			}
			if c == '^' {
				labMap.GuardInit = Point{i, lineNum}
			}
		}
		lineNum = lineNum + 1
	})
	labMap.Columns = lineNum
	labMap.Rows = lineNum

	fmt.Printf("part1: %d\n", part1(labMap))
	fmt.Printf("part2: %d\n", part2(labMap))
}

type Point struct {
	X int
	Y int
}

func Advance(p Point, dir int) Point {
	switch dir {
	case 0:
		return Point{X: p.X, Y: p.Y - 1}
	case 1:
		return Point{X: p.X + 1, Y: p.Y}
	case 2:
		return Point{X: p.X, Y: p.Y + 1}
	case 3:
		return Point{X: p.X - 1, Y: p.Y}
	}
	return p
}

type LabMap struct {
	Obstacles []Point
	GuardInit Point
	Rows      int
	Columns   int
}

func part1(labMap LabMap) int {
	count := 1
	direction := 0
	currentPos := labMap.GuardInit
	visited := make(map[Point]bool)
	visited[currentPos] = true
	for {
		newPoint := Advance(currentPos, (4+direction)%4) //uso l'addizione modulare per tenere traccia della direzione
		if newPoint.X < 0 || newPoint.X >= labMap.Columns || newPoint.Y < 0 || newPoint.Y >= labMap.Rows {
			break
		}
		isObstacle := false
		for _, o := range labMap.Obstacles {
			if newPoint == o {
				isObstacle = true
				break
			}
		}
		if isObstacle {
			direction = direction + 1
			continue
		}
		// devo contare i punti in maniera univoca
		currentPos = newPoint
		if _, ok := visited[currentPos]; !ok {
			count = count + 1
			visited[currentPos] = true

		}

	}
	return count
}

func part2(labMap LabMap) int {
	// colpisco 2 volte il nuovo ostacolo: sono in un loop
	direction := 0
	currentPos := labMap.GuardInit
	visited := make(map[Point]bool)
	for {
		newPoint := Advance(currentPos, (4+direction)%4) //uso l'addizione modulare per tenere traccia della direzione
		if newPoint.X < 0 || newPoint.X >= labMap.Columns || newPoint.Y < 0 || newPoint.Y >= labMap.Rows {
			break
		}
		isObstacle := false
		for _, o := range labMap.Obstacles {
			if newPoint == o {
				isObstacle = true
				break
			}
		}
		if isObstacle {
			direction = direction + 1
			continue
		}
		// devo contare i punti in maniera univoca
		currentPos = newPoint
		if _, ok := visited[currentPos]; !ok {
			visited[currentPos] = true

		}

	}

	obstaclesMap := make(map[Point]bool)
	for _, v := range labMap.Obstacles {
		obstaclesMap[v] = true
	}

	// fmt.Println(visited)
	// fmt.Println(obstaclesMap)
	// fmt.Println("-----")

	countStuck := 0
	for k := range visited {
		tDirection := 0
		obstacles := maps.Clone(obstaclesMap)
		obstacles[k] = true

		tCurrentPos := labMap.GuardInit
		newObstaclesHits := 0
		tVisited := make(map[Point]bool)
		isStuckInLoop := false
		for {
			newPoint := Advance(tCurrentPos, (4+tDirection)%4) //uso l'addizione modulare per tenere traccia della direzione
			if newPoint == k {
				newObstaclesHits = newObstaclesHits + 1
			}
			if newObstaclesHits >= 2 {
				isStuckInLoop = true
				break
			}
			if newPoint.X < 0 || newPoint.X >= labMap.Columns || newPoint.Y < 0 || newPoint.Y >= labMap.Rows {
				break
			}
			isObstacle := false
			if _, ok := obstacles[newPoint]; ok {
				isObstacle = true
			}
			if isObstacle {
				tDirection = tDirection + 1
				continue
			}
			// devo contare i punti in maniera univoca
			tCurrentPos = newPoint
			if _, ok := tVisited[tCurrentPos]; !ok {
				tVisited[tCurrentPos] = true
			}
		}
		if isStuckInLoop {
			countStuck = countStuck + 1
		}
	}

	return countStuck
}
