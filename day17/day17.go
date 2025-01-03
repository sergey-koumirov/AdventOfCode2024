package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Computer struct {
	A, B, C int
	Program []int
	Pointer int
	Output  []int
}

var INames = map[int]string{0: "adv", 1: "bxl", 2: "bst", 3: "jnz", 4: "bxc", 5: "out", 6: "bdv", 7: "cdv"}

func main() {
	// part1()
	part2()
}

func part2() {
	c1 := Computer{A: 0, B: 0, C: 0, Program: []int{2, 4, 1, 1, 7, 5, 0, 3, 4, 3, 1, 6, 5, 5, 3, 0}, Pointer: 0, Output: []int{}}

	// a := 35_184_372_000_000

	base := 0b10_000_000_000_000_000_000_011_101_110_101_001_100_000_000_000

	additions := []int{
		0b010_111_010,
		0b010_111_011,
		0b010_111_111,
		0b111_111_010,
		0b111_111_011,
		0b111_111_110,
	}

	// a := 281_474_977_000_000
	// a := 0
	i := 0
	for {

		found := false
		for _, n := range additions {
			if runA(&c1, base+n) {
				found = true
				fmt.Println("Solution:", base+n)
			}
			if c1.Output[0] == 2 &&
				c1.Output[1] == 4 &&
				c1.Output[2] == 1 &&
				c1.Output[3] == 1 &&
				c1.Output[4] == 7 &&
				c1.Output[5] == 5 &&
				c1.Output[6] == 0 &&
				c1.Output[7] == 3 &&
				c1.Output[8] == 4 &&
				c1.Output[9] == 3 &&
				c1.Output[10] == 1 &&
				c1.Output[11] == 6 &&
				c1.Output[12] == 5 &&
				c1.Output[13] == 5 &&
				c1.Output[14] == 3 &&
				c1.Output[15] == 0 {
				fmt.Println(":: ", base+n, format3(strconv.FormatInt(int64(base+n), 2)), c1.Output)
			}
		}

		if found || i > 10_000_000 {
			break
		}
		i++
		base = base + 0b1_000_000_000_000_000_000_000_000_000
	}

	fmt.Println("Part2:")
}

func runA(c1 *Computer, a int) bool {
	c1.A = a
	c1.B = 0
	c1.C = 0
	c1.Pointer = 0
	c1.Output = []int{}

	for {
		step(c1)
		if c1.Pointer >= len(c1.Program) {
			break
		}
	}
	return slices.Equal(c1.Output, c1.Program)
}

func format3(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i%3 == 0 && i != 0 {
			res = "_" + res
		}
		res = string(s[len(s)-i-1]) + res
	}
	return res
}

func part1() {
	cmp := loadInput()
	for {
		step(&cmp)
		if cmp.Pointer >= len(cmp.Program) {
			break
		}
	}
	ss := make([]string, len(cmp.Output))
	for i, n := range cmp.Output {
		ss[i] = strconv.Itoa(n)
	}
	fmt.Println("Part1:", strings.Join(ss, ","))
}

func step(cmp *Computer) {
	instr := cmp.Program[cmp.Pointer]
	operand := cmp.Program[cmp.Pointer+1]

	verbose := false
	if verbose {
		fmt.Printf("%s %d    A: %d  B: %d  C: %d\n", INames[instr], operand, cmp.A, cmp.B, cmp.C)
	}

	if instr == 0 {
		res := cmp.A >> comboOp(cmp, operand) //int(float64(cmp.A) / math.Pow(2.0, float64(comboOp(cmp, operand))))
		if verbose {
			fmt.Printf("  A = %d / 2**%d = %d\n", cmp.A, comboOp(cmp, operand), res)
		}
		cmp.A = res
		cmp.Pointer += 2
	}

	if instr == 1 {
		res := cmp.B ^ operand
		cmp.B = res
		cmp.Pointer += 2
	}

	if instr == 2 {
		res := comboOp(cmp, operand) % 8
		cmp.B = res
		cmp.Pointer += 2
	}

	if instr == 3 {
		if verbose {
			fmt.Printf("  A=%d Jump to %d\n", cmp.A, operand)
		}

		if cmp.A == 0 {
			cmp.Pointer += 2
		} else {
			cmp.Pointer = operand
		}
	}

	if instr == 4 {
		res := cmp.B ^ cmp.C
		cmp.B = res
		cmp.Pointer += 2
	}

	if instr == 5 {
		res := comboOp(cmp, operand) % 8
		if verbose {
			fmt.Printf("  Out: %d %% 8 = %d\n", comboOp(cmp, operand), res)
		}
		cmp.Output = append(cmp.Output, res)
		cmp.Pointer += 2
	}

	if instr == 6 {
		res := cmp.A >> comboOp(cmp, operand) //int(float64(cmp.A) / math.Pow(2.0, float64(comboOp(cmp, operand))))
		cmp.B = res
		cmp.Pointer += 2
	}

	if instr == 7 {
		res := cmp.A >> comboOp(cmp, operand) // int(float64(cmp.A) / math.Pow(2.0, float64(comboOp(cmp, operand))))
		cmp.C = res
		cmp.Pointer += 2
	}
}

func comboOp(cmp *Computer, v int) int {
	if v <= 3 {
		return v
	}
	if v == 4 {
		return cmp.A
	}
	if v == 5 {
		return cmp.B
	}
	if v == 6 {
		return cmp.C
	}
	panic("Wrong operand")
}

func loadInput() Computer {
	c1 := Computer{A: 18427963, B: 0, C: 0, Program: []int{2, 4, 1, 1, 7, 5, 0, 3, 4, 3, 1, 6, 5, 5, 3, 0}, Pointer: 0, Output: []int{}}
	// c1 := Computer{A: 729, B: 0, C: 0, Program: []int{0, 1, 5, 4, 3, 0}, Pointer: 0, Output: []int{}}
	// c1 := Computer{A: 117440, B: 0, C: 0, Program: []int{0, 3, 5, 4, 3, 0}, Pointer: 0, Output: []int{}}

	return c1
}
