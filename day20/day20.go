package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Coord struct {
	Row, Col, Index int
}

type Path []Coord
type PathMap map[int]Coord
type Cheat struct {
	From, To Coord
	Delta    int
}
type StatRec struct {
	Delta, Cnt int
}

func main() {
	fmt.Println("Day 20:")
	part1()
	part2()
}

func part1() {
	lines := loadInput()

	_, pathMap := linesToPath(lines)

	cheats := findCheats(lines, pathMap)

	statsMap := map[int]int{}
	for _, cheat := range cheats {
		statsMap[cheat.Delta] += 1
	}

	stats := []StatRec{}
	for k, v := range statsMap {
		stats = append(stats, StatRec{Delta: k, Cnt: v})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Delta < stats[j].Delta
	})

	sum := 0
	for _, s := range stats {
		// fmt.Println("C=", s.Cnt, "D=", s.Delta)
		if s.Delta >= 100 {
			sum += s.Cnt
		}
	}

	// print(lines, cheats, 4)

	fmt.Println("Part1:", sum)
}

func part2() {
	lines := loadInput()

	path, _ := linesToPath(lines)

	statsMap := map[int]int{}

	cheats := []Cheat{}

	for from := 0; from < len(path)-1; from++ {
		for to := from + 1; to < len(path); to++ {
			dist := abs(path[to].Row-path[from].Row) + abs(path[to].Col-path[from].Col)
			delta := path[to].Index - path[from].Index - dist
			if delta > 0 && dist <= 20 {
				statsMap[delta] += 1
				cheats = append(cheats, Cheat{From: path[from], To: path[to], Delta: delta})
			}
		}
	}

	stats := []StatRec{}
	for k, v := range statsMap {
		stats = append(stats, StatRec{Delta: k, Cnt: v})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Delta < stats[j].Delta
	})

	sum := 0

	for _, s := range stats {
		// fmt.Println("C=", s.Cnt, "D=", s.Delta)
		if s.Delta >= 100 {
			sum += s.Cnt
		}
	}

	// print(lines, cheats, 76)

	fmt.Println("Part2:", sum)
}

func findCheats(lines []string, pathMap PathMap) []Cheat {
	res := []Cheat{}

	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[0])-1; col++ {
			if lines[row][col] == '#' {
				key1 := (row+1)*1000 + col
				cell1, ex1 := pathMap[key1]

				key2 := (row-1)*1000 + col
				cell2, ex2 := pathMap[key2]

				if ex1 && ex2 {
					res = append(res, Cheat{From: cell1, To: cell2, Delta: abs(cell1.Index-cell2.Index) - 2})
				}

				key1 = row*1000 + col + 1
				cell1, ex1 = pathMap[key1]

				key2 = row*1000 + col - 1
				cell2, ex2 = pathMap[key2]

				if ex1 && ex2 {
					res = append(res, Cheat{From: cell1, To: cell2, Delta: abs(cell1.Index-cell2.Index) - 2})
				}
			}
		}
	}

	return res
}

func print(lines []string, cheats []Cheat, delta int) {
	var Blue = "\033[34m"
	var Reset = "\033[0m"

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[0]); col++ {
			isBlue := false
			for _, cheat := range cheats {
				if cheat.Delta == delta && (cheat.From.Row == row && cheat.From.Col == col || cheat.To.Row == row && cheat.To.Col == col) {

					isBlue = true
				}
			}

			if isBlue {
				fmt.Print(Blue + "?" + Reset)
			} else {
				fmt.Print(string(lines[row][col]))
			}
		}
		fmt.Println()
	}
}

func linesToPath(lines []string) (Path, PathMap) {
	res := Path{}
	resMap := PathMap{}

	rowX := -1
	colX := -1

	for row, line := range lines {
		for col, char := range line {
			if char == 'S' {
				rowX = row
				colX = col
			}
		}
	}

	index := 0

	for {
		res = append(res, Coord{Row: rowX, Col: colX, Index: index})
		resMap[rowX*1000+colX] = res[index]

		if lines[rowX][colX] == 'E' {
			break
		}

		if index == 0 {
			rowX, colX = findNext(lines, rowX, colX, -1, -1)
		} else {
			rowX, colX = findNext(lines, rowX, colX, res[index-1].Row, res[index-1].Col)
		}

		index++
	}

	return res, resMap
}

func findNext(lines []string, row, col, prevRow, prevCol int) (int, int) {
	if (lines[row][col+1] == '.' || lines[row][col+1] == 'E') && col+1 != prevCol {
		return row, col + 1
	}
	if (lines[row][col-1] == '.' || lines[row][col-1] == 'E') && col-1 != prevCol {
		return row, col - 1
	}
	if (lines[row+1][col] == '.' || lines[row+1][col] == 'E') && row+1 != prevRow {
		return row + 1, col
	}
	if (lines[row-1][col] == '.' || lines[row-1][col] == 'E') && row-1 != prevRow {
		return row - 1, col
	}
	panic("No Next")
}

func loadInput() []string {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	res := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return res
}

func abs(v int) int {
	if v > 0 {
		return v
	}

	return -v
}
