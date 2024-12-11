package main

import (
	"bufio"
	"fmt"
	"os"
)

type RC struct {
	Row, Col int
}

type Nodes map[rune][]RC

func main() {
	fmt.Println("Day 08:")
	part1()
	fmt.Println()
	part2()
}

func part1() {
	nodes, maxR, maxC := loadInput()

	uniqs := map[int]bool{}
	for _, vv := range nodes {
		// fmt.Println(string(k))
		for i1 := 0; i1 < len(vv)-1; i1++ {
			for i2 := i1 + 1; i2 < len(vv); i2++ {
				// fmt.Println(vv[i1], vv[i2], antiRC(vv[i1], vv[i2]))
				for _, v := range antiRC(vv[i1], vv[i2]) {
					if v.Row >= 0 && v.Row <= maxR && v.Col >= 0 && v.Col <= maxC {
						uniqs[v.Row*100+v.Col] = true
					}
				}
			}
		}
	}

	fmt.Println("Part1:", len(uniqs))
}

func antiRC(rc1, rc2 RC) []RC {
	res := make([]RC, 2)

	deltaR := rc2.Row - rc1.Row
	deltaC := rc2.Col - rc1.Col

	res[0] = RC{Row: rc1.Row - deltaR, Col: rc1.Col - deltaC}
	res[1] = RC{Row: rc2.Row + deltaR, Col: rc2.Col + deltaC}

	return res
}

func part2() {
	nodes, maxR, maxC := loadInput()

	fmt.Println(maxR, maxC)

	uniqs := map[int]bool{}
	for _, vv := range nodes {
		// fmt.Println(string(k))
		for i1 := 0; i1 < len(vv)-1; i1++ {
			for i2 := i1 + 1; i2 < len(vv); i2++ {
				// fmt.Println(vv[i1], vv[i2], antiRC(vv[i1], vv[i2]))
				for _, v := range antiRCPart2(vv[i1], vv[i2], maxR, maxC) {
					if v.Row >= 0 && v.Row <= maxR && v.Col >= 0 && v.Col <= maxC {
						uniqs[v.Row*100+v.Col] = true
					}
				}
			}
		}
	}

	fmt.Println(uniqs)

	fmt.Println("Part2:", len(uniqs))
}

func antiRCPart2(rc1, rc2 RC, maxR, maxC int) []RC {
	res := make([]RC, 0)

	deltaR := rc2.Row - rc1.Row
	deltaC := rc2.Col - rc1.Col

	nextR := rc1.Row
	nextC := rc1.Col
	for nextR >= 0 && nextR <= maxR && nextC >= 0 && nextC <= maxC {
		res = append(res, RC{Row: nextR, Col: nextC})
		nextR -= deltaR
		nextC -= deltaC
	}

	nextR = rc2.Row
	nextC = rc2.Col
	for nextR >= 0 && nextR <= maxR && nextC >= 0 && nextC <= maxC {
		res = append(res, RC{Row: nextR, Col: nextC})
		nextR += deltaR
		nextC += deltaC
	}

	return res
}

func loadInput() (Nodes, int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := make(Nodes)

	row := 0
	col := 0
	for scanner.Scan() {
		line := scanner.Text()

		col = len(line)

		for col, char := range line {
			if char != '.' {
				vv, ex := res[char]
				if !ex {
					vv = make([]RC, 0)
				}

				vv = append(vv, RC{Row: row, Col: col})

				res[char] = vv
			}
		}

		row++
	}

	return res, row - 1, col - 1
}
