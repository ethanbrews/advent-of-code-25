package main

import (
	"strconv"
	"strings"
)

func isValid(id string) bool {
	// Compare the left half of the string to the right half.
	// Uneven strings are valid by definition.
	l := len(id)
	return l%2 != 0 || id[:l/2] != id[l/2:]
}

func day02() int {
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
			if !isValid(s) {
				total += v
			}
		}
	}

	return total
}
