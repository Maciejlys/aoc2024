package main

import (
	_ "embed"
	"image"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

var offsets = []image.Point{
	{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

var diagOffsets []image.Point = []image.Point{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}

func parse(input string) map[image.Point]rune {
	grid := map[image.Point]rune{}
	for i, s := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, r := range s {
			grid[image.Point{i, j}] = r
		}
	}

	return grid
}

func part1(input string) int {
	grid := parse(input)
	count := 0

	for pos := range grid {
		for _, offset := range offsets {
			curr := ""

			for i := range 4 {
				curr += string(grid[pos.Add(offset.Mul(i))])
			}

			if curr == "XMAS" {
				count++
			}
		}
	}

	return count
}

func part2(input string) int {
	grid := parse(input)
	count := 0

	for pos, r := range grid {
		curr := ""
		if r == 'A' {
			for _, offset := range diagOffsets {
				curr += string(grid[pos.Add(offset)])
			}

			if curr == "MMSS" || curr == "SSMM" || curr == "SMSM" || curr == "MSMS" {
				count++
			}
		}
	}

	return count
}
