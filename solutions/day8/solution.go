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

func parse(input string) (map[rune][]image.Point, int, int) {
	antenas := make(map[rune][]image.Point)
	maxI := 0
	maxJ := 0
	for i, s := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, r := range s {
			if r != '.' {
				antenas[r] = append(antenas[r], image.Point{i, j})
			}
			maxJ = j
		}
		maxI = i
	}
	return antenas, maxI, maxJ
}

func calculateAntinodesForPair(a image.Point, b image.Point) [2]image.Point {
	p1 := image.Point{
		X: b.X + 2*(a.X-b.X),
		Y: b.Y + 2*(a.Y-b.Y),
	}

	p2 := image.Point{
		X: a.X + 2*(b.X-a.X),
		Y: a.Y + 2*(b.Y-a.Y),
	}

	return [2]image.Point{p1, p2}
}

func calculateAntinodesPart1(points []image.Point, rec image.Rectangle) map[image.Point]struct{} {
	antinodes := make(map[image.Point]struct{})

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairAntinodes := calculateAntinodesForPair(points[i], points[j])
			if pairAntinodes[0].In(rec) {
				antinodes[pairAntinodes[0]] = struct{}{}
			}
			if pairAntinodes[1].In(rec) {
				antinodes[pairAntinodes[1]] = struct{}{}
			}
		}
	}

	return antinodes
}

func part1(input string) int {
	antenas, maxI, maxJ := parse(input)
	globalAntinodes := make(map[image.Point]struct{})

	rect := image.Rectangle{image.Point{0, 0}, image.Point{maxI + 1, maxJ + 1}}

	for _, points := range antenas {
		antinodes := calculateAntinodesPart1(points, rect)
		for antinode := range antinodes {
			globalAntinodes[antinode] = struct{}{}
		}
	}

	return len(globalAntinodes)
}

func pointsOnLine(p1, p2 image.Point, rect image.Rectangle) []image.Point {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	points := []image.Point{}

	current := p1
	for current.In(rect) {
		points = append(points, current)
		current = image.Point{X: current.X - dx, Y: current.Y - dy}
	}

	current = p2
	for current.In(rect) {
		points = append(points, current)
		current = image.Point{X: current.X + dx, Y: current.Y + dy}
	}

	return points
}

func part2(input string) int {
	antenas, maxI, maxJ := parse(input)
	globalAntinodes := make(map[image.Point]struct{})

	rect := image.Rectangle{image.Point{0, 0}, image.Point{maxI + 1, maxJ + 1}}

	for _, points := range antenas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				onLine := pointsOnLine(points[i], points[j], rect)
				for _, on := range onLine {
					globalAntinodes[on] = struct{}{}
				}
			}
		}
	}

	return len(globalAntinodes)
}
