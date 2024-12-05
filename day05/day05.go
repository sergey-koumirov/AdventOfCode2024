package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule map[int]bool
type Rules map[int]Rule

type Update []int
type Updates []Update

func main() {
	fmt.Println("Day 05:")
	part1()
	part2()
}

func part1() {
	rules, updates := loadInput()

	sum := 0
	for _, update := range updates {
		if updateIsCorrect(update, rules) {
			sum += update[len(update)/2]
		}
	}

	fmt.Println("Part1:", sum)
	fmt.Println()
}

func part2() {
	rules, updates := loadInput()

	sum := 0
	for _, update := range updates {
		if !updateIsCorrect(update, rules) {
			fixed := fixUpdate(update, rules)

			fmt.Println(fixed)
			sum += fixed[len(fixed)/2]
		}
	}

	fmt.Println("Part2:", sum)
}

func fixUpdate(update Update, rules Rules) Update {
	res := make(Update, len(update))

	for i, el := range update {
		res[i] = el
	}

	for index := 1; index < len(update); index++ {
		allowed, exAllowed := rules[res[index]]

		if exAllowed {
			failIndex := -1
			for prevIndex := 0; prevIndex < index; prevIndex++ {
				found, ex := allowed[res[prevIndex]]
				if ex && found {
					failIndex = prevIndex
					break
				}
			}

			if failIndex > -1 {
				part1 := res[:failIndex]
				part2 := res[failIndex:index]
				part3 := res[index+1:]
				res = slices.Concat(part1, []int{res[index]}, part2, part3)
			}
		}
	}

	return res
}

func updateIsCorrect(update Update, rules Rules) bool {
	isCorrect := true
	for index := len(update) - 2; index >= 0; index-- {
		candidate := update[index]

		for checkIndex := index; checkIndex < len(update); checkIndex++ {
			element := update[checkIndex]

			forbiddens, exEl := rules[element]
			if exEl {
				found, exFr := forbiddens[candidate]
				if exFr && found {
					isCorrect = false
					break
				}
			}
		}

		if !isCorrect {
			break
		}
	}

	return isCorrect
}

func loadInput() (Rules, Updates) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	rules := make(Rules)
	updates := make(Updates, 0)

	processRule := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			processRule = false
		} else if processRule {
			parts := strings.Split(line, "|")
			n1, _ := strconv.Atoi(parts[0])
			n2, _ := strconv.Atoi(parts[1])

			_, ex := rules[n1]
			if !ex {
				rules[n1] = make(Rule)
			}
			rules[n1][n2] = true
		} else {
			updates = append(updates, parseInts(line))
		}
	}

	return rules, updates
}

func parseInts(line string) []int {
	parts := strings.Split(line, ",")

	res := []int{}
	for _, part := range parts {
		if part != "" {
			n, _ := strconv.Atoi(part)
			res = append(res, n)
		}
	}
	return res
}
