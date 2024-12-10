package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

func parse(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getReports(input string) [][]int {
	parsed := parse(input)
	reports := make([][]int, 0)

	for _, line := range parsed {
		fields := strings.Fields(line)
		report := make([]int, 0)
		for _, n := range fields {
			num, _ := strconv.Atoi(string(n))
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	return reports
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isSortedWithConstraint(arr []int, tolerateRemoval bool) bool {
	isValid := func(subArr []int) bool {
		increasing := true
		decreasing := true

		for i := 1; i < len(subArr); i++ {
			diff := subArr[i] - subArr[i-1]

			if diff > 3 || diff < -3 || diff == 0 {
				return false
			}

			if diff < 0 {
				increasing = false
			}
			if diff > 0 {
				decreasing = false
			}
		}

		return increasing || decreasing
	}

	if !tolerateRemoval {
		return isValid(arr)
	}

	if isValid(arr) {
		return true
	}

	for i := 0; i < len(arr); i++ {
		subArr := removeIndex(arr, i)
		if isValid(subArr) {
			return true
		}
	}

	return false
}

func part1(input string) int {
	sum := 0
	reports := getReports(input)

	for _, report := range reports {
		if ok := isSortedWithConstraint(report, false); ok {
			sum++
		}
	}

	return sum
}

func part2(input string) int {
	sum := 0
	reports := getReports(input)

	for _, report := range reports {
		if ok := isSortedWithConstraint(report, true); ok {
			sum++
		}
	}

	return sum
}
