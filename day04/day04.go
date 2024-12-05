package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 04:")
	part1()
	part2()
}

func part1() {
	lines := loadInput()

	cnt := 0
	for deltaCol := -1; deltaCol <= 1; deltaCol++ {
		for deltaRow := -1; deltaRow <= 1; deltaRow++ {
			if deltaCol != 0 || deltaRow != 0 {
				cnt += checkWord(lines, "XMAS", deltaRow, deltaCol)
			}
		}
	}

	fmt.Println("Part1:", cnt)
}

func part2() {
	lines := loadInput()

	cnt := 0
	for row := 1; row <= len(lines)-2; row++ {
		for col := 1; col <= len(lines[0])-2; col++ {
			if checkPattern1(lines, row, col) {
				cnt++
			}
			if checkPattern2(lines, row, col) {
				cnt++
			}
			if checkPattern3(lines, row, col) {
				cnt++
			}
			if checkPattern4(lines, row, col) {
				cnt++
			}
		}
	}

	fmt.Println("Part2:", cnt)
}

// M.S
// .A.
// M.S
func checkPattern1(lines []string, row, col int) bool {
	if lines[row][col] != 'A' {
		return false
	}
	if lines[row-1][col-1] != 'M' {
		return false
	}
	if lines[row+1][col-1] != 'M' {
		return false
	}
	if lines[row-1][col+1] != 'S' {
		return false
	}
	if lines[row+1][col+1] != 'S' {
		return false
	}
	return true
}

// S.M
// .A.
// S.M
func checkPattern2(lines []string, row, col int) bool {
	if lines[row][col] != 'A' {
		return false
	}
	if lines[row-1][col-1] != 'S' {
		return false
	}
	if lines[row+1][col-1] != 'S' {
		return false
	}
	if lines[row-1][col+1] != 'M' {
		return false
	}
	if lines[row+1][col+1] != 'M' {
		return false
	}
	return true
}

// M.M
// .A.
// S.S
func checkPattern3(lines []string, row, col int) bool {
	if lines[row][col] != 'A' {
		return false
	}
	if lines[row-1][col-1] != 'M' {
		return false
	}
	if lines[row+1][col-1] != 'S' {
		return false
	}
	if lines[row-1][col+1] != 'M' {
		return false
	}
	if lines[row+1][col+1] != 'S' {
		return false
	}
	return true
}

// S.S
// .A.
// M.M
func checkPattern4(lines []string, row, col int) bool {
	if lines[row][col] != 'A' {
		return false
	}
	if lines[row-1][col-1] != 'S' {
		return false
	}
	if lines[row+1][col-1] != 'M' {
		return false
	}
	if lines[row-1][col+1] != 'S' {
		return false
	}
	if lines[row+1][col+1] != 'M' {
		return false
	}
	return true
}

func checkWord(lines []string, word string, deltaRow, deltaCol int) int {
	var (
		fromRow int
		fromCol int
		toRow   int
		toCol   int
	)

	if deltaRow == -1 {
		fromRow = len(word) - 1
		toRow = len(lines) - 1
	} else if deltaRow == 0 {
		fromRow = 0
		toRow = len(lines) - 1
	} else if deltaRow == 1 {
		fromRow = 0
		toRow = len(lines) - len(word)
	}

	if deltaCol == -1 {
		fromCol = len(word) - 1
		toCol = len(lines[0]) - 1
	} else if deltaCol == 0 {
		fromCol = 0
		toCol = len(lines[0]) - 1
	} else if deltaCol == 1 {
		fromCol = 0
		toCol = len(lines[0]) - len(word)
	}

	cnt := 0

	for row := fromRow; row <= toRow; row++ {
		for col := fromCol; col <= toCol; col++ {
			equal := true
			for index, char := range word {
				if byte(char) != lines[row+index*deltaRow][col+index*deltaCol] {
					equal = false
					break
				}
			}

			if equal {
				cnt++
			}
		}
	}

	return cnt
}

func loadInput() []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return res
}
