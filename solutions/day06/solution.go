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

type base int

const (
	DOWN base = iota
	UP
	LEFT
	RIGHT
)

var offsets map[base]image.Point = map[base]image.Point{
	LEFT:  {0, -1},
	RIGHT: {0, 1},
	UP:    {-1, 0},
	DOWN:  {1, 0},
}

func parse(input string) map[image.Point]rune {
	grid := map[image.Point]rune{}
	for i, s := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, r := range s {
			grid[image.Point{i, j}] = r
		}
	}

	return grid
}

func findStartingPoint(grid map[image.Point]rune) image.Point {
	for pos, r := range grid {
		if r != '.' && r != '#' {
			return pos
		}
	}

	return image.Point{-1, -1}
}

func turnRight(orientation base) base {
	switch orientation {
	case DOWN:
		return LEFT
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case LEFT:
		return UP
	}

	return DOWN
}

func canGoInDirection(orientation base, current image.Point, grid map[image.Point]rune) bool {
	if r := grid[current.Add(offsets[orientation])]; r != '#' {
		return true
	}
	return false
}

func willGoOut(orientation base, current image.Point, grid map[image.Point]rune) bool {
	if _, ok := grid[current.Add(offsets[orientation])]; ok {
		return false
	}
	return true
}

func traverseTheGrid(current image.Point, orientation base, visited map[image.Point]bool, grid map[image.Point]rune) {
outside:
	for true {
		canGo := canGoInDirection(orientation, current, grid)

		if !canGo {
			orientation = turnRight(orientation)
		}

		for canGoInDirection(orientation, current, grid) {
			out := willGoOut(orientation, current, grid)

			if out {
				break outside
			}
			if _, ok := visited[current]; !ok {
				visited[current] = true
			}
			current = current.Add(offsets[orientation])
		}

	}
}

func part1(input string) int {
	grid := parse(input)
	current := findStartingPoint(grid)
	orientation := UP
	visited := make(map[image.Point]bool)

	traverseTheGrid(current, orientation, visited, grid)

	return len(visited) + 1
}

func pathToTheRightWasVisited(current image.Point, orientation base, grid map[image.Point]rune, visited map[image.Point]bool) bool {
	orientation = turnRight(orientation)

	for canGoInDirection(orientation, current, grid) && !willGoOut(orientation, current, grid) {
		current = current.Add(offsets[orientation])
	}

	if r, ok := grid[current.Add(offsets[orientation])]; ok && r == '#' && visited[current] {
		return true
	}

	return false
}

func traverseTheGridWithLoopCheck(current image.Point, orientation base, visited map[image.Point]int, grid map[image.Point]rune) bool {
	count := 0
	for true {
		canGo := canGoInDirection(orientation, current, grid)

		if !canGo {
			orientation = turnRight(orientation)
		}

		for canGoInDirection(orientation, current, grid) {
			out := willGoOut(orientation, current, grid)

			if out {
				return false
			}
			if count > 10000 {
				return true
			}

			visited[current]++
			count++
			current = current.Add(offsets[orientation])
		}

	}

	return false
}

func part2(input string) int {
	grid := parse(input)
	current := findStartingPoint(grid)
	orientation := UP
	visited := make(map[image.Point]int)
	count := 0

	for pos := range grid {
		cp := parse(input)
		if cp[pos] == '#' {
			continue
		} else {
			cp[pos] = '#'
		}

		if ok := traverseTheGridWithLoopCheck(current, orientation, visited, cp); ok {
			count++
		}
	}

	return count
}
