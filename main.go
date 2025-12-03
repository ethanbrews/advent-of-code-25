package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var funcDayMap = [][]func() int{
	{day01},
	{day02Part1, day02Part2},
}

func main() {
	args := os.Args[1:]
	var day int
	if len(args) == 0 {
		date := time.Now()
		if date.Year() != 2025 || date.Day() > 12 || date.Month() != time.December {
			panic("Advent of code has ended!")
		}
		day = time.Now().Day()
	} else {
		var err error
		day, err = strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
	}

	for i, f := range funcDayMap[day-1] {
		fmt.Printf("Day %d part %d: %d\n", day, i+1, f())
	}
}
