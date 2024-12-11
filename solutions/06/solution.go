package solution06

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"strings"
)

type Solution06 struct{}

type Direction int

const (
	Top    Direction = 1
	Right  Direction = 2
	Bottom Direction = 3
	Left   Direction = 4
)

func (s Solution06) PartA(lineIterator *utils.LineIterator) int {
	resultComputingFunc := func(input [][]string, i int, j int, direction Direction, visited [][][]Direction) int {

		travelWithoutAddingObstacles(input, i, j, Top, visited)

		result := 0
		for _, row := range visited {
			for _, col := range row {
				if len(col) > 0 {
					result += 1
				}
			}
		}

		return result
	}

	return runSolution(lineIterator, resultComputingFunc)
}

func (s Solution06) PartB(lineIterator *utils.LineIterator) int {
	resultComputingFunc := func(input [][]string, i int, j int, direction Direction, visited [][][]Direction) int {
		// Build matrix to keep track of possible extra obstacles
		possibleObstacles := [][]bool{}
		for _, row := range input {
			possibleObstacles = append(possibleObstacles, make([]bool, len(row)))
		}

		travelAddingObstacles(input, i, j, Top, visited, possibleObstacles)

		result := 0
		for rowIdx, row := range possibleObstacles {
			for colIdx, col := range row {
				if col && (rowIdx != i || colIdx != j) {
					result += 1
				}
			}
		}

		return result
	}

	return runSolution(lineIterator, resultComputingFunc)
}

func runSolution(lineIterator *utils.LineIterator, resultComputingFunc func([][]string, int, int, Direction, [][][]Direction) int) int {
	input := [][]string{}
	visited := [][][]Direction{}

	lineIndex := 0
	i := 0
	j := 0
	// Parse input and build matrix to keep track of visited cells
	for lineIterator.Next() {
		line := lineIterator.Value()
		input = append(input, strings.Split(line, ""))
		visited = append(visited, make([][]Direction, len(line)))

		if idx := strings.Index(line, "^"); idx != -1 {
			i = lineIndex
			j = idx
		}
		lineIndex += 1
	}

	return resultComputingFunc(input, i, j, Top, visited)
}

func travelWithoutAddingObstacles(input [][]string, i int, j int, direction Direction, visited [][][]Direction) bool {
	return travel(input, i, j, direction, visited, false, nil)
}

func travelAddingObstacles(input [][]string, i int, j int, direction Direction, visited [][][]Direction, possibleObstacles [][]bool) bool {
	return travel(input, i, j, direction, visited, true, possibleObstacles)
}

func travel(input [][]string, i int, j int, direction Direction, visited [][][]Direction, shouldTryObstacles bool, possibleObstacles [][]bool) bool {
	directionMoves := getDirectionMoves(direction)
	var prevI, prevJ int
	// Start straight line travel
	for {
		// If combination of position and direction already visited, return true, it's a loop
		if isDirectionInSlice(visited[i][j], direction) {
			return true
		}

		prevI, prevJ = i, j
		// Compute next position
		i, j = i+directionMoves[0], j+directionMoves[1]

		// If out of the matrix (i.e. travel finished),
		// add move to the visited matrix and return false, it's not a loop
		if utils.IsOutbound(input, i, j) {
			visited[prevI][prevJ] = append(visited[prevI][prevJ], direction)
			return false
		}

		// If next position is obstacle, exit straight line loop
		// and recursively call the function with direction rotated
		if input[i][j] == "#" {
			break
		}

		// If should try adding obstacles (part B) and the next position
		// is not altready visited (obstacles can only be added at the start,
		// so only when a position has not been already visited), try adding obstacle
		// and call the travel function with shouldTryObstacles = false (part A) to
		// just detect if it will end up in a loop
		if shouldTryObstacles && len(visited[i][j]) == 0 {
			inputCopied := copyInput(input)
			inputCopied[i][j] = "#"
			possibleObstacles[i][j] = possibleObstacles[i][j] || travelWithoutAddingObstacles(inputCopied, prevI, prevJ, getNextDirection(direction), copyVisited(visited))
		}

		// Finally add the move to the visited matrix and go on
		visited[prevI][prevJ] = append(visited[prevI][prevJ], direction)
	}

	return travel(input, prevI, prevJ, getNextDirection(direction), visited, shouldTryObstacles, possibleObstacles)
}

func copyInput(input [][]string) [][]string {
	duplicate := make([][]string, len(input))
	for i := range input {
		duplicate[i] = make([]string, len(input[i]))
		copy(duplicate[i], input[i])
	}
	return duplicate
}

func copyVisited(visited [][][]Direction) [][][]Direction {
	duplicatedVisited := make([][][]Direction, len(visited))
	for i := range visited {
		duplicatedVisited[i] = make([][]Direction, len(visited[i]))
		for j := range visited[i] {
			duplicatedVisited[i][j] = make([]Direction, len(visited[i][j]))
			copy(duplicatedVisited[i][j], visited[i][j])
		}
	}
	return duplicatedVisited
}

func isDirectionInSlice(directions []Direction, direction Direction) bool {
	for _, dir := range directions {
		if direction == dir {
			return true
		}
	}
	return false
}

func getNextDirection(direction Direction) Direction {
	switch direction {
	case Top:
		return Right
	case Right:
		return Bottom
	case Bottom:
		return Left
	case Left:
		return Top
	}
	panic("Invalid direction")
}

func getDirectionMoves(direction Direction) []int {
	switch direction {
	case Top:
		return []int{-1, 0}
	case Right:
		return []int{0, 1}
	case Bottom:
		return []int{1, 0}
	case Left:
		return []int{0, -1}
	}
	panic("Invalid direction")
}

func init() {
	solutions.RegisterSolution(6, Solution06{})
}
