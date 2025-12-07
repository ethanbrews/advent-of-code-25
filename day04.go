package main

import "fmt"

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%2d ", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func day04() int {
	puzzleData, err := ReadLines("data/day04.txt", "\n")
	if err != nil {
		panic(err)
	}

	var grid = make([][]int, len(puzzleData))
	grid[0] = make([]int, len(puzzleData[0]))

	for i := 0; i < len(grid); i++ {
		if i < len(grid)-1 {
			grid[i+1] = make([]int, len(puzzleData[0]))
		}
		for j := 0; j < len(grid[i]); j++ {
			if puzzleData[i][j] != '@' {
				continue
			}

			l := j > 0
			r := j < len(grid[i])-1
			u := i > 0
			d := i < len(puzzleData)-1

			if l {
				grid[i][j-1] += 1
			}
			if l && u {
				grid[i-1][j-1] += 1
			}
			if u {
				grid[i-1][j] += 1
			}
			if u && r {
				grid[i-1][j+1] += 1
			}
			if r {
				grid[i][j+1] += 1
			}
			if d && r {
				grid[i+1][j+1] += 1
			}
			if d {
				grid[i+1][j] += 1
			}
			if d && l {
				grid[i+1][j-1] += 1
			}
		}
	}

	total := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 4 && puzzleData[i][j] == '@' {
				total++
			}
		}
	}

	return total
}
