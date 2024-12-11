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
	return strings.Split(strings.TrimRight(input, "\n"), " ")
}

func run(stones []string, times int) int {
	stoneCounts := make(map[string]int)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for range times {
		newStoneCounts := make(map[string]int)

		for stone, count := range stoneCounts {
			if stone == "0" {
				newStoneCounts["1"] += count
			} else if len(stone)%2 == 0 {
				mid := len(stone) / 2
				left := stone[:mid]
				right := strings.TrimLeft(stone[mid:], "0")
				if right == "" {
					right = "0"
				}
				newStoneCounts[left] += count
				newStoneCounts[right] += count
			} else {
				num, _ := strconv.Atoi(stone)
				transformed := strconv.Itoa(num * 2024)
				newStoneCounts[transformed] += count
			}
		}

		stoneCounts = newStoneCounts
	}

	totalStones := 0
	for _, count := range stoneCounts {
		totalStones += count
	}
	return totalStones
}

func part1(input string) int {
	stones := parse(input)
	return run(stones, 25)
}

func part2(input string) int {
	stones := parse(input)
	return run(stones, 75)
}
