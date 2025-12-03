package main

import "strconv"

func day01() int {
	puzzle_data, err := ReadLines("data/day01.txt")
	if err != nil {
		panic(err)
	}

	var position = 50
	var zeroCount = 0

	for _, line := range puzzle_data {
		// Parse the line which will be [R/L][Number]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if line[:1] == "L" {
			value = value * -1
		} else if line[:1] != "R" {
			panic("Invalid line doesn't start with L or R: " + line)
		}
		position = (((position + value) % 100) + 100) % 100
		if position == 0 {
			zeroCount++
		}
	}
	return zeroCount
}
