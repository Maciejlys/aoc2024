package main

import (
	_ "embed"
	"image"
	"log"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

var offsets = []image.Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type grid map[image.Point]int

func parse(input string) grid {
	grid := grid{}

	for i, s := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, r := range s {
			num, _ := strconv.Atoi(string(r))
			pos := image.Point{i, j}
			grid[pos] = num
		}
	}

	return grid
}

func (this grid) log() {
	log.Print(this)
}

func (this grid) startingPoints() []image.Point {
	startingPoints := make([]image.Point, 0)
	for pos, num := range this {
		if num == 0 {
			startingPoints = append(startingPoints, pos)
		}

	}
	return startingPoints
}

func (this grid) dfs(curr image.Point, visited map[image.Point]bool) (score int) {
	if this[curr] == 9 {
		if visited[curr] {
			return 0
		} else if visited != nil {
			visited[curr] = true
		}
		return 1
	}

	for _, offset := range offsets {
		if n := curr.Add(offset); this[n] == this[curr]+1 {
			score += this.dfs(n, visited)
		}
	}
	return score
}

func part1(input string) int {
	grid := parse(input)
	starting := grid.startingPoints()
	score := 0

	for _, pos := range starting {
		score += grid.dfs(pos, make(map[image.Point]bool))
	}

	return score
}

func part2(input string) int {
	grid := parse(input)
	starting := grid.startingPoints()
	score := 0

	for _, pos := range starting {
		score += grid.dfs(pos, nil)
	}

	return score
}
