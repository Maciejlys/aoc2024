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

type Robot struct {
	P, V image.Point
}

func sgn(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}

func parse(input string) ([]Robot, map[image.Point]int) {
	area := image.Rectangle{image.Point{0, 0}, image.Point{101, 103}}
	robots, quads := []Robot{}, map[image.Point]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var r Robot
		fmt.Sscanf(s, "p=%d,%d v=%d,%d", &r.P.X, &r.P.Y, &r.V.X, &r.V.Y)
		robots = append(robots, r)
		r.P = r.P.Add(r.V.Mul(100)).Mod(area)
		quads[image.Point{sgn(r.P.X - area.Dx()/2), sgn(r.P.Y - area.Dy()/2)}]++
	}

	return robots, quads
}

func part1(input string) int {
	_, quads := parse(input)
	result := 1

	for quad, count := range quads {
		if quad.X != 0 && quad.Y != 0 {
			result *= count
		}
	}

	return result
}

func part2(input string) int {
	robots, _ := parse(input)
	area := image.Rectangle{image.Point{0, 0}, image.Point{101, 103}}

	for t := 1; ; t++ {
		seen := map[image.Point]struct{}{}
		for i := range robots {
			robots[i].P = robots[i].P.Add(robots[i].V).Mod(area)
			seen[robots[i].P] = struct{}{}
		}

		if len(seen) == len(robots) {
			return t
		}
	}
}
