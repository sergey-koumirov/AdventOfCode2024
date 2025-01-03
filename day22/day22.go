package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 22:")

	part1()
	part2()
}

func part1() {
	nn := loadInput()

	for i := 1; i <= 2000; i++ {
		for j, n := range nn {
			nn[j] = evolve(n)
		}
	}

	sum := 0
	for _, n := range nn {
		sum += n
	}

	fmt.Println("Part1:", sum)
}

type Fours map[int]int

func part2() {
	starts := loadInput()

	steps := 2000
	// steps := 9

	allFours := []Fours{}
	allKeys := map[int]bool{}

	for _, n := range starts {
		nn := make([]int, steps+1)

		nn[0] = n
		for i := 1; i <= steps; i++ {
			nn[i] = evolve(nn[i-1])
		}

		changes := make([]int, steps)
		for i := range changes {
			changes[i] = nn[i+1]%10 - nn[i]%10
		}

		fours := Fours{}

		for i := 3; i < len(changes); i++ {
			key := k(changes[i], changes[i-1], changes[i-2], changes[i-3])
			_, ex := fours[key]
			if !ex {
				fours[key] = nn[i+1] % 10
				allKeys[key] = true
			}
		}

		allFours = append(allFours, fours)
	}

	max := 0
	for k := range allKeys {
		test := 0
		for _, fours := range allFours {
			val, ex := fours[k]
			if ex {
				test += val
			}
		}

		if test > max {
			max = test
		}

		// fmt.Println(keyToStr(k), "/", test)
	}

	fmt.Println("Part2:", max)
}

func k(n1, n2, n3, n4 int) int {
	return (n1+10)*1000000 + (n2+10)*10000 + (n3+10)*100 + (n4 + 10)
}

func keyToStr(n int) string {
	n1 := n%100 - 10

	n2 := (n/100)%100 - 10

	n3 := (n/10000)%100 - 10

	n4 := (n/1000000)%100 - 10

	return fmt.Sprintf("%d %d %d %d", n1, n2, n3, n4)
}

func evolve(n int) int {
	n64 := n * 64
	nMixed := mix(n64, n)
	nPrune := prune(nMixed)

	n32 := nPrune / 32
	nMixed2 := mix(n32, nPrune)
	nPrune2 := prune(nMixed2)

	n2048 := nPrune2 * 2048
	nMixed3 := mix(n2048, nPrune2)
	nPrune3 := prune(nMixed3)

	return nPrune3
}

func mix(n64, n int) int {
	return n64 ^ n
}

func prune(n int) int {
	return n % 16777216
}

func loadInput() []int {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	res := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		n, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println(err)
		} else {
			res = append(res, n)
		}
	}

	return res
}
