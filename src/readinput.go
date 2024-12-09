package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ReadLines int = iota
	ReadGrid  int = iota
	ReadList  int = iota
)

func readInput(fileName string) (State, error) {

	readType := ReadList

	state := State{
		readType:  readType,
		data:      DataSet{},
		grid:      Grid{},
		locations: Locations{},
	}

	file, err := os.Open(fileName)
	if err != nil {
		return state, fmt.Errorf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	// read in lines
	x := 0
	state.maxWidth = 0
	for scanner.Scan() {
		line := scanner.Text()
		// part := strings.Fields(line)
		part := []string{" "} //dummy

		switch readType {
		case ReadLines:

			if len(part) != 2 {
				return state,
					fmt.Errorf(
						"invalid input format, expected: '<num1> <num2>', found: '%s'",
						line,
					)
			}

			a, err := getInt64(part[0])
			if err != nil {
				return state, err
			}

			if len(part) > state.maxWidth {
				state.maxWidth = len(part)
			}
			state.width = len(part)

			state.data = append(state.data,
				Data{
					line: line,
					item: int(a),
				})

		case ReadGrid:

			gridLine := GridLine{}

			for y, char := range line {
				gridLine = append(gridLine, char)

				// capture list of coordinates for significant elements
				if char != '.' {
					if coord, ok := state.locations[char]; !ok {
						state.locations[char] = CoordinateList{
							Coordinate{x, y},
						}
					} else {
						state.locations[char] = append(coord, Coordinate{x, y})
					}
				}

			}

			state.grid = append(state.grid, gridLine)

			if len(line) > state.maxWidth {
				state.maxWidth = len(line)
			}
			state.width = len(line)

			state.data = append(state.data,
				Data{
					line: line,
				})

		case ReadList:

			// ******************************
			// *** EXTRA PROCESSING LOGIC ***
			// ******************************
			// if !strings.HasSuffix(parts[0], ":") {
			// 	return state, fmt.Errorf("invalid input format, expected: '<num1>: <num2> <num3> <num...>', found: '%s'", line)
			// }

			// parts[0] = strings.TrimSuffix(parts[0], ":")
			// result, err := strconv.ParseInt(parts[0], 10, 64)
			// if err != nil {
			// 	return state, fmt.Errorf("invalid result value. must be valid 64-bit signed integers. found: '%s'", parts[0])
			// }
			// ******************************

			vars := VariableList{}

			sum := 0
			for i := 0; i < len(line); i += 2 {
				used := int(line[i]) - int('0')
				sum += used
				free := -1
				if i+1 < len(line) {
					free = int(line[i+1]) - int('0')
					sum += free
				}

				vars = append(vars, Pair{int(used), int(free)})
			}

			state.width = len(line)
			state.data = append(state.data,
				Data{
					// line:    line,
					varList:   vars,
					numBlocks: sum,
				})

		default:
			return state, fmt.Errorf("invalid read type: %d", readType)
		}

		x++
	}

	state.height = x

	return state, nil
}
