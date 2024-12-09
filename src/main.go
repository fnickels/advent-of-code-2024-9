package main

import (
	"fmt"
	"os"
	"strings"
)

type Result struct {
	part1 int
	part2 int
}

type GridLine []rune
type Grid []GridLine

type GridMarkLine []bool
type GridMarks []GridMarkLine

type Coordinate struct {
	x int
	y int
}

type CoordinateList []Coordinate
type Locations map[rune]CoordinateList

type Pair struct {
	used int
	free int
}

type VariableList []Pair
type Data struct {
	line      string
	item      int
	varList   VariableList
	numBlocks int
}

type DataSet []Data

type State struct {
	readType  int
	data      DataSet
	grid      Grid
	marks     GridMarks
	locations Locations
	height    int
	width     int
	maxWidth  int
}

// main is the entry point of the Game of Life application. It reads and validates the input,
// processes the locations according to the game rules, and displays the results.
func main() {

	// read & validate input
	state, err := readInput("input-file.txt")
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		os.Exit(1)
	}

	// process state
	result := process(state)

	// display results
	display(result)
}

func process(state State) Result {
	// create a new state to hold the results
	result := Result{}

	switch state.readType {
	case ReadLines:
		// show line by line of data
		for i, v := range state.data {
			fmt.Printf("(%d) Processing pt 1 %v\n", i, v.line)

			//result.part1 += part1(state)
		}

		fmt.Printf("\n")

		for i, v := range state.data {
			if i == 0 {
				break // remove after part 1 is complete
			}

			fmt.Printf("(%d) Processing pt 2 %v\n", i, v.varList)

			//result.part2 += part2(state)
		}

	case ReadGrid:

		gridHeight, gridWidth := gridSize(state.grid)

		// create shadow matrix to hold marks
		state.marks = createGridMarks(gridHeight, gridWidth)

		state.marks = make(GridMarks, state.height)
		for i := range state.marks {
			state.marks[i] = make([]bool, state.width)
		}

		displayGrid(state.grid)
		fmt.Printf("Hight: %d, Width: %d\n\n", gridHeight, gridWidth)

		// display location map
		for k, v := range state.locations {
			fmt.Printf("key %c has %v\n", k, v)
		}

		displayMarks(state.marks)

		// process Locations
		for i, v := range state.locations {
			fmt.Printf("(%d) Processing pt 1 Locations: %v\n", i, v)

			// result.part1 += part1(state)

		}

		displayMarks(state.marks)

		fmt.Printf("\n")

		// re-create shadow matrix to hold marks
		state.marks = createGridMarks(gridHeight, gridWidth)

		// process Locations
		for i, v := range state.locations {
			fmt.Printf("(%d) Processing pt 1 Locations: %v\n", i, v)

			// result.part2 += part2(state)
		}

		displayMarks(state.marks)

	case ReadList:
		// show line by line of data
		for i, v := range state.data {
			// fmt.Printf("(%d) Line       pt 1 %v\n", i, v.line)
			fmt.Printf("(%d) Processing pt 1 %v\n", i, v.varList)

			result.part1 += part1(v.varList, v.numBlocks)
		}

		fmt.Printf("\n")

		for i, v := range state.data {
			fmt.Printf("(%d) Processing pt 2 %v\n", i, v.varList)

			result.part2 += part2(v.varList, v.numBlocks)
		}
	}

	return result
}

func part1(list VariableList, blocks int) int {
	// displayMap(list)
	fmt.Printf("Blocks: %d\n", blocks)

	slots := make([]int, blocks)
	pos := 0
	for i, pair := range list {
		for j := 0; j < pair.used; j++ {
			slots[pos] = i
			pos++
		}
		if pair.free > 0 {
			for j := 0; j < pair.free; j++ {
				slots[pos] = -1
				pos++
			}
		}
	}

	fmt.Printf("Slots   : %d\n", slots)

	// reshuffle

	for {
		firstOpen := getFirstOpen(slots)
		lastNotOpen := lastNotOpen(slots)

		if firstOpen > lastNotOpen {
			break
		}

		slots[firstOpen] = slots[lastNotOpen]
		slots[lastNotOpen] = -1
	}

	fmt.Printf("Re-Slots: %d\n", slots)

	checksum := 0
	for i, v := range slots {
		if v != -1 {
			checksum += i * v
		}
	}

	return checksum
}

func part2(list VariableList, blocks int) int {
	// displayMap(list)
	fmt.Printf("Blocks: %d\n", blocks)

	slots := make([]int, blocks)
	pos := 0
	for i, pair := range list {
		for j := 0; j < pair.used; j++ {
			slots[pos] = i
			pos++
		}
		if pair.free > 0 {
			for j := 0; j < pair.free; j++ {
				slots[pos] = -1
				pos++
			}
		}
	}

	fmt.Printf("Slots   : %d\n", slots)

	// reshuffle

	filled := getFirstOpen(slots)
	tailFree := lastNotOpen(slots)

	for {
		firstOpen, blockLen := getFirstOpenBlock(slots, filled)

		// try to move last file to first open slot
		file, fileLength := lastNotOpenBlock(slots, tailFree)
		// last file moved skip files already looked at
		tailFree = file - 1

		//	fmt.Printf("First Open: %d, Block Len: %d, File: %d, File Length: %d\n", firstOpen, blockLen, file, fileLength)
		//	displaySlots(slots)

		if firstOpen == -1 || file == -1 {
			break
		}

		if firstOpen > file {
			break
		}

		for {
			if blockLen >= fileLength {
				// move file to open slot
				for i := 0; i < fileLength; i++ {
					slots[firstOpen+i] = slots[file+i]
					slots[file+i] = -1
				}
				// reset as new space may have opened
				filled = getFirstOpen(slots)
				// tailFree = lastNotOpen(slots)
				break
			}

			// get next free slot to see if the file will fit
			firstOpen, blockLen = getFirstOpenBlock(slots, firstOpen+blockLen)
			if firstOpen == -1 {
				// no free slots found move to next file
				break
			}
			if firstOpen > file { // done move backwards
				break
			}
		}
	}

	fmt.Printf("Re-Slots: %d\n", slots)

	checksum := 0
	for i, v := range slots {
		if v != -1 {
			checksum += i * v
		}
	}

	return checksum
}

func displaySlots(list []int) {
	for _, v := range list {
		if v == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", v)
		}
	}
	fmt.Printf("\n")
}

func displayMap(list VariableList) {
	for i, pair := range list {

		fmt.Printf("%v", repeat(fmt.Sprintf("%d", i), pair.used))
		if pair.free > 0 {
			fmt.Printf("%v", repeat(".", pair.free))
		}
	}
	fmt.Printf("\n")
}

func repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

func getFirstOpen(slots []int) int {
	for i, v := range slots {
		if v == -1 {
			return i
		}
	}
	return -1
}

func lastNotOpen(slots []int) int {
	for i := len(slots) - 1; i >= 0; i-- {
		if slots[i] != -1 {
			return i
		}
	}
	return -1
}

func getFirstOpenBlock(slots []int, startPos int) (int, int) {
	target := -1
	length := 0
	for i := startPos; i < len(slots); i++ {
		if slots[i] == -1 {
			if target == -1 {
				target = i
				length = 1
			} else {
				length++
			}
		} else {
			if target != -1 {
				break
			}
		}
	}
	return target, length
}

func lastNotOpenBlock(slots []int, tail int) (int, int) {
	fileid := -1
	target := -1
	length := 0
	for i := tail; i >= 0; i-- {
		if fileid == -1 {
			if slots[i] != -1 {
				fileid = slots[i]
				target = i
				length = 1
			}
		} else {
			if slots[i] == fileid {
				target--
				length++
			} else {
				break
			}
		}
	}
	return target, length
}
