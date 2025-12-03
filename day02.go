package main

import (
	"strconv"
	"strings"
)

func isValidPart1(id string) bool {
	// Compare the left half of the string to the right half.
	// Uneven strings are valid by definition.
	l := len(id)
	return l%2 != 0 || id[:l/2] != id[l/2:]
}

func isValidPart2(id string) bool {
	return true
}

func day02Part1() int {
	return day02(isValidPart1)
}

func day02Part2() int {
	return day02(isValidPart2)
}

func day02(validator func(string) bool) int {
	puzzleData, err := ReadLines("data/day02.txt", ",")
	if err != nil {
		panic(err)
	}

	var total = 0

	for _, line := range puzzleData {
		split := strings.Split(line, "-")
		a, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		for v := a; v <= b; v++ {
			s := strconv.Itoa(v)
			if !validator(s) {
				total += v
			}
		}
	}

	return total
}
