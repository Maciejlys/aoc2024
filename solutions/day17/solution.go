package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

func parse(input string) (int, int, int, []int) {
	split := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	var a int
	var b int
	var c int

	fmt.Sscanf(split[0], "Register A: %d", &a)
	fmt.Sscanf(split[0], "Register B: %d", &b)
	fmt.Sscanf(split[0], "Register C: %d", &c)

	trimmed := strings.TrimLeft(strings.Split(split[1], ":")[1], " ")
	var program []int
	json.Unmarshal([]byte("["+trimmed+"]"), &program)

	return a, b, c, program
}

func run(a, b, c int, program []int) (out []int) {
	for ip := 0; ip < len(program); ip += 2 {
		op, literal := program[ip], program[ip+1]

		var combo int
		switch literal {
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		default:
			combo = literal
		}

		switch op {
		case 0:
			a >>= combo
		case 1:
			b ^= literal
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				ip = literal - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, combo%8)
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}
	}

	return out
}

func part1(input string) string {
	a, b, c, program := parse(input)
	output := make([]string, 0)
	result := run(a, b, c, program)

	for _, num := range result {
		str := strconv.Itoa(num)
		output = append(output, str)
	}

	return strings.Join(output, ",")
}

func part2(input string) int {
	a, b, c, program := parse(input)

	a = 0
	for n := len(program) - 1; n >= 0; n-- {
		a <<= 3
		for !slices.Equal(run(a, b, c, program), program[n:]) {
			a++
		}
	}
	return a
}
