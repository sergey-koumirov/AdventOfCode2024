package main

import (
	"bufio"
	"fmt"
	"os"
)

type RC struct {
	Row, Col int
}

func main() {
	fmt.Println("Day 10:")
	part1()
	fmt.Println()
	part2()
}

func part1() {
	lines, maxRow, maxCol := loadInput()

	cnt := 0

	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if lines[row][col] == '0' {
				cnt += find9(lines, row, col)
			}
		}
	}

	fmt.Println("Part1:", cnt)
}

func part2() {
	lines, maxRow, maxCol := loadInput()

	cnt := 0

	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if lines[row][col] == '0' {
				cnt += find9Part2(lines, row, col, maxRow, maxCol)
			}
		}
	}

	fmt.Println("Part2:", cnt)
}

func find9Part2(lines []string, row, col, maxRow, maxCol int) int {
	cnt := 0
	deep9(lines, row, col, maxRow, maxCol, byte('1'), &cnt)
	return cnt
}

func deep9(lines []string, row, col, maxRow, maxCol int, nextChar byte, cnt *int) {
	if col-1 >= 0 && lines[row][col-1] == nextChar {
		if nextChar == byte('9') {
			*cnt++
		} else {
			deep9(lines, row, col-1, maxRow, maxCol, nextChar+1, cnt)
		}
	}
	if row-1 >= 0 && lines[row-1][col] == nextChar {
		if nextChar == byte('9') {
			*cnt++
		} else {
			deep9(lines, row-1, col, maxRow, maxCol, nextChar+1, cnt)
		}
	}
	if col+1 <= maxCol && lines[row][col+1] == nextChar {
		if nextChar == byte('9') {
			*cnt++
		} else {
			deep9(lines, row, col+1, maxRow, maxCol, nextChar+1, cnt)
		}
	}
	if row+1 <= maxRow && lines[row+1][col] == nextChar {
		if nextChar == byte('9') {
			*cnt++
		} else {
			deep9(lines, row+1, col, maxRow, maxCol, nextChar+1, cnt)
		}
	}
}

func find9(lines []string, row, col int) int {
	buffer := []RC{{Row: row, Col: col}}

	for i := 1; i <= 9; i++ {
		newBuffer := []RC{}

		for _, rc := range buffer {
			if rc.Col-1 >= 0 && lines[rc.Row][rc.Col-1] == byte(int('0')+i) {
				newBuffer = append(newBuffer, RC{Row: rc.Row, Col: rc.Col - 1})
			}
			if rc.Row-1 >= 0 && lines[rc.Row-1][rc.Col] == byte(int('0')+i) {
				newBuffer = append(newBuffer, RC{Row: rc.Row - 1, Col: rc.Col})
			}
			if rc.Col+1 < len(lines[0]) && lines[rc.Row][rc.Col+1] == byte(int('0')+i) {
				newBuffer = append(newBuffer, RC{Row: rc.Row, Col: rc.Col + 1})
			}
			if rc.Row+1 < len(lines) && lines[rc.Row+1][rc.Col] == byte(int('0')+i) {
				newBuffer = append(newBuffer, RC{Row: rc.Row + 1, Col: rc.Col})
			}
		}

		buffer = newBuffer
	}

	uniq := map[int]bool{}

	for _, v := range buffer {
		uniq[v.Row*100+v.Col] = true
	}

	return len(uniq)
}

func loadInput() ([]string, int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	maxRow := len(res) - 1
	maxCol := len(res[0]) - 1

	return res, maxRow, maxCol
}
