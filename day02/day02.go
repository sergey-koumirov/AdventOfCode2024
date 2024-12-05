package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 02:")
	part1()
	part2()
}

func part1() {
	reports := loadInput()

	cnt := 0
	for _, report := range reports {
		isAsc := report[0]-report[1] > 0
		isGood := true

		for i := 1; i < len(report); i++ {
			delta := report[i-1] - report[i]
			isAscPair := delta > 0

			if isAsc != isAscPair || math.Abs(float64(delta)) < 1 || math.Abs(float64(delta)) > 3 {
				isGood = false
			}
		}

		if isGood {
			cnt += 1
		}
	}

	fmt.Println("P1:", cnt)
}

func part2() {
	reports := loadInput()

	cnt := 0
	for _, report := range reports {
		if isGoodReport(report) {
			cnt += 1
		} else {
			removeIndex := 0
			testIsGood := false
			for removeIndex < len(report) && !testIsGood {
				test := copyWithoutIndex(report, removeIndex)
				testIsGood = isGoodReport(test)
				removeIndex += 1
			}

			if testIsGood {
				cnt += 1
			}
		}
	}

	fmt.Println("P2:", cnt)
}

func copyWithoutIndex(report []int, index int) []int {
	res := make([]int, 0)

	for i := 0; i < len(report); i++ {
		if i != index {
			res = append(res, report[i])
		}
	}

	return res
}

func isGoodReport(report []int) bool {
	isAsc := report[0]-report[1] > 0
	isGood := true

	for i := 1; i < len(report); i++ {
		delta := report[i-1] - report[i]
		isAscPair := delta > 0

		if isAsc != isAscPair || math.Abs(float64(delta)) < 1 || math.Abs(float64(delta)) > 3 {
			isGood = false
		}
	}

	return isGood
}

func loadInput() [][]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		temp := make([]int, len(parts))

		for i, str := range parts {
			num, _ := strconv.Atoi(str)
			temp[i] = num
		}

		res = append(res, temp)
	}

	return res
}
