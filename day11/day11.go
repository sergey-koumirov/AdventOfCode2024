package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RC struct {
	Row, Col int
}

func main() {
	fmt.Println("Day 11:")
	part1()
	fmt.Println()
	part2()
}

func part1() {
	nums := loadInput()

	cnt := 0

	for _, num := range nums {
		temp := 1
		deep(num, 1, 25, &temp)
		cnt += temp
	}

	fmt.Println("Part1:", cnt)
}

func part2() {
	nums := loadInput()

	cnt := 0

	cache := map[string]int{}

	for _, num := range nums {
		runs := 0
		test := deepPart2(num, 75, cache, &runs)
		cnt += test
	}

	fmt.Println("Part2:", cnt)
}

func deepPart2(num string, level int, cache map[string]int, runs *int) int {
	*runs++

	if level <= 0 {
		return 1
	}

	key := num + "-" + strconv.Itoa(level)
	precalc, ex := cache[key]
	if ex {
		// fmt.Println(key, precalc)
		return precalc
	}

	if num == "0" {
		cnt := deepPart2("1", level-1, cache, runs)
		cache[key] = cnt
		return cnt
	}

	if len(num)%2 == 0 {
		left := num[:len(num)/2]
		right := num[len(num)/2:]

		if strings.Count(right, "0") == len(right) {
			right = "0"
		} else {
			right = strings.TrimLeft(right, "0")
		}

		lCnt := deepPart2(left, level-1, cache, runs)
		rCnt := deepPart2(right, level-1, cache, runs)
		return lCnt + rCnt
	}

	newNum, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("E:", err)
	}

	newStr := strconv.Itoa(newNum * 2024)

	cnt2 := deepPart2(newStr, level-1, cache, runs)
	cache[key] = cnt2

	return cnt2
}

func deep(num string, level, max int, cnt *int) {

	if level > max {
		return
	}

	if num == "0" {
		deep("1", level+1, max, cnt)
	} else if len(num)%2 == 0 {
		left := num[:len(num)/2]
		right := num[len(num)/2:]

		if strings.Count(right, "0") == len(right) {
			right = "0"
		} else {
			right = strings.TrimLeft(right, "0")
		}

		*cnt++

		deep(left, level+1, max, cnt)
		deep(right, level+1, max, cnt)
	} else {
		newNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("E:", err)
		}

		newStr := strconv.Itoa(newNum * 2024)
		deep(newStr, level+1, max, cnt)
	}
}

func loadInput() []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	return strings.Split(line, " ")
}
