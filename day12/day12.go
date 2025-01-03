package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Cell struct {
	Used bool
	Kind rune
}

type Field struct {
	Cells          [][]Cell
	MaxRow, MaxCol int
}

type RC struct {
	Row, Col int
}

type Side struct {
	AreaCode, Dir, Row, Col int
}

func main() {
	fmt.Println("Day 12:")
	part1()
}

func part1() {
	field := loadInput()

	sum := 0
	sides := []Side{}
	areaCode := 0
	codeToArea := map[int]int{}

	for row := 0; row <= field.MaxRow; row++ {
		for col := 0; col <= field.MaxCol; col++ {
			if !field.Cells[row][col].Used {
				area, perimeter := evaluate(&field, row, col, &sides, areaCode)
				sum += area * perimeter
				codeToArea[areaCode] = area
				areaCode++
			}
		}
	}

	fmt.Println("Part1:", sum)

	areaSides := map[int][]Side{}
	for index, side := range sides {
		sidesForCode, ex := areaSides[side.AreaCode]
		if !ex {
			sidesForCode = []Side{}
		}
		sidesForCode = append(sidesForCode, sides[index])
		areaSides[side.AreaCode] = sidesForCode
	}

	sum2 := 0

	for areaCode, vv := range areaSides {
		// // 1, 7 top + bottom
		// fmt.Println("1/7")
		sort.Slice(vv, func(i, j int) bool {
			return vv[i].Dir < vv[j].Dir ||
				vv[i].Dir == vv[j].Dir && vv[i].Row < vv[j].Row ||
				vv[i].Dir == vv[j].Dir && vv[i].Row == vv[j].Row && vv[i].Col < vv[j].Col
		})

		linesCnt := 0
		lineLen := 0
		for index := range vv {
			if vv[index].Dir == 1 || vv[index].Dir == 7 {
				// fmt.Print(vv[index])
				if index == 0 ||

					vv[index-1].Dir == vv[index].Dir &&
						vv[index-1].Row == vv[index].Row &&
						vv[index-1].Col+1 == vv[index].Col {
					lineLen++
					// fmt.Println(" L++")
				} else if lineLen > 0 {
					linesCnt++
					lineLen = 1
					// fmt.Println(" C++")
				} else {
					lineLen = 1
					// fmt.Println(" R")
				}
			}
		}
		if lineLen > 0 {
			linesCnt++
			// fmt.Println("LAST", linesCnt)
		}

		// // 3, 5 left + right
		// fmt.Println("3/5")
		lineLen = 0

		sort.Slice(vv, func(i, j int) bool {
			return vv[i].Dir < vv[j].Dir ||
				vv[i].Dir == vv[j].Dir && vv[i].Col < vv[j].Col ||
				vv[i].Dir == vv[j].Dir && vv[i].Col == vv[j].Col && vv[i].Row < vv[j].Row
		})

		for index := range vv {
			if vv[index].Dir == 3 || vv[index].Dir == 5 {
				// fmt.Print(vv[index])
				if index == 0 ||
					vv[index-1].Dir == vv[index].Dir &&
						vv[index-1].Col == vv[index].Col &&
						vv[index-1].Row+1 == vv[index].Row {
					lineLen++
					// fmt.Println(" L++")
				} else if lineLen > 0 {
					linesCnt++
					lineLen = 1
					// fmt.Println(" C++")
				} else {
					lineLen = 1
					// fmt.Println(" R")
				}
			}
		}
		if lineLen > 0 {
			linesCnt++
			// fmt.Println("LAST")
		}

		sum2 += codeToArea[areaCode] * linesCnt
		// fmt.Println("Area #", areaCode, " = ", codeToArea[areaCode], linesCnt, sum2)
		// fmt.Println()
	}

	fmt.Println("Part2:", sum2)
}

func evaluate(f *Field, row, col int, sides *[]Side, areaCode int) (int, int) {

	f.Cells[row][col].Used = true
	buffer := []RC{{Row: row, Col: col}}
	area := 1
	perimeter := 0

	for {
		if len(buffer) == 0 {
			break
		}

		newBuffer := []RC{}
		for _, el := range buffer {
			baseCell := f.Cells[el.Row][el.Col]
			evaluateCell(f, el, baseCell, 0, -1, &newBuffer, &perimeter, &area, sides, areaCode)
			evaluateCell(f, el, baseCell, 0, 1, &newBuffer, &perimeter, &area, sides, areaCode)
			evaluateCell(f, el, baseCell, -1, 0, &newBuffer, &perimeter, &area, sides, areaCode)
			evaluateCell(f, el, baseCell, 1, 0, &newBuffer, &perimeter, &area, sides, areaCode)
		}

		buffer = newBuffer
	}

	return area, perimeter
}

func evaluateCell(f *Field, el RC, baseCell Cell, dr, dc int, newBuffer *[]RC, perimeter *int, area *int, sides *[]Side, areaCode int) {
	if el.Row+dr >= 0 && el.Row+dr <= f.MaxRow && el.Col+dc >= 0 && el.Col+dc <= f.MaxCol {
		cell := f.Cells[el.Row+dr][el.Col+dc]
		if cell.Kind == baseCell.Kind {
			if !cell.Used {
				*newBuffer = append(*newBuffer, RC{Row: el.Row + dr, Col: el.Col + dc})
				f.Cells[el.Row+dr][el.Col+dc].Used = true
				*area++
			}
		} else {
			*perimeter++
			*sides = append(*sides, Side{AreaCode: areaCode, Dir: 3*(dr+1) + dc + 1, Row: el.Row, Col: el.Col})
		}
	} else {
		*perimeter++
		*sides = append(*sides, Side{AreaCode: areaCode, Dir: 3*(dr+1) + dc + 1, Row: el.Row, Col: el.Col})
	}
}

func loadInput() Field {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := Field{}
	res.Cells = make([][]Cell, 0)

	for scanner.Scan() {
		line := scanner.Text()

		temp := []Cell{}
		for _, char := range line {
			temp = append(temp, Cell{Used: false, Kind: char})
		}

		res.Cells = append(res.Cells, temp)
	}

	res.MaxRow = len(res.Cells) - 1
	res.MaxCol = len(res.Cells[0]) - 1

	return res
}
