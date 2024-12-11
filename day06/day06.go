package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Day 06:")
	part1()
	fmt.Println()
	part2()
}

func part1() {
	lines, row, col := loadInput()

	maxRow := len(lines) - 1
	maxCol := len(lines[0]) - 1

	paintPossible(lines, row, col, maxRow, maxCol)

	cnt := 0
	for _, line := range lines {
		for _, el := range line {
			if el == 2 {
				cnt++
			}
		}
	}

	fmt.Println("Part1:", cnt)
}

func paintPossible(lines [][]int, row, col, maxRow, maxCol int) {
	dir := 'U'
	for {
		nextRow, nextCol := nextRowCol(row, col, dir)

		if nextRow < 0 || nextRow > maxRow || nextCol < 0 || nextCol > maxCol {
			break
		}

		if lines[nextRow][nextCol] == 1 {
			dir = rotateR90(dir)
		} else {
			lines[nextRow][nextCol] = 2
			row = nextRow
			col = nextCol
		}
	}
}

func part2() {
	lines, row, col := loadInput()

	maxRow := len(lines) - 1
	maxCol := len(lines[0]) - 1

	start := time.Now()

	paintPossible(lines, row, col, maxRow, maxCol)
	lines[row][col] = 0

	cnt := 0
	for obsRow := 0; obsRow <= maxRow; obsRow++ {
		for obsCol := 0; obsCol <= maxCol; obsCol++ {
			if lines[obsRow][obsCol] == 2 {
				lines[obsRow][obsCol] = 1
				if isLoop(lines, row, col, maxRow, maxCol) {
					cnt++
				}
				lines[obsRow][obsCol] = 2
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Part2:", cnt, elapsed)
}

func isLoop(lines [][]int, row, col, maxRow, maxCol int) bool {
	path := map[int64]bool{}
	dir := 'U'

	key := int64(row*1_000_000) + int64(col*1_000) + int64(dir)
	path[key] = true

	for {
		nextRow, nextCol := nextRowCol(row, col, dir)

		if nextRow < 0 || nextRow > maxRow || nextCol < 0 || nextCol > maxCol {
			break
		}

		if lines[nextRow][nextCol] == 1 {
			dir = rotateR90(dir)
		} else {
			row = nextRow
			col = nextCol
		}

		key = int64(row*1_000_000) + int64(col*1_000) + int64(dir)
		_, ex := path[key]
		if ex {
			return true
		}
		path[key] = true
	}

	return false
}

func rotateR90(dir rune) rune {
	if dir == 'U' {
		return 'R'
	}
	if dir == 'R' {
		return 'D'
	}
	if dir == 'D' {
		return 'L'
	}
	if dir == 'L' {
		return 'U'
	}
	return dir
}

func nextRowCol(row, col int, dir rune) (int, int) {
	if dir == 'U' {
		return row - 1, col
	}
	if dir == 'R' {
		return row, col + 1
	}
	if dir == 'D' {
		return row + 1, col
	}
	if dir == 'L' {
		return row, col - 1
	}
	return row, col
}

func loadInput() ([][]int, int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := make([][]int, 0)
	row := 0
	col := 0
	lineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		temp := make([]int, len(line))

		for index, char := range line {
			if char == '^' {
				col = index
				row = lineIndex
				temp[index] = 2
			} else if char == '#' {
				temp[index] = 1
			} else {
				temp[index] = 0
			}
		}
		lineIndex += 1
		res = append(res, temp)
	}

	return res, row, col
}
