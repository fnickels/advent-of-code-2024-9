package main

import "fmt"

// displayGrid prints the grid to the standard output.
// Each element of the grid is printed as a character, and each line of the grid is printed on a new line.
//
// Parameters:
//
//	grid (Grid): The grid to be displayed.
//
// Example:
//
//	displayGrid(myGrid)

func displayGrid(grid Grid) {
	for _, line := range grid {
		for _, v := range line {
			fmt.Printf("%c", v)
		}
		fmt.Printf("\n")
	}
}

// gridSize returns the dimensions of the grid as a pair of integers (rows, columns).
// If the grid is empty, it returns (0, 0).
//
// Parameters:
//
//	grid (Grid): The grid whose size is to be determined.
//
// Returns:
//
//	int: The number of rows in the grid.
//	int: The number of columns in the grid.
//
// Example:
//
//	rows, cols := gridSize(myGrid)

func gridSize(grid Grid) (int, int) {
	if len(grid) == 0 {
		return 0, 0
	}
	return len(grid), len(grid[0])
}
