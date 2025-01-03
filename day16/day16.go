package main

import (
	"bufio"
	"fmt"
	"os"
)

type Move struct {
	Row, Col int
	Dir      rune
	Addition int
}

func main() {
	fmt.Println("Day 16:")
	part1()
	fmt.Println()
	fmt.Println()
	// part2()
}

func part1() {
	field := loadInput()

	dir := 'E'
	row, col := findRune(field, 'S')

	best := -1
	visited := map[int]int{}

	path := map[int]bool{}
	path[k2(row, col)] = true

	allVisited := map[int]bool{}

	findPath(field, visited, dir, row, col, 0, &best, path, allVisited)

	// printVis(field, allVisited)

	fmt.Println("Part12:", best, len(allVisited)+1)
}

func findPath(field [][]rune, visited map[int]int, dir rune, row, col, current int, best *int, path map[int]bool, allVisited map[int]bool) {
	moves := getMoves(field, visited, path, dir, row, col, current)

	// if row == 7 && col == 15 {
	// 	fmt.Println(string(dir), current)
	// 	fmt.Println("N:", visited[k(row, col, 'N')])
	// 	fmt.Println("S:", visited[k(row, col, 'S')])

	// 	for _, m := range moves {
	// 		fmt.Printf("%+v\n", m)
	// 	}
	// 	printVis(field, path)
	// 	fmt.Scanln()
	// }

	for _, move := range moves {
		if field[move.Row][move.Col] == 'E' {
			if *best == -1 || *best >= current+move.Addition {
				// fmt.Println("Solution:", current+move.Addition)
				// print(field, visited)
				// path[k2(move.Row, move.Col)] = true

				*best = current + move.Addition

				if *best == 108504 {
					for k, v := range path {
						if v {
							allVisited[k] = true
						}
					}
					// fmt.Println("=", allVisited)
					// printVis(field, allVisited)
				}
			}
		} else {
			key := k(move.Row, move.Col, move.Dir)
			visited[key] = current + move.Addition

			key2 := k2(move.Row, move.Col)
			path[key2] = true
			findPath(field, visited, move.Dir, move.Row, move.Col, current+move.Addition, best, path, allVisited)
			path[key2] = false
		}
	}
}

var deltas = []Move{
	{Row: 0, Col: 1, Dir: 'E'},
	{Row: 0, Col: -1, Dir: 'W'},
	{Row: -1, Col: 0, Dir: 'N'},
	{Row: 1, Col: 0, Dir: 'S'},
}

func getMoves(field [][]rune, visited map[int]int, path map[int]bool, dir rune, row, col, current int) []Move {
	res := []Move{}

	for _, delta := range deltas {
		key2 := k2(row+delta.Row, col+delta.Col)
		wasHere := path[key2]

		if !wasHere && notOpposite(dir, delta.Dir) && field[row+delta.Row][col+delta.Col] != '#' {
			key := k(row+delta.Row, col+delta.Col, delta.Dir)

			addition := 1
			if dir != delta.Dir {
				addition += 1000
			}

			score, ex := visited[key]
			if !ex || score >= current+addition {
				res = append(res, Move{Row: row + delta.Row, Col: col + delta.Col, Dir: delta.Dir, Addition: addition})
			}
		}
	}

	return res
}

func notOpposite(dir1, dir2 rune) bool {
	return dir1 == 'N' && dir2 != 'S' || dir1 == 'S' && dir2 != 'N' || dir1 == 'E' && dir2 != 'W' || dir1 == 'W' && dir2 != 'E'
}

func k(r, c int, d rune) int {
	return r*1_000_000 + c*1_000 + int(d)
}

func k2(r, c int) int {
	return r*1_000 + c
}

var Red = "\033[91m"
var Reset = "\033[0m"

func print(field [][]rune) {
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[0]); col++ {
			fmt.Print(string(field[row][col]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func printVis(field [][]rune, visited map[int]bool) {
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[0]); col++ {
			if visited[k2(row, col)] {
				fmt.Print(Red + "+" + Reset)
			} else {
				fmt.Print(string(field[row][col]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func findRune(field [][]rune, char rune) (int, int) {
	for row := 1; row < len(field)-1; row++ {
		for col := 1; col < len(field[0])-1; col++ {
			if field[row][col] == char {
				return row, col
			}
		}
	}
	return -1, -1
}

func loadInput() [][]rune {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	res := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		res = append(res, []rune(line))
	}

	return res
}
