package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Day 23:")

	part1()
	part2()
}

func part1() {
	pairs, nodes := loadInput()

	cnt := 0

	for i1 := 0; i1 < len(nodes)-2; i1++ {
		for i2 := i1; i2 < len(nodes)-1; i2++ {
			for i3 := i2; i3 < len(nodes); i3++ {
				if connected(pairs, nodes[i1], nodes[i2]) && connected(pairs, nodes[i1], nodes[i3]) && connected(pairs, nodes[i2], nodes[i3]) {
					if nodes[i1][0] == 't' || nodes[i2][0] == 't' || nodes[i3][0] == 't' {
						// fmt.Println(nodes[i1], nodes[i2], nodes[i3])
						cnt++
					}
				}
			}
		}
	}

	fmt.Println("Part1:", cnt)
}

func part2() {
	pairs, nodes := loadInput()

	skip := map[string]bool{}

	max := 0
	pass := ""

	for _, node := range nodes {
		if !skip[node] {
			password := findFull([]string{node}, copyExcept(nodes, node), pairs)

			if max < len(password) {
				max = len(password)
				pass = strings.Join(password, ",")
			}
		}
	}

	fmt.Println("Part2:", pass)
}

func findFull(used []string, unused []string, pairs map[string]map[string]bool) []string {

	newUnused := make([]string, len(unused))
	copy(newUnused, unused)

	for _, test := range unused {
		withAll := true
		for _, el := range used {
			if !connected(pairs, test, el) {
				withAll = false
				break
			}
		}

		newUnused := copyExcept(newUnused, test)

		if withAll {
			newUsed := append(used, test)
			return findFull(newUsed, newUnused, pairs)
		}
	}

	sort.Strings(used)

	return used
}

func copyExcept(nodes []string, node string) []string {
	res := make([]string, len(nodes)-1)
	index := 0
	for _, test := range nodes {
		if test != node {
			res[index] = test
			index++
		}
	}
	return res
}

func connected(pairs map[string]map[string]bool, n1, n2 string) bool {
	nn, ex1 := pairs[n1]
	if ex1 {
		_, ex2 := nn[n2]
		return ex2
	}
	return false
}

func loadInput() (map[string]map[string]bool, []string) {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	res := map[string]map[string]bool{}
	uniqsMap := map[string]bool{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "-")

		nodes1, ex1 := res[parts[0]]
		if !ex1 {
			nodes1 = map[string]bool{}
		}
		nodes1[parts[1]] = true
		res[parts[0]] = nodes1

		nodes2, ex2 := res[parts[1]]
		if !ex2 {
			nodes2 = map[string]bool{}
		}
		nodes2[parts[0]] = true
		res[parts[1]] = nodes2

		uniqsMap[parts[0]] = true
		uniqsMap[parts[1]] = true
	}

	uniqs := []string{}

	for k := range uniqsMap {
		uniqs = append(uniqs, k)
	}

	return res, uniqs
}
