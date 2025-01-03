package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	Row, Col int
}

func main() {
	part1()
	part2()
}

func part2() {
	max, _, coords := loadInput()

	left := 1
	right := len(coords)

	for {
		if right-left == 1 {
			break
		}

		middle := (left + right) / 2
		ex := checkFirstN(middle, max, coords)

		if ex {
			left = middle
		} else {
			right = middle
		}

		fmt.Println(left, middle, right)
	}

	fmt.Printf("Part2: %d %d  %d,%d\n", left, right, coords[right-1].Col, coords[right-1].Row)
}

func checkFirstN(first, max int, coords []Coord) bool {
	corrupted := map[int]bool{}
	for i := 0; i < first; i++ {
		corrupted[k(coords[i].Row, coords[i].Col)] = true
	}

	visited := map[int]int{}
	visited[k(0, 0)] = 0

	findPath(visited, corrupted, 0, 0, max, 0)

	_, ex := visited[k(max, max)]
	return ex
}

func part1() {
	max, first, coords := loadInput()

	corrupted := map[int]bool{}
	for i := 0; i < first; i++ {
		corrupted[k(coords[i].Row, coords[i].Col)] = true
	}

	visited := map[int]int{}
	visited[k(0, 0)] = 0

	findPath(visited, corrupted, 0, 0, max, 0)

	fmt.Println("Part1:", visited[k(max, max)])
}

func findPath(visited map[int]int, corrupted map[int]bool, row, col, max, current int) {
	nexts := getNext(corrupted, row, col, max)

	// fmt.Println(nexts)
	// fmt.Scanln()

	for _, next := range nexts {
		key := k(next.Row, next.Col)
		score, ex := visited[key]

		if !ex || score > current+1 {
			visited[key] = current + 1
			findPath(visited, corrupted, next.Row, next.Col, max, current+1)
		}
	}
}

func getNext(corrupted map[int]bool, row, col, max int) []Coord {
	res := []Coord{}

	if row+1 <= max && !corrupted[k(row+1, col)] {
		res = append(res, Coord{Row: row + 1, Col: col})
	}

	if row-1 >= 0 && !corrupted[k(row-1, col)] {
		res = append(res, Coord{Row: row - 1, Col: col})
	}

	if col+1 <= max && !corrupted[k(row, col+1)] {
		res = append(res, Coord{Row: row, Col: col + 1})
	}

	if col-1 >= 0 && !corrupted[k(row, col-1)] {
		res = append(res, Coord{Row: row, Col: col - 1})
	}

	return res
}

func k(row, col int) int {
	return row*100 + col
}

func loadInput() (int, int, []Coord) {
	max := 70
	first := 1024
	file, _ := os.Open("input.txt")

	// max := 6
	// first := 12
	// file, _ := os.Open("input-0.txt")

	scanner := bufio.NewScanner(file)

	res := []Coord{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")

		row, _ := strconv.Atoi(parts[1])
		col, _ := strconv.Atoi(parts[0])

		res = append(res, Coord{Row: row, Col: col})
	}

	return max, first, res
}
