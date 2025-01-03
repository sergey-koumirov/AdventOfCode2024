package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	AX, AY, BX, BY, X, Y int
}

func main() {
	fmt.Println("Day 13:")
	part1()
	part2()
}

func part1() {
	eqs := loadInput()

	sum := 0

	for _, eq := range eqs {
		n1 := float64(eq.BY*eq.X-eq.BX*eq.Y) / float64(eq.AX*eq.BY-eq.BX*eq.AY)

		n2 := (float64(eq.Y) - float64(eq.AY)*n1) / float64(eq.BY)

		if n1 == math.Floor(n1) && n2 == math.Floor(n2) && n1 <= 100 && n2 <= 100 {
			// fmt.Println(n1, n2, n1*3+n2)
			sum += int(n1)*3 + int(n2)
		}
	}

	fmt.Println("Part1:", sum)
}

func part2() {
	eqs := loadInput()

	for i := 0; i < len(eqs); i++ {
		eqs[i].X = eqs[i].X + 10000000000000
		eqs[i].Y = eqs[i].Y + 10000000000000
	}

	sum := 0

	for i, eq := range eqs {

		num1 := eq.BY*eq.X - eq.BX*eq.Y
		num2 := eq.AX*eq.BY - eq.BX*eq.AY
		if num1%num2 == 0 {
			n1 := num1 / num2

			num3 := eq.Y - eq.AY*n1
			num4 := eq.BY

			if num3%num4 == 0 {
				n2 := num3 / num4

				// fmt.Println(i+1, n1, n2, n1*3+n2)
				sum += int(n1)*3 + int(n2)
			}
		}
	}

	fmt.Println("Part2:", sum)
}

func loadInput() []Equation {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := []Equation{}

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	for i := 0; i < len(lines)/3; i++ {
		parts1 := strings.Split(lines[i*3+0], ": ")
		parts2 := strings.Split(parts1[1], ", ")
		axStr := strings.Split(parts2[0], "+")[1]
		ayStr := strings.Split(parts2[1], "+")[1]

		parts3 := strings.Split(lines[i*3+1], ": ")
		parts4 := strings.Split(parts3[1], ", ")
		bxStr := strings.Split(parts4[0], "+")[1]
		byStr := strings.Split(parts4[1], "+")[1]

		parts5 := strings.Split(lines[i*3+2], ": ")
		parts6 := strings.Split(parts5[1], ", ")
		xStr := strings.Split(parts6[0], "=")[1]
		yStr := strings.Split(parts6[1], "=")[1]

		ax, _ := strconv.Atoi(axStr)
		ay, _ := strconv.Atoi(ayStr)
		bx, _ := strconv.Atoi(bxStr)
		by, _ := strconv.Atoi(byStr)
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)

		res = append(res, Equation{AX: ax, BX: bx, AY: ay, BY: by, X: x, Y: y})
	}

	return res
}
