package main

import (
	_ "embed"
	"image"
	"log"
	"slices"
	"strings"
)

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

//go:embed example3.txt
var example3 string

//go:embed input.txt
var actual string

var offsets = []image.Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type grid map[image.Point]rune
type group []image.Point
type groups [][]image.Point

func (this grid) log() {
	log.Print(this)
}

func (this groups) log() {
	for _, line := range this {
		log.Print(line)
	}
}

func parse(input string) grid {
	grid := make(grid)

	for i, s := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, r := range s {
			pos := image.Point{i, j}
			grid[pos] = r
		}
	}

	return grid
}

func (this grid) gatherGroups() groups {
	groups := make(groups, 0)
	visited := make(map[image.Point]struct{})

	var dfs func(current image.Point, group *group, r rune)
	dfs = func(current image.Point, group *group, r rune) {
		if _, ok := visited[current]; ok {
			return
		}
		if g, ok := this[current]; !ok || g != r {
			return
		}

		visited[current] = struct{}{}
		*group = append(*group, current)

		for _, offset := range offsets {
			dfs(current.Add(offset), group, r)
		}
	}

	for pos, r := range this {
		if _, ok := visited[pos]; ok {
			continue
		}
		group := make(group, 0)
		dfs(pos, &group, r)
		groups = append(groups, group)
	}

	return groups
}

func part1(input string) int {
	grid := parse(input)
	groups := grid.gatherGroups()
	result := 0

	for _, group := range groups {
		perimeter := 0

		for _, node := range group {
			for _, offset := range offsets {
				if !slices.Contains(group, node.Add(offset)) {
					perimeter++
				}
			}
		}

		result += perimeter * len(group)
	}

	return result
}

func part2(input string) int {
	grid := parse(input)
	groups := grid.gatherGroups()
	result := 0

	for _, group := range groups {
		sides := 0

		for _, node := range group {
			for _, offset := range offsets {
				if !slices.Contains(group, node.Add(offset)) {
					r := node.Add(image.Point{-offset.Y, offset.X})
					if grid[r] != grid[node] || grid[r.Add(offset)] == grid[node] {
						sides++
					}
				}
			}
		}

		result += sides * len(group)
	}

	return result
}
