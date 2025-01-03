package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 25:")

	part1()
	part2()
}

func part1() {
	keys, locks := loadInput()

	cnt := 0
	for _, key := range keys {
		for _, lock := range locks {
			fit := true
			for i := 0; i <= 4; i++ {
				if key[i]+lock[i] > 5 {
					fit = false
					break
				}
			}

			if fit {
				cnt++
			}
		}
	}

	fmt.Println("Part1:", cnt)
}

func part2() {
	// pairs, nodes := loadInput()

	fmt.Println("Part2:")
}

func loadInput() ([][]int, [][]int) {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	keys := [][]int{}
	locks := [][]int{}

	buffer := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		buffer = append(buffer, line)
		if line == "" {
			processBuffer(buffer, &keys, &locks)
			buffer = []string{}
		}

	}

	processBuffer(buffer, &keys, &locks)

	return keys, locks
}

func processBuffer(buffer []string, keys *[][]int, locks *[][]int) {
	if buffer[0] == "....." {
		key := make([]int, 5)
		for i := 1; i <= 5; i++ {
			for index, char := range buffer[i] {
				if char == '#' {
					key[index]++
				}
			}
		}
		*keys = append(*keys, key)
	}

	if buffer[0] == "#####" {
		lock := make([]int, 5)
		for i := 1; i <= 5; i++ {
			for index, char := range buffer[i] {
				if char == '#' {
					lock[index]++
				}
			}
		}
		*locks = append(*locks, lock)
	}
}
