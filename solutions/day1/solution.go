package main

import (
	_ "embed"
	"math"
	"sort"
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

type lists struct {
	first  []int
	second []int
}

func getLists(input string) lists {
	parsed := parse(input)
	first := make([]int, 0)
	second := make([]int, 0)

	for _, n := range parsed {
		numbers := strings.Fields(n)
		firstN, _ := strconv.Atoi(numbers[0])
		secondN, _ := strconv.Atoi(numbers[1])

		first = append(first, firstN)
		second = append(second, secondN)
	}

	return lists{first: first, second: second}
}

func part1(input string) int {
	sum := 0
	lists := getLists(input)
	sort.Ints(lists.first)
	sort.Ints(lists.second)

	for i := 0; i < len(lists.first); i++ {
		sum += int(math.Abs(float64(lists.first[i]) - float64(lists.second[i])))
	}

	return sum
}

func part2(input string) int {
	sum := 0
	occurences := make(map[int]int)
	lists := getLists(input)

	for _, n := range lists.second {
		occurences[n]++
	}

	for _, n := range lists.first {
		sum += n * occurences[n]
	}

	return sum
}
