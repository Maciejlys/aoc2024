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

func generateCombinations(n int, symbols []string) []string {
	if n <= 0 {
		return []string{}
	}

	var results []string

	var generate func(current string, depth int)
	generate = func(current string, depth int) {
		if depth == n {
			results = append(results, current)
			return
		}
		for _, symbol := range symbols {
			generate(current+symbol, depth+1)
		}
	}

	generate("", 0)
	return results
}

func parse(input string) map[int][]int {
	nums := make(map[int][]int)

	return nums
}

func findTestCases(input string, symbols []string) int {

	sum := 0

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	for _, line := range lines {
		entry := strings.Split(line, ":")
		test, _ := strconv.Atoi(entry[0])
		nums := make([]int, 0)

		for _, field := range strings.Fields(entry[1]) {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}

		combinations := generateCombinations(len(nums)-1, symbols)

		for _, combination := range combinations {
			fields := strings.Split(combination, "")
			temp := nums[0]

			for i := 1; i < len(nums); i++ {
				if fields[i-1] == "+" {
					temp += nums[i]

				} else if fields[i-1] == "*" {
					temp *= nums[i]
				} else {
					tempStr := strconv.Itoa(temp)
					numStr := strconv.Itoa(nums[i])
					str := tempStr + numStr
					converted, _ := strconv.Atoi(str)
					temp = converted
				}
			}

			if temp == test {
				sum += test
				break
			}
		}
	}

	return sum
}

func part1(input string) int {
	return findTestCases(input, []string{"+", "*"})
}

func part2(input string) int {
	return findTestCases(input, []string{"+", "*", "|"})
}
