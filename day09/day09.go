package day09

import (
	"fmt"
	"slices"

	"github.com/umbe77/aoc-2024/utils"
)

func Execute() {

	var diskMap string
	utils.ReadFile("day09/input.txt", func(line string) {
		diskMap = line
	})

	fmt.Printf("part1: %d\n", part1(diskMap))
	fmt.Printf("part2: %d\n", part2(diskMap))
	// fmt.Printf("part1: %d\n", part1("2333133121414131402"))
	// fmt.Printf("part2: %d\n", part2("2333133121414131402"))
}

func findNextEmptyBlock(i int, defragStorage []int) (int, int) {
	for {
		if i >= len(defragStorage) {
			return -1, -1
		}
		if defragStorage[i] == -1 {
			capacity := 0
			for j := i; j < len(defragStorage) && defragStorage[j] == -1; j++ {
				capacity = capacity + 1
			}
			return i, capacity
		}
		i = i + 1
	}
}
func getStorage(diskMap string) ([]int, int) {
	storage := make([]int, 0)
	fileIndex := 0
	emptyBlockCount := 0
	for i, s := range diskMap {
		file := i%2 == 0
		v := utils.Atoi(string(s))
		sv := -1
		if file {
			sv = fileIndex
			fileIndex = fileIndex + 1
		} else {
			emptyBlockCount = emptyBlockCount + v
		}

		for range v {
			storage = append(storage, sv)
		}

	}
	return storage, emptyBlockCount
}
func part1(diskMap string) int {

	storage, emptyBlockCount := getStorage(diskMap)

	defragStorage := slices.Clone(storage)
	currentEmpyIndex := 0
	for i := len(storage) - 1; i > len(storage)-1-emptyBlockCount; i-- {
		currentBlock := storage[i]
		if currentBlock > -1 {
			currentEmpyIndex, _ = findNextEmptyBlock(currentEmpyIndex, defragStorage)
			if currentEmpyIndex == -1 {
				break
			}
			defragStorage[currentEmpyIndex] = currentBlock
			defragStorage[i] = -1
		}

	}

	checksum := 0

	for i, v := range defragStorage {
		if v == -1 {
			break
		}

		checksum = checksum + i*v

	}

	return checksum
}

func part2(diskMap string) int {
	storage, _ := getStorage(diskMap)
	defragStorage := slices.Clone(storage)

	for i := len(storage) - 1; i >= 0; i-- {
		currentFileSize := 0
		currentBlock := storage[i]
		if currentBlock > -1 {
			sizeBlock := currentBlock

			for sizeBlock == currentBlock {
				currentFileSize = currentFileSize + 1
				i--
				if i < 0 {
					break
				}
				sizeBlock = storage[i]
			}

			i++
			currentEmptyIndex := 0
			currentEmptyCapacity := 0
			for currentEmptyIndex < len(defragStorage) {
				currentEmptyIndex, currentEmptyCapacity = findNextEmptyBlock(currentEmptyIndex, defragStorage)
				if currentEmptyIndex == -1 {
					break
				}
				if i <= currentEmptyIndex {
					break
				}
				if currentFileSize <= currentEmptyCapacity {
					originalI := i + currentFileSize

					for range currentFileSize {
						originalI--
						defragStorage[currentEmptyIndex] = currentBlock
						defragStorage[originalI] = -1
						currentEmptyIndex++
					}
					currentEmptyIndex = currentEmptyIndex + (currentEmptyCapacity - currentFileSize)
					storage = slices.Clone(defragStorage)
					break
				}
				currentEmptyIndex++
			}
		}
	}

	checksum := 0

	for i, v := range defragStorage {
		if v != -1 {
			checksum = checksum + i*v
		}
	}

	return checksum
}
