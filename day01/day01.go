package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 01:")
	part1()
	part2()
}

func part1() {
	n1, n2 := loadInput()

	sort.Ints(n1)
	sort.Ints(n2)

	sum := 0
	for i := 0; i < len(n1); i++ {
		d := n1[i] - n2[i]

		if d > 0 {
			sum += d
		} else {
			sum -= d
		}
	}

	fmt.Println(sum)
}

func part2() {
	n1, n2 := loadInput()

	counts := map[int]int{}

	for _, n := range n2 {
		counts[n] += 1
	}

	sum := 0
	for _, n := range n1 {
		sum += counts[n] * n
	}
	fmt.Println(sum)
}

func loadInput() ([]int, []int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res1 := make([]int, 0)
	res2 := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		num1, _ := strconv.Atoi(parts[0])
		res1 = append(res1, num1)

		num2, _ := strconv.Atoi(parts[1])
		res2 = append(res2, num2)
	}

	return res1, res2
}
