package main

import (
	_ "embed"
	"strconv"
)

//go:embed example.txt
var example string

//go:embed input.txt
var actual string

func parse(input string) []int {
	str := make([]int, 0)

	id := 0
	for i, r := range input {
		toAppend := -1
		if i%2 == 0 {
			toAppend = id
			id++
		}
		amount, _ := strconv.Atoi(string(r))
		for range amount {
			str = append(str, toAppend)
		}
	}

	return str
}

func getLastFileIndex(files []int) int {
	for i := len(files) - 1; i >= 0; i-- {
		if files[i] != -1 {
			return i
		}
	}

	return -1
}

func part1(input string) int {
	files := parse(input)

	checksum := 0

	for i, file := range files {
		index := getLastFileIndex(files)
		if file == -1 && index >= i {
			files[i], files[index] = files[index], files[i]
		}
		if files[i] != -1 {
			checksum += files[i] * i
		}
	}

	return checksum
}
