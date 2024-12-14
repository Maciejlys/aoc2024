package main

import (
	_ "embed"
	"fmt"
	"image"
	"log"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

type game struct {
	A     image.Point
	B     image.Point
	Prize image.Point
}

type games []game

func (this games) log() {
	for _, line := range this {
		log.Print(line)
	}
}

func parse(input string) games {
	games := make(games, 0)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		var a, b, prize image.Point
		fmt.Sscanf(s, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&a.X, &a.Y, &b.X, &b.Y, &prize.X, &prize.Y)
		games = append(games, game{A: a, B: b, Prize: prize})
	}

	return games
}

func (this game) tokensToWin(offset image.Point) int {
	prizeWithOffset := this.Prize.Add(offset)
	ap := (this.B.Y*prizeWithOffset.X - this.B.X*prizeWithOffset.Y) / (this.A.X*this.B.Y - this.A.Y*this.B.X)
	bp := (this.A.Y*prizeWithOffset.X - this.A.X*prizeWithOffset.Y) / (this.A.Y*this.B.X - this.A.X*this.B.Y)
	if this.A.Mul(ap).Add(this.B.Mul(bp)) == prizeWithOffset {
		return ap*3 + bp
	}
	return 0
}

func part1(input string) int {
	games := parse(input)
	result := 0

	for _, game := range games {
		result += game.tokensToWin(image.Point{0, 0})
	}

	return result
}

func part2(input string) int {
	games := parse(input)
	result := 0

	for _, game := range games {
		result += game.tokensToWin(image.Point{10000000000000, 10000000000000})
	}

	return result
}
