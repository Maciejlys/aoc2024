package main

import (
	_ "embed"
	"image"
	"strings"
)

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

//go:embed input.txt
var actual string

var offsets = map[rune]image.Point{
	'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
	'[': {1, 0}, ']': {-1, 0},
}

func parse(input string) (map[image.Point]rune, image.Point, string) {
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	grid, robot := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(split[0]) {
		for x, r := range s {
			if r == '@' {
				robot = image.Point{x, y}
				r = '.'
			}
			grid[image.Point{x, y}] = r
		}
	}

	return grid, robot, strings.ReplaceAll(split[1], "\n", "")
}

func run(input string) int {
	grid, robot, moves := parse(input)

loop:
	for _, r := range moves {
		queue, boxes := []image.Point{robot}, map[image.Point]rune{}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if _, ok := boxes[p]; ok {
				continue
			}
			boxes[p] = grid[p]

			switch n := p.Add(offsets[r]); grid[p.Add(offsets[r])] {
			case '#':
				continue loop
			case '[', ']':
				queue = append(queue, n.Add(offsets[grid[n]]))
				fallthrough
			case 'O':
				queue = append(queue, n)
			}
		}

		for b := range boxes {
			grid[b] = '.'
		}
		for b := range boxes {
			grid[b.Add(offsets[r])] = boxes[b]
		}
		robot = robot.Add(offsets[r])
	}

	gps := 0
	for p, r := range grid {
		if r == 'O' || r == '[' {
			gps += 100*p.Y + p.X
		}
	}
	return gps
}

func part1(input string) int {
	return run(input)
}

func part2(input string) int {
	replacer := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")
	return run(replacer.Replace(input))
}
