package main

import (
	_ "embed"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

func parse(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func part1(input string) int {

	return 0
}

func part2(input string) int {
	return 0
}
