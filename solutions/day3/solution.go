package main

import (
	_ "embed"
	"regexp"
	"strconv"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

func part1(input string) int {
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		sum += first * second
	}

	return sum
}

func part2(input string) int {
	enabled := true
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't|do`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if match[0] == "do" {
			enabled = true
		} else if match[0] == "don't" {
			enabled = false
		} else if enabled {
			first, _ := strconv.Atoi(match[1])
			second, _ := strconv.Atoi(match[2])
			sum += first * second
		}
	}

	return sum
}
