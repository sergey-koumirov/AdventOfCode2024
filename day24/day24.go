package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Operation struct {
	Kind, Input1, Input2, Output string
	Flag                         bool
}

var Red = "\033[91m"
var Reset = "\033[0m"

func main() {
	fmt.Println("Day 24:")

	// part1()
	part2()
}

func part1() {
	ops, zz := loadInput()

	for _, z := range zz {
		ops[z].Flag = deepCalc(ops, z)
	}

	fmt.Println("Part1:", zToNum(ops, zz))
}

func part2() {
	ops, zz := loadInput()

	validate(ops, zz)
	analize(ops)

	switchGates(ops, "jmq", "z06") //jmq qrh gmh

	validate(ops, zz)
	analize(ops)

	switchGates(ops, "gmh", "z13") // qrh gmh

	validate(ops, zz)
	analize(ops)

	switchGates(ops, "qrh", "z38") // qrh

	validate(ops, zz)
	analize(ops)

	switchGates(ops, "rqf", "cbd") // kdt + rqf | mrj + cbd
	validate(ops, zz)
	analize(ops)

	// fmt.Println("\nz05")
	// deepPrint(ops, "z05", 0)
	// fmt.Println()
	// fmt.Println("\nz06")
	// deepPrint(ops, "z06", 0)

	result := []string{"jmq", "z06", "gmh", "z13", "qrh", "z38", "rqf", "cbd"}
	sort.Strings(result)

	fmt.Println()
	fmt.Println("Part2:", strings.Join(result, ","))
}

func analize(ops map[string]*Operation) {
	fmt.Println()
	index := 0
	for key := range ops {
		if !isValid(ops, key) {
			op := ops[key]
			op1 := ops[op.Input1]
			op2 := ops[op.Input2]
			fmt.Println(index, op.Output, "=", op.Input1, op.Kind, op.Input2, "(", op1.Kind, op1.Output, op2.Kind, op2.Output, ")")
			index++
		}
	}
}

func switchGates(ops map[string]*Operation, op1, op2 string) {
	temp := ops[op2]

	ops[op2] = ops[op1]
	ops[op2].Output = op2

	ops[op1] = temp
	ops[op1].Output = op1
}

func validate(ops map[string]*Operation, zz []string) {
	fmt.Println()

	for key := range ops {
		if key[0] == 'y' {
			ops[key].Flag = false
		}
		if key[0] == 'x' {
			ops[key].Flag = true
		}
	}
	for _, z := range zz {
		ops[z].Flag = deepCalc(ops, z)
	}
	fmt.Printf("x1 + y0 %b\n", zToNum(ops, zz))

	for key := range ops {
		if key[0] == 'y' {
			ops[key].Flag = true
		}
		if key[0] == 'x' {
			ops[key].Flag = false
		}
	}
	for _, z := range zz {
		ops[z].Flag = deepCalc(ops, z)
	}
	fmt.Printf("x0 + y1 %b\n", zToNum(ops, zz))

	for key := range ops {
		if key[0] == 'y' {
			ops[key].Flag = true
		}
		if key[0] == 'x' {
			ops[key].Flag = false
		}
	}
	ops["x00"].Flag = true
	for _, z := range zz {
		ops[z].Flag = deepCalc(ops, z)
	}
	fmt.Printf("111 + 1 %b\n", zToNum(ops, zz))

	for key := range ops {
		if key[0] == 'y' {
			ops[key].Flag = false
		}
		if key[0] == 'x' {
			ops[key].Flag = true
		}
	}
	ops["y00"].Flag = true
	for _, z := range zz {
		ops[z].Flag = deepCalc(ops, z)
	}
	fmt.Printf("1 + 111 %b\n", zToNum(ops, zz))
}

func isValid(ops map[string]*Operation, key string) bool {
	op := ops[key]

	if op.Kind == "CONST" {
		return true
	}

	if op.Output == "z45" {
		return true
	}

	if op.Output[0] == 'z' {
		return op.Kind == "XOR"
	}

	op1 := ops[op.Input1]
	op2 := ops[op.Input2]

	if op.Kind == "OR" && op1.Kind == "AND" && op2.Kind == "AND" {
		return true
	}

	if op.Kind == "AND" && op1.Kind == "CONST" && op2.Kind == "CONST" {
		return true
	}

	if op.Kind == "XOR" && op1.Kind == "CONST" && op2.Kind == "CONST" {
		return true
	}

	if op.Kind == "AND" && (op1.Kind == "AND" && op2.Kind == "XOR" || op1.Kind == "XOR" && op2.Kind == "AND") {
		return true
	}

	if op.Kind == "AND" && (op1.Kind == "OR" && op2.Kind == "XOR" || op1.Kind == "XOR" && op2.Kind == "OR") {
		return true
	}

	return false
}

func deepPrint(ops map[string]*Operation, z string, level int) {
	op := ops[z]

	op1 := ops[op.Input1]
	op2 := ops[op.Input2]

	ss := strings.Repeat(" ", level)

	if op.Kind == "CONST" {
		return
	}

	if isXY(op1) {
		fmt.Println(ss, op.Kind, op.Input1)
		deepPrint(ops, op.Input1, level+4)

		fmt.Println(ss, "   ", op.Input2)
		deepPrint(ops, op.Input2, level+4)
	} else if isXY(op2) {
		fmt.Println(ss, op.Kind, op.Input2)
		deepPrint(ops, op.Input2, level+4)

		fmt.Println(ss, "   ", op.Input1)
		deepPrint(ops, op.Input1, level+4)
	} else {
		fmt.Println(ss, op.Kind, op.Input1)
		deepPrint(ops, op.Input1, level+4)

		fmt.Println(ss, "   ", op.Input2)
		deepPrint(ops, op.Input2, level+4)
	}
}

func isXY(op *Operation) bool {
	if op.Kind == "CONST" {
		return false
	}
	return op.Kind == "XOR" && (op.Input1[0] == 'x' && op.Input2[0] == 'y' || op.Input1[0] == 'y' && op.Input2[0] == 'x')
}

func deepCalc(ops map[string]*Operation, z string) bool {
	opsz := ops[z]

	if opsz.Kind == "CONST" {
		return opsz.Flag
	}

	if opsz.Kind == "OR" {
		return deepCalc(ops, opsz.Input1) || deepCalc(ops, opsz.Input2)
	}

	if opsz.Kind == "AND" {
		return deepCalc(ops, opsz.Input1) && deepCalc(ops, opsz.Input2)
	}

	if opsz.Kind == "XOR" {
		return deepCalc(ops, opsz.Input1) != deepCalc(ops, opsz.Input2)
	}

	panic("Unknown operation")
}

func zToNum(ops map[string]*Operation, zz []string) int {
	res := 0

	for i := len(zz) - 1; i >= 0; i-- {
		if ops[zz[i]].Flag {
			res = (res << 1) | 1
		} else {
			res = res << 1
		}
	}

	return res
}

func loadInput() (map[string]*Operation, []string) {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	res := map[string]*Operation{}
	znn := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			if strings.ContainsRune(line, ':') {
				parts := strings.Split(line, ": ")
				res[parts[0]] = &Operation{
					Kind:   "CONST",
					Output: parts[0],
					Flag:   parts[1] == "1",
				}
			} else {
				parts := strings.Split(line, " ")

				res[parts[4]] = &Operation{
					Kind:   parts[1],
					Output: parts[4],
					Input1: parts[0],
					Input2: parts[2],
				}

				if parts[4][0] == 'z' {
					znn = append(znn, parts[4])
				}
			}
		}

	}

	sort.Strings(znn)

	return res, znn

}
