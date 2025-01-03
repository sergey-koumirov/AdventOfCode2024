package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 19")
	// part1()
	part2()
}

func part2() {
	towels, patterns, _, max := loadInput()

	cnt := 0
	for _, pattern := range patterns {
		cache := map[int]int{}
		countPatterns(towels, pattern, 0, max, cache)
		cntPattern := cache[0]
		cnt += cntPattern
	}

	fmt.Println("Part2", cnt)
}

func countPatterns(towels map[string]bool, pattern string, index, max int, cache map[int]int) int {

	cnt, ex := cache[index]
	if ex {
		return cnt
	}

	newCnt := 0

	for i := 1; i <= max; i++ {
		if index+i <= len(pattern) {
			part := pattern[index : index+i]

			_, ex := towels[part]

			if ex {
				// fmt.Println("   E: ", part)
				if index+i >= len(pattern) {
					newCnt++
				} else {
					newCnt += countPatterns(towels, pattern, index+i, max, cache)
				}
			}
		}
	}
	cache[index] = newCnt

	return newCnt
}

func part1() {
	towels, patterns, _, max := loadInput()

	cnt := 0
	for i, pattern := range patterns {
		possible := false
		checkPattern(towels, pattern, max, 0, &possible)
		if possible {
			cnt++
		}
		fmt.Println(i+1, "/", len(patterns))
	}

	fmt.Println("Part1", cnt)
}

func checkPattern(towels map[string]bool, pattern string, max, index int, res *bool) {
	// fmt.Println(pattern, index)

	for i := 1; i <= max; i++ {
		if !*res && index+i <= len(pattern) {
			part := pattern[index : index+i]

			_, ex := towels[part]

			if ex {
				// fmt.Println("   E: ", part)
				if index+i >= len(pattern) {
					*res = true
				} else {
					checkPattern(towels, pattern, max, index+i, res)
				}
			}
		}
	}
}

func loadInput() (map[string]bool, []string, int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	max := 0
	min := 9999
	towels := map[string]bool{}

	parts := strings.Split(line, ", ")
	for _, part := range parts {
		towels[part] = true
		if len(part) > max {
			max = len(part)
		}
		if len(part) < min {
			min = len(part)
		}
	}

	scanner.Scan()

	res := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return towels, res, min, max
}
