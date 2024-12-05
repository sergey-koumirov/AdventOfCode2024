package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 03:")
	part1()
	part2()
}

func part1() {
	lines := loadInput()

	sum := 0
	for _, line := range lines {
		n := processLine(line)
		sum += n
	}

	fmt.Println("Part1:", sum)
}

func part2() {
	lines := loadInput()

	sum := 0
	canAdd := true
	for _, line := range lines {
		for index := 0; index < len(line)-8; index++ {
			exists, n1, n2 := checkPattern(line, index)

			if exists && canAdd {
				// fmt.Println(n1, n2)
				sum += n1 * n2
			} else if isStr(line, "do()", index) {
				canAdd = true
			} else if isStr(line, "don't()", index) {
				canAdd = false
			}
		}
	}

	fmt.Println("Part2:", sum)
}

func processLine(line string) int {
	sum := 0

	for index := 0; index < len(line)-8; index++ {
		exists, n1, n2 := checkPattern(line, index)

		if exists {
			// fmt.Println(n1, n2)
			sum += n1 * n2
		}
	}

	return sum
}

func isStr(line string, str string, startIndex int) bool {
	lineLen := len(line)

	if lineLen-1-startIndex < len(str) {
		return false
	}

	for index, char := range str {
		if line[startIndex+index] != byte(char) {
			return false
		}
	}

	return true
}

func checkPattern(line string, startIndex int) (bool, int, int) {
	lineLen := len(line)
	ch1 := lineLen > startIndex && line[startIndex] != 'm'
	ch2 := lineLen > startIndex+1 && line[startIndex+1] != 'u'
	ch3 := lineLen > startIndex+2 && line[startIndex+2] != 'l'
	ch4 := lineLen > startIndex+3 && line[startIndex+3] != '('

	if ch1 || ch2 || ch3 || ch4 {
		return false, 0, 1
	}

	nextIndex := startIndex + 4

	n1 := 0
	for delta := 0; delta <= 2; delta++ {
		if isNum(line[nextIndex]) {
			n1 = n1*10 + int(line[nextIndex]-48)
			nextIndex += 1
		} else if line[nextIndex] == ',' {
			break
		} else {
			return false, 0, 2
		}
	}

	if line[nextIndex] != ',' {
		return false, 0, 3
	}
	nextIndex += 1

	n2 := 0

	for delta := 0; delta <= 2; delta++ {
		if isNum(line[nextIndex]) {
			n2 = n2*10 + int(line[nextIndex]-48)
			nextIndex += 1
		} else if line[nextIndex] == ')' {
			break
		} else {
			return false, 0, 4
		}
	}

	if line[nextIndex] != ')' {
		return false, 0, 5
	}

	return true, n1, n2
}

func isNum(s byte) bool {
	return s == '0' || s == '1' || s == '2' || s == '3' || s == '4' || s == '5' || s == '6' || s == '7' || s == '8' || s == '9'
}

func loadInput() []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return res
}
