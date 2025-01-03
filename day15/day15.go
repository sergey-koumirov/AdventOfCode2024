package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cell struct {
	Row, Col int
	Char     rune
}

func main() {
	fmt.Println("Day 15:")
	part1()
	part2()
}

func part1() {
	field, maxRow, maxCol, moves := loadInput()

	rRow, rCol := findRune(field, '@', maxRow, maxCol)

	// print(field)

	for _, move := range moves {
		applyMove(field, &rRow, &rCol, move)

		// fmt.Println(string(move))
		// print(field)
	}

	fmt.Println("Part1:", score(field, 'O'))
}

func part2() {
	field, maxRow, maxCol, moves := loadInput()
	field2 := makeWider(field)
	print(field2)

	rRow, rCol := findRune(field2, '@', maxRow*2, maxCol)

	for _, move := range moves {
		applyMovePart2(field2, &rRow, &rCol, move)

		// fmt.Println(string(move))
		// print(field2)
		// fmt.Scanln()
	}
	print(field2)

	fmt.Println("Part2:", score(field2, '['))
}

func applyMovePart2(f [][]rune, rRow, rCol *int, move rune) {
	dlRow, dlCol := moveToDl(move)

	// left/right
	if dlRow == 0 {
		length := 1

		for {
			nextCell := f[*rRow][*rCol+dlCol*length]
			if nextCell == '.' {
				break
			}
			if nextCell == '#' {
				length = 0
				break
			}
			length++
		}

		if length > 0 {
			for index := *rCol + dlCol*length; index != *rCol; index -= dlCol {
				f[*rRow][index] = f[*rRow][index-dlCol]
			}
			f[*rRow][*rCol] = '.'
			*rCol += dlCol
		}
	}

	// up/down
	if dlCol == 0 {
		cells := []Cell{{Row: *rRow, Col: *rCol, Char: '@'}}
		canMove := cellsToMove(f, &cells, dlRow)

		if canMove {
			for _, cell := range cells {
				f[cell.Row][cell.Col] = '.'
			}
			for _, cell := range cells {
				f[cell.Row+dlRow][cell.Col] = cell.Char
			}
			*rRow += dlRow
		}
	}
}

func cellsToMove(f [][]rune, cells *[]Cell, delta int) bool {
	prevCells := make([]Cell, len(*cells))
	copy(prevCells, *cells)

	for {
		newCells := []Cell{}

		for _, cell := range prevCells {
			if f[cell.Row+delta][cell.Col] == '#' {
				return false
			}
			if f[cell.Row+delta][cell.Col] == '[' {
				newCells = append(
					newCells,
					Cell{Row: cell.Row + delta, Col: cell.Col, Char: '['},
					Cell{Row: cell.Row + delta, Col: cell.Col + 1, Char: ']'},
				)
			}
			if f[cell.Row+delta][cell.Col] == ']' {
				newCells = append(
					newCells,
					Cell{Row: cell.Row + delta, Col: cell.Col, Char: ']'},
					Cell{Row: cell.Row + delta, Col: cell.Col - 1, Char: '['},
				)
			}
		}

		if len(newCells) == 0 {
			break
		}

		*cells = append(*cells, newCells...)
		prevCells = newCells
	}

	return true
}

func makeWider(f [][]rune) [][]rune {
	res := [][]rune{}

	for _, row := range f {
		temp := []rune{}
		for _, char := range row {
			if char == '#' {
				temp = append(temp, '#', '#')
			}
			if char == 'O' {
				temp = append(temp, '[', ']')
			}
			if char == '.' {
				temp = append(temp, '.', '.')
			}
			if char == '@' {
				temp = append(temp, '@', '.')
			}
		}

		res = append(res, temp)
	}

	return res
}

func score(field [][]rune, box rune) int {
	res := 0

	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[row]); col++ {
			if field[row][col] == box {
				res += row*100 + col
			}
		}
	}

	return res
}

func print(field [][]rune) {
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[row]); col++ {
			fmt.Print(string(field[row][col]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func applyMove(field [][]rune, rRow, rCol *int, move rune) {
	dlRow, dlCol := moveToDl(move)

	nextCell := field[*rRow+dlRow][*rCol+dlCol]
	if nextCell == '#' {
		return
	}
	if nextCell == '.' {
		field[*rRow+dlRow][*rCol+dlCol] = '@'
		field[*rRow][*rCol] = '.'

		*rRow = *rRow + dlRow
		*rCol = *rCol + dlCol
		return
	}
	if nextCell == 'O' {
		i := 2
		for {
			cell := field[*rRow+dlRow*i][*rCol+dlCol*i]

			if cell == '#' {
				return
			}

			if cell == '.' {
				field[*rRow+dlRow*i][*rCol+dlCol*i] = 'O'
				field[*rRow+dlRow][*rCol+dlCol] = '@'
				field[*rRow][*rCol] = '.'

				*rRow = *rRow + dlRow
				*rCol = *rCol + dlCol
				return
			}

			i++
		}
	}
}

func moveToDl(move rune) (int, int) {
	if move == '>' {
		return 0, 1
	}
	if move == '<' {
		return 0, -1
	}
	if move == 'v' {
		return 1, 0
	}
	if move == '^' {
		return -1, 0
	}
	return 0, 0
}

func findRune(field [][]rune, char rune, maxRow, maxCol int) (int, int) {
	for row := 1; row < maxRow; row++ {
		for col := 1; col < maxCol; col++ {
			if field[row][col] == char {
				return row, col
			}
		}
	}
	return -1, -1
}

func loadInput() ([][]rune, int, int, []rune) {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	res := [][]rune{}

	moves := []rune{}

	readMoves := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readMoves = true
		}

		if !readMoves {
			res = append(res, []rune(line))
		} else {
			moves = append(moves, []rune(line)...)
		}
	}

	return res, len(res), len(res[0]), moves
}
