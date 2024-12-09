package main

import "fmt"

// createGridMarks initializes a GridMarks structure with the specified height and width.
// Each element in the GridMarks structure is initialized to false.
//
// Parameters:
//   - height: The number of rows in the grid.
//   - width: The number of columns in the grid.
//
// Returns:
//   A GridMarks structure with the specified dimensions.

func createGridMarks(height, width int) GridMarks {
	marks := make(GridMarks, height)
	for i := range marks {
		marks[i] = make(GridMarkLine, width)
	}
	return marks
}

// displayMarks prints the GridMarks structure to the console.
// The grid is displayed with '|' characters as vertical boundaries and '-' characters as horizontal boundaries.
// A '#' character represents a true value, and a space represents a false value.
//
// Parameters:
//   - marks: The GridMarks structure to display.

func displayMarks(marks GridMarks) {

	fmt.Printf("%v\n", gridBoundary(len(marks[0])))
	for _, line := range marks {
		fmt.Printf("|")
		for _, v := range line {
			o := " "
			if v {
				o = "#"
			}
			fmt.Printf("%v", o)
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("%v\n", gridBoundary(len(marks[0])))
}

// gridMarkSize returns the dimensions of the GridMarks structure.
//
// Parameters:
//   - gridMark: The GridMarks structure to measure.
//
// Returns:
//   Two integers representing the height and width of the grid, respectively.

func gridMarkSize(gridMark GridMarks) (int, int) {
	if len(gridMark) == 0 {
		return 0, 0
	}
	return len(gridMark), len(gridMark[0])
}

// gridBoundary generates a string representing the top or bottom boundary of a grid.
// The boundary consists of a '+' character at the beginning and end, with '-' characters in between.
// The number of '-' characters is determined by the width parameter.
//
// Parameters:
//   - width: The number of '-' characters to include between the '+' characters.
//
// Returns:
//
//	A string representing the grid boundary.

func gridBoundary(width int) string {
	boundary := "+"
	for i := 0; i < width; i++ {
		boundary += "-"
	}
	boundary += "+"
	return boundary
}
