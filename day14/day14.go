package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Robot struct {
	X, Y, Vx, Vy, DestX, DestY int
}

func main() {
	fmt.Println("Day 14:")
	part1()
	part2()
}

func part1() {
	robots, width, height := loadInput()

	seconds := 100
	centerX := (width - 1) / 2
	centerY := (height - 1) / 2

	q0 := 0
	q1 := 0
	q2 := 0
	q3 := 0

	for i := 0; i < len(robots); i++ {
		robots[i].DestX = mod(robots[i].X+robots[i].Vx*seconds, width)
		robots[i].DestY = mod(robots[i].Y+robots[i].Vy*seconds, height)

		if robots[i].DestX < centerX && robots[i].DestY < centerY {
			q0++
		}

		if robots[i].DestX > centerX && robots[i].DestY < centerY {
			q1++
		}

		if robots[i].DestX < centerX && robots[i].DestY > centerY {
			q3++
		}

		if robots[i].DestX > centerX && robots[i].DestY > centerY {
			q2++
		}
	}

	fmt.Println("Part1:", q0, q1, q2, q3, "::", q0*q1*q2*q3)
}

func part2() {
	robots, width, height := loadInput()

	seconds := 0
	// centerX := (width - 1) / 2
	// centerY := (height - 1) / 2

	for {
		// q0 := 0
		// q1 := 0
		// q2 := 0
		// q3 := 0

		for i := 0; i < len(robots); i++ {
			robots[i].DestX = mod(robots[i].X+robots[i].Vx*seconds, width)
			robots[i].DestY = mod(robots[i].Y+robots[i].Vy*seconds, height)
			// if robots[i].DestX < centerX && robots[i].DestY < centerY {
			// 	q0++
			// }

			// if robots[i].DestX > centerX && robots[i].DestY < centerY {
			// 	q1++
			// }

			// if robots[i].DestX < centerX && robots[i].DestY > centerY {
			// 	q3++
			// }

			// if robots[i].DestX > centerX && robots[i].DestY > centerY {
			// 	q2++
			// }
		}

		if checkCond1(robots) {
			fmt.Println("S:", seconds)
			print(robots, width, height)
			fmt.Println()

			fmt.Scanln()
		}

		if seconds%1_000_000 == 0 {
			fmt.Println("S:", seconds)
		}

		seconds++
	}

	// fmt.Println("Part2:")
}

func checkCond1(robots []Robot) bool {
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].DestY < robots[j].DestY ||
			robots[i].DestY == robots[j].DestY && robots[i].DestX < robots[j].DestX
	})

	cnt := 0
	for i := 2; i < len(robots); i++ {
		if robots[i-2].DestX+1 == robots[i-1].DestX && robots[i-1].DestX+1 == robots[i].DestX {
			cnt++
		}
	}

	fmt.Println("CNT:", cnt)
	return cnt >= 10
}

func print(robots []Robot, width, height int) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			has := false
			for i := 0; i < len(robots); i++ {
				if robots[i].DestX == col && robots[i].DestY == row {
					has = true
					break
				}
			}
			if has {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func loadInput() ([]Robot, int, int) {
	file, _ := os.Open("input.txt")
	width := 101
	height := 103

	// file, _ := os.Open("input-0.txt")
	// width := 11
	// height := 7

	scanner := bufio.NewScanner(file)

	// p=9,5 v=-3,-3
	res := []Robot{}

	for scanner.Scan() {
		line := scanner.Text()

		temp := Robot{}

		parts0 := strings.Split(line, " v=")
		parts1 := strings.Split(parts0[0], "=")

		coordsStr := strings.Split(parts1[1], ",")
		temp.X, _ = strconv.Atoi(coordsStr[0])
		temp.Y, _ = strconv.Atoi(coordsStr[1])

		velStr := strings.Split(parts0[1], ",")
		temp.Vx, _ = strconv.Atoi(velStr[0])
		temp.Vy, _ = strconv.Atoi(velStr[1])

		res = append(res, temp)
	}

	return res, width, height
}

func mod(a, b int) int {
	return (a%b + b) % b
}
