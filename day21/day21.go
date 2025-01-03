package main

import (
	"fmt"
	"strconv"
	"strings"
)

var passcodes = []string{"826A", "341A", "582A", "983A", "670A"}

// var passcodes = []string{"3A"}
// var passcodes = []string{"029A"}
// var passcodes = []string{"029A", "980A", "179A", "456A", "379A"}

var keypad1 = []string{
	"789",
	"456",
	"123",
	" 0A",
}

var keypad2 = []string{
	" ^A",
	"<v>",
}

var pathes1 map[string][]string
var pathes2 map[string][]string
var gc int

func main() {
	pathes1 = precalc(keypad1)

	pathes2 = precalc(keypad2)

	part1()
	part2()
}

func part1() {
	sum := 0
	gc = 0

	cache := map[string]int{}

	for _, passcode := range passcodes {
		minLen := loop1(passcode, 2, cache)
		n, _ := strconv.Atoi(passcode[0 : len(passcode)-1])
		sum += minLen * n
		fmt.Println(passcode, minLen, "*", n, "=", sum)
	}

	fmt.Println("Part1", sum, "[", gc, "]")
}

func part2() {
	sum := 0
	gc = 0

	cache := map[string]int{}

	for _, passcode := range passcodes {
		minLen := loop1(passcode, 25, cache)
		n, _ := strconv.Atoi(passcode[0 : len(passcode)-1])
		sum += minLen * n
		fmt.Println(passcode, minLen, "*", n, "=", sum)
	}

	fmt.Println("Part2", sum, "[", gc, "]")
}

func loop1(passcode string, maxDeep int, cache map[string]int) int {
	fullPasscode := "A" + passcode

	sum := 0

	for i := 1; i < len(fullPasscode); i++ {
		key := string(fullPasscode[i-1]) + "-" + string(fullPasscode[i])

		min := -1
		pathes := pathes1[key]

		for _, path := range pathes {
			pathLen := deepLoop(path, 1, maxDeep, cache)
			if min == -1 || min > pathLen {
				min = pathLen
			}
		}

		if min == -1 {
			fmt.Println("Problem: ", key)
		} else {
			sum += min
		}
	}

	return sum
}

func deepLoop(path string, deep, maxDeep int, cache map[string]int) int {
	gc++
	if deep > maxDeep {
		return len(path)
	}

	fullPath := "A" + path

	sum := 0

	for i := 1; i < len(fullPath); i++ {
		key := string(fullPath[i-1]) + "-" + string(fullPath[i])

		min := -1
		pathes := pathes2[key]

		for _, path := range pathes {

			cacheKey := fmt.Sprintf("%d-%s", deep+1, path)
			pathLen, ex := cache[cacheKey]
			if !ex {
				pathLen = deepLoop(path, deep+1, maxDeep, cache)
				cache[cacheKey] = pathLen
			}

			if min == -1 || min > pathLen {
				min = pathLen
			}
		}

		if min == -1 {
			fmt.Println("Problem: ", key)
		} else {
			sum += min
		}
	}

	return sum
}

func precalc(keypad []string) map[string][]string {
	symbols := strings.Join(keypad, "")

	res := map[string][]string{}

	for i1, char1 := range symbols {
		for i2, char2 := range symbols {
			if i1 != i2 && char1 != ' ' && char2 != ' ' {
				row1 := i1 / 3
				col1 := i1 % 3

				row2 := i2 / 3
				col2 := i2 % 3

				pathes := findAllPathes(row1, col1, row2, col2, keypad)
				key := fmt.Sprintf("%c-%c", char1, char2)

				res[key] = pathes
			} else if i1 == i2 && char1 != ' ' && char2 != ' ' {
				key := fmt.Sprintf("%c-%c", char1, char2)
				res[key] = []string{"A"}
			}
		}
	}

	return res
}

func findAllPathes(row1, col1, row2, col2 int, keypad []string) []string {
	dist := abs(row2-row1) + abs(col2-col1)

	res := []string{}

	deepCheck(row1, col1, row2, col2, 1, dist, "", &res, keypad)

	return res
}

func deepCheck(row1, col1, row2, col2, level, maxLevel int, path string, res *[]string, keypad []string) {
	if keypad[row1][col1] == ' ' {
		return
	}

	if row1 == row2 && col1 == col2 {
		*res = append(*res, path+"A")
		return
	}

	if level <= maxLevel {
		if row1 > 0 {
			deepCheck(row1-1, col1, row2, col2, level+1, maxLevel, path+"^", res, keypad)
		}

		if row1 < len(keypad)-1 {
			deepCheck(row1+1, col1, row2, col2, level+1, maxLevel, path+"v", res, keypad)
		}

		if col1 > 0 {
			deepCheck(row1, col1-1, row2, col2, level+1, maxLevel, path+"<", res, keypad)
		}
		if col1 < len(keypad[0])-1 {
			deepCheck(row1, col1+1, row2, col2, level+1, maxLevel, path+">", res, keypad)
		}
	}
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
