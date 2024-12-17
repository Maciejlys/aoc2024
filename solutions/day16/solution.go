package main

import (
	"container/heap"
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

func parse(input string) (map[image.Point]rune, image.Point, image.Point) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := map[image.Point]rune{}
	start, end := image.Point{}, image.Point{}

	for y, line := range lines {
		for x, r := range line {
			point := image.Point{x, y}
			grid[point] = r
			if r == 'S' {
				start = point
			} else if r == 'E' {
				end = point
			}
		}
	}

	return grid, start, end
}

type state struct {
	point     image.Point
	direction image.Point
	score     int
	path      []image.Point
}

type PriorityQueue []state

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(state))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func findBestPaths(input string) (int, [][]image.Point) {
	grid, start, end := parse(input)

	pq := &PriorityQueue{}
	heap.Init(pq)

	offsets := []image.Point{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	initialDirection := image.Point{1, 0}
	heap.Push(pq, state{start, initialDirection, 0, []image.Point{start}})

	visited := make(map[image.Point]map[image.Point]int)

	bestScore := -1
	var bestPaths [][]image.Point

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(state)

		if curr.point == end {
			if bestScore == -1 || curr.score < bestScore {
				bestScore = curr.score
				bestPaths = [][]image.Point{curr.path}
			} else if curr.score == bestScore {
				bestPaths = append(bestPaths, curr.path)
			}
			continue
		}

		if visited[curr.point] == nil {
			visited[curr.point] = make(map[image.Point]int)
		}
		if prevScore, ok := visited[curr.point][curr.direction]; ok && prevScore < curr.score {
			continue
		}
		visited[curr.point][curr.direction] = curr.score

		for _, offset := range offsets {
			nextPoint := curr.point.Add(offset)

			if r, ok := grid[nextPoint]; !ok || r == '#' {
				continue
			}

			additionalCost := 1
			if curr.direction != offset {
				additionalCost = 1001
			}

			newPath := append([]image.Point{}, curr.path...)
			newPath = append(newPath, nextPoint)

			nextState := state{nextPoint, offset, curr.score + additionalCost, newPath}
			heap.Push(pq, nextState)
		}
	}

	return bestScore, bestPaths
}

func part1(input string) int {
	bestScore, _ := findBestPaths(input)
	return bestScore
}

func part2(input string) int {
	_, paths := findBestPaths(input)
	visited := make(map[image.Point]struct{})

	for _, path := range paths {
		for _, p := range path {
			visited[p] = struct{}{}
		}
	}

	return len(visited)
}
