package solution15

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"strings"
)

type Solution15 struct{}

func (s Solution15) PartA(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, false, "O", moveSingleBox)
}

func (s Solution15) PartB(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, true, "[", moveDoubleBox)
}

func runSolution(lineIterator *utils.LineIterator, doubleObjects bool, scoreSymbol string, moveBox func([][]string, [2]int, int, int) [][]string) int {
	warehouse, moves := parseInput(lineIterator, doubleObjects)
	startI, startJ := 0, 0
	for i, row := range warehouse {
		for j, col := range row {
			if col == "@" {
				startI, startJ = i, j
			}
		}
	}
	warehouse = moveRobot(warehouse, moves, 0, startI, startJ, moveBox)
	result := 0
	for i := range warehouse {
		for j, col := range warehouse[i] {
			if col == scoreSymbol {
				result += i*100 + j
			}
		}
	}
	return result
}

func parseInput(lineIterator *utils.LineIterator, doubleObjects bool) ([][]string, [][2]int) {
	warehouse := [][]string{}
	moves := [][2]int{}
	isMovesSection := false
	for lineIterator.Next() {
		line := lineIterator.Value()
		if len(line) == 0 {
			isMovesSection = true
			continue
		}
		lineSplit := strings.Split(line, "")
		if !isMovesSection {
			if !doubleObjects {
				warehouse = append(warehouse, lineSplit)
			} else {
				row := []string{}
				for _, object := range lineSplit {
					var doubledObject []string
					if object == "#" {
						doubledObject = []string{"#", "#"}
					} else if object == "O" {
						doubledObject = []string{"[", "]"}
					} else if object == "." {
						doubledObject = []string{".", "."}
					} else {
						doubledObject = []string{"@", "."}
					}
					row = append(row, doubledObject...)
				}
				warehouse = append(warehouse, row)
			}
		} else {
			for _, move := range lineSplit {
				if move == "^" {
					moves = append(moves, [2]int{-1, 0})
				} else if move == ">" {
					moves = append(moves, [2]int{0, 1})
				} else if move == "<" {
					moves = append(moves, [2]int{0, -1})
				} else {
					moves = append(moves, [2]int{1, 0})
				}
			}
		}
	}
	return warehouse, moves
}

func moveRobot(warehouse [][]string, moves [][2]int, currentMoveIdx int, i int, j int, moveBox func([][]string, [2]int, int, int) [][]string) [][]string {
	currentMove := moves[currentMoveIdx]

	// Apply move to current position
	nextI, nextJ := i+currentMove[0], j+currentMove[1]

	// If next position is a box try moving it
	if warehouse[nextI][nextJ] == "O" || warehouse[nextI][nextJ] == "[" || warehouse[nextI][nextJ] == "]" {
		warehouse = moveBox(warehouse, currentMove, nextI, nextJ)
	}

	// If next position is now empty, move the robot
	if warehouse[nextI][nextJ] == "." {
		warehouse[i][j] = "."
		warehouse[nextI][nextJ] = "@"
	} else { // Otherwise restore previous position
		nextI, nextJ = i, j
	}

	// Recursively call until all moves have been performed
	nextMove := currentMoveIdx + 1
	if nextMove != len(moves) {
		return moveRobot(warehouse, moves, nextMove, nextI, nextJ, moveBox)
	}
	return warehouse
}

// Moves a single box "O", or "[]" only when coming from left or right
func moveSingleBox(warehouse [][]string, move [2]int, i int, j int) [][]string {
	// Apply move
	nextI, nextJ := i+move[0], j+move[1]

	// If next position is edge, box cannot be moved
	if warehouse[nextI][nextJ] == "#" {
		return warehouse
	}

	// If next position is another box, try moving it
	if warehouse[nextI][nextJ] == "O" || warehouse[nextI][nextJ] == "[" || warehouse[nextI][nextJ] == "]" {
		warehouse = moveSingleBox(warehouse, move, nextI, nextJ)
	}

	// If next box has been moved, move the current one
	if warehouse[nextI][nextJ] == "." {
		warehouse[nextI][nextJ] = warehouse[i][j]
		warehouse[i][j] = "."
	}

	return warehouse
}

// Moves a double box "[]"
func moveDoubleBox(warehouse [][]string, move [2]int, i int, j int) [][]string {
	// If we are moving toward the box from left or right its movement is the same
	// as the single box case ("O")
	if move[0] == 0 {
		return moveSingleBox(warehouse, move, i, j)
	}

	// Get the coordinates of both box cells "[" and "]"
	pairI := i
	var pairJ int
	if warehouse[i][j] == "[" {
		pairJ = j + 1
	} else {
		pairJ = j - 1
	}

	// Apply the move to the two box cells and if at least one next cell is "#"
	// the box cannot be moved
	nextI, nextJ := i+move[0], j+move[1]
	if warehouse[nextI][nextJ] == "#" {
		return warehouse
	}
	nextPairI, nextPairJ := pairI+move[0], pairJ+move[1]
	if warehouse[nextPairI][nextPairJ] == "#" {
		return warehouse
	}

	// Copy the warehouse matrix and if in the next two cells there is a box try moving it
	warehouseCopy := utils.CopyMatrix(warehouse)
	if warehouseCopy[nextI][nextJ] == "[" || warehouseCopy[nextI][nextJ] == "]" {
		warehouseCopy = moveDoubleBox(warehouseCopy, move, nextI, nextJ)
	}
	if warehouseCopy[nextPairI][nextPairJ] == "[" || warehouseCopy[nextPairI][nextPairJ] == "]" {
		warehouseCopy = moveDoubleBox(warehouseCopy, move, nextPairI, nextPairJ)
	}

	// If both next cells are now empty, move the current box
	// Otherwise, restore the warehouse to avoid inconsistencies, 
	// as it's possible that only one of the two boxes in the next cells was moved.
	// Both boxes must move together or not at all
	if warehouseCopy[nextI][nextJ] == "." && warehouseCopy[nextPairI][nextPairJ] == "." {
		warehouseCopy[nextI][nextJ] = warehouseCopy[i][j]
		warehouseCopy[i][j] = "."

		warehouseCopy[nextPairI][nextPairJ] = warehouseCopy[pairI][pairJ]
		warehouseCopy[pairI][pairJ] = "."
		warehouse = warehouseCopy
	}
	return warehouse
}

func init() {
	solutions.RegisterSolution(15, Solution15{})
}
