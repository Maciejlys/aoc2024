package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

func findBestPath(start, end image.Point, grid map[image.Point]bool) int {
	queue, dist := []image.Point{start}, map[image.Point]int{start: 0}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p == end {
			return dist[p]
		}

		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			n := p.Add(d)
			if _, ok := dist[n]; !ok && grid[n] {
				queue, dist[n] = append(queue, n), dist[p]+1
			}
		}
	}

	return -1
}

func parse(input string, size int) (map[image.Point]bool, []image.Point) {
	bytes := []image.Point{}
	for _, s := range strings.Fields(string(input)) {
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		bytes = append(bytes, image.Point{x, y})
	}

	grid := map[image.Point]bool{}
	for y := range size + 1 {
		for x := range size + 1 {
			grid[image.Point{x, y}] = true
		}
	}

	return grid, bytes
}

func part1(input string, size, amount int) int {
	grid, bytes := parse(input, size)

	for i := 0; i <= amount; i++ {
		grid[bytes[i]] = false
	}

	return findBestPath(image.Point{0, 0}, image.Point{size, size}, grid)
}

func part2(input string, size int) string {
	grid, bytes := parse(input, size)

	for b := range bytes {
		grid[bytes[b]] = false
		if findBestPath(image.Point{0, 0}, image.Point{size, size}, grid) != -1 {
			continue
		}
		return fmt.Sprintf("%d,%d", bytes[b].X, bytes[b].Y)
	}

	return ""
}
