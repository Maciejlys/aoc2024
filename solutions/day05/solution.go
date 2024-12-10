package main

import (
	_ "embed"
	"slices"
	"sort"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

func parse(input string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	updates := make([][]int, 0)

	splitted := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, l := range strings.Split(strings.TrimSpace(splitted[0]), "\n") {
		nums := strings.Split(l, "|")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		rules[a] = append(rules[a], b)
	}

	for _, l := range strings.Split(strings.TrimSpace(splitted[1]), "\n") {
		temp := []int{}
		nums := strings.Split(l, ",")

		for _, r := range nums {
			num, _ := strconv.Atoi(r)
			temp = append(temp, num)
		}

		updates = append(updates, temp)
	}

	return rules, updates
}

func part1(input string) int {
	sum := 0
	rules, updates := parse(input)

	for _, line := range updates {
		keep := true

		for i, j := 0, 1; j < len(line); i, j = i+1, j+1 {
			if !slices.Contains(rules[line[i]], line[j]) {
				keep = false
			}
		}

		if keep {
			sum += line[len(line)>>1]
		}
	}

	return sum
}

func part2(input string) int {
	sum := 0
	rules, updates := parse(input)

	for _, line := range updates {
		for i, j := 0, 1; j < len(line); i, j = i+1, j+1 {
			if !slices.Contains(rules[line[i]], line[j]) {
				sort.Slice(line, func(i, j int) bool {
					return slices.Contains(rules[line[i]], line[j])
				})
				sum += line[len(line)>>1]
			}
		}
	}

	return sum
}
