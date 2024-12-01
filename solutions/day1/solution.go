package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

//go:embed input.txt
var actual string

func parse(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

const numbers = "0123456789"

func part1(input string) int {
	parsed := parse(input)
	var first string
	var second string
	var sum int

	for _, line := range parsed {
		for i := 0; i < len(line); i++ {
			if strings.ContainsAny(line[i:i+1], numbers) {
				first = line[i : i+1]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if strings.ContainsAny(line[i:i+1], numbers) {
				second = line[i : i+1]
				break
			}
		}
		number, err := strconv.Atoi(first + second)

		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}

	return sum
}

func part2(input string) int {
	parsed := parse(input)
	words := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var first string
	var second string
	var sum int

	for _, line := range parsed {
	outer:
		for i := 0; i < len(line); i++ {
			for k := range words {
				if strings.Contains(line[0:i+1], k) {
					first = words[k]
					break outer
				}
			}
			if strings.ContainsAny(line[i:i+1], numbers) {
				first = line[i : i+1]
				break
			}
		}
	outer2:
		for i := len(line) - 1; i >= 0; i-- {
			for k := range words {
				log.Print(line[i:len(line)])
				if strings.Contains(line[i:len(line)], k) {
					first = words[k]
					break outer2
				}
			}
			if strings.ContainsAny(line[i:i+1], numbers) {
				second = line[i : i+1]
				break
			}
		}
		log.Print("Second:", second)
		log.Print("in line:", line)

		log.Print("------------------:")
		number, err := strconv.Atoi(first + second)

		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	return sum
}
