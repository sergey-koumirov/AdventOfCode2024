package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Fragment struct {
	ID, Len int
}

type Fragments []Fragment

func main() {
	fmt.Println("Day 09:")
	part1()
	fmt.Println()
	part2()
}

func part1() {
	fragments := loadInput()

	for {
		emptyIndex := firstEmpty(fragments)
		if emptyIndex == -1 {
			break
		}

		lastIndex := len(fragments) - 1

		delta := fragments[emptyIndex].Len - fragments[lastIndex].Len

		if delta > 0 {
			fragments[emptyIndex] = fragments[lastIndex]
			fragments = slices.Concat(
				fragments[:emptyIndex+1],
				Fragments{Fragment{ID: -1, Len: delta}},
				fragments[emptyIndex+1:lastIndex],
			)
		} else if delta == 0 {
			fragments[emptyIndex] = fragments[lastIndex]
			fragments = fragments[:lastIndex]
		} else if delta < 0 {
			fragments[emptyIndex].ID = fragments[lastIndex].ID
			fragments[lastIndex].Len = fragments[lastIndex].Len - fragments[emptyIndex].Len
		}
	}

	cs := checksum(fragments)

	fmt.Println("Part1:", cs)
}

func part2() {
	fragments := loadInput()

	lastIndex := len(fragments) - 1

	// print(fragments)

	for {
		emptyIndex := firstEmptyPart2(fragments, fragments[lastIndex].Len, lastIndex)

		if emptyIndex > -1 {
			delta := fragments[emptyIndex].Len - fragments[lastIndex].Len

			if delta > 0 {
				fragments[emptyIndex] = fragments[lastIndex]
				fragments[lastIndex].ID = -1

				fragments = slices.Concat(
					fragments[:emptyIndex+1],
					Fragments{Fragment{ID: -1, Len: delta}},
					fragments[emptyIndex+1:],
				)

				lastIndex++
			} else if delta == 0 {
				fragments[emptyIndex] = fragments[lastIndex]
				fragments[lastIndex].ID = -1
			}
		}

		for lastIndex > -1 {
			lastIndex--
			if lastIndex < 0 || fragments[lastIndex].ID > -1 {
				break
			}
		}

		// print(fragments)

		if lastIndex < 0 {
			break
		}
	}

	cs := checksum(fragments)

	fmt.Println("Part2:", cs)
}

func print(ff Fragments) {
	for _, f := range ff {
		for i := 0; i < f.Len; i++ {
			if f.ID == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(f.ID)
			}

		}
	}
	fmt.Println()
}

func checksum(ff Fragments) int {
	mainIndex := 0
	res := 0
	for _, f := range ff {
		if f.ID > -1 {
			res += f.ID * (mainIndex*f.Len + (f.Len-1)*(f.Len-1+1)/2)
		}
		mainIndex += f.Len
	}
	return res
}

func firstEmpty(ff Fragments) int {
	for index, f := range ff {
		if f.ID == -1 {
			return index
		}
	}
	return -1
}

func firstEmptyPart2(ff Fragments, minLen int, beforeIndex int) int {
	for index, f := range ff {
		if index >= beforeIndex {
			return -1
		}
		if f.ID == -1 && f.Len >= minLen {
			return index
		}
	}
	return -1
}

func loadInput() Fragments {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	res := Fragments{}

	id := 0
	for index, char := range line {
		partLen := int(char) - 48
		if index%2 == 0 {
			res = append(res, Fragment{ID: id, Len: partLen})
			id++
		} else if partLen > 0 {
			res = append(res, Fragment{ID: -1, Len: partLen})
		}
	}

	return res
}
