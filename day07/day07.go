package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Result int
	Params []int
}

func main() {
	fmt.Println("Day 07:")
	part1()
	fmt.Println()
	part2()
}

func part1() {
	records := loadInput()

	sum := 0
	for _, record := range records {
		if isCorrect(record) {
			sum += record.Result
		}
	}

	fmt.Println("Part1:", sum)
}

func part2() {
	records := loadInput()

	sum := 0
	for _, record := range records {
		if isCorrectPart2(record) {
			sum += record.Result
		}
	}

	fmt.Println("Part2:", sum)
}

func isCorrectPart2(r Record) bool {
	res := r.Params[0]
	return applyNextPart2(r, res, 1)
}

func applyNextPart2(r Record, res int, index int) bool {
	if index >= len(r.Params) {
		return res == r.Result
	}

	if applyNextPart2(r, res+r.Params[index], index+1) {
		return true
	}

	if applyNextPart2(r, res*r.Params[index], index+1) {
		return true
	}

	str := fmt.Sprintf("%d%d", res, r.Params[index])
	test, _ := strconv.Atoi(str)

	if applyNextPart2(r, test, index+1) {
		return true
	}

	return false
}

func isCorrect(r Record) bool {
	res := r.Params[0]

	return applyNext(r, res, 1)
}

func applyNext(r Record, res int, index int) bool {
	if index >= len(r.Params) {
		return res == r.Result
	}

	if applyNext(r, res+r.Params[index], index+1) {
		return true
	}

	if applyNext(r, res*r.Params[index], index+1) {
		return true
	}

	return false
}

func loadInput() []Record {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := make([]Record, 0)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")

		r, _ := strconv.Atoi(parts[0])

		temp := Record{
			Result: r,
			Params: parseInts(parts[1], " "),
		}

		res = append(res, temp)
	}

	return res
}

func parseInts(line string, divider string) []int {
	parts := strings.Split(line, divider)

	res := []int{}
	for _, part := range parts {
		if part != "" {
			n, _ := strconv.Atoi(part)
			res = append(res, n)
		}
	}
	return res
}
