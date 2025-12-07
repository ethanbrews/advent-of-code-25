package main

// makeAccessibilityGrid returns a grid of int, representing the number of adjacent rolls for a given square.
func makeAccessibilityGrid(puzzleData [][]byte) [][]int {
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

	return grid
}

// countAccessible counts the number of rolls which are accessible (fewer than 4 adjacent rolls)
func countAccessible(grid [][]int, puzzleData [][]byte) int {
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

// cleanAccessible removes any accessible rolls, mutating the puzzleData input.
func cleanAccessible(grid [][]int, puzzleData *[][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 4 && (*puzzleData)[i][j] == '@' {
				(*puzzleData)[i][j] = '.'
			}
		}
	}
}

// countRolls counts the number of rolls present in the puzzleData
func countRolls(puzzleData [][]byte) int {
	total := 0

	for i := 0; i < len(puzzleData); i++ {
		for j := 0; j < len(puzzleData[i]); j++ {
			if puzzleData[i][j] == '@' {
				total++
			}
		}
	}

	return total
}

// readData reads the day04 input and returns it as a 2d byte array
func readData() [][]byte {
	puzzleData, err := ReadLines("data/day04.txt", "\n")
	if err != nil {
		panic(err)
	}
	result := make([][]byte, len(puzzleData))
	for i, row := range puzzleData {
		result[i] = []byte(row)
	}
	return result
}

func day04Part1() int {
	puzzleData := readData()
	grid := makeAccessibilityGrid(puzzleData)
	return countAccessible(grid, puzzleData)
}

func day04Part2() int {
	puzzleData := readData()
	grid := makeAccessibilityGrid(puzzleData)
	initialCount := countRolls(puzzleData)
	for countAccessible(grid, puzzleData) > 0 {
		cleanAccessible(grid, &puzzleData)
		grid = makeAccessibilityGrid(puzzleData)
	}
	return initialCount - countRolls(puzzleData)
}
