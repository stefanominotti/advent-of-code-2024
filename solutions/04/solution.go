package solution04

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"strings"
)

type Solution04 struct{}

func directions() []int {
	return []int{1, 0, -1}
}

func xmas() []string {
	return []string{"X", "M", "A", "S"}
}

func (s Solution04) PartA(lineIterator *utils.LineIterator) int {
	return wordSearch(lineIterator, countXmasFromPosition)
}

func countXmasFromPosition(input [][]string, i int, j int) int {
	xmasCount := 0
	// Iterate over directions (0, 1 or -1) for both i and j
	for _, iDir := range directions() {
		for _, jDir := range directions() {
			// If increment is 0 for both i and j skip
			if iDir == 0 && jDir == 0 {
				continue
			}
			// Given the direction, for each XMAS word check whether the expected word is found
			for idx, expected := range xmas() {
				// Increment i and j in the given direction
				incrementedI := i+iDir*idx
				incrementedJ := j+jDir*idx
				// Position is out of bound
				if (isOutbound(input, incrementedI, incrementedJ)) {
					break
				}
				// Expected XMAS word not found
				if input[incrementedI][incrementedJ] != expected {
					break
				}
				if (idx == len(xmas()) - 1) {
					xmasCount += 1
				}
			}
		}
	}
	return xmasCount
}

type Diagonal [][]int

func leftRight() Diagonal {
	return [][]int{{-1, -1}, {1, 1}}
}
func rightLeft() Diagonal {
	return [][]int{{1, -1}, {-1, 1}}
}

func (s Solution04) PartB(lineIterator *utils.LineIterator) int {
	searchFunction := func(input [][]string, i int, j int) int {
		if (isMasDiagonal(input, i, j, leftRight()) && isMasDiagonal(input, i, j, rightLeft())) {
			return 1
		}
		return 0
	}
	return wordSearch(lineIterator, searchFunction)
}

func isMasDiagonal(input [][]string, i int, j int, diagonal Diagonal) bool {
	if input[i][j] != "A" {
		return false
	}

	startI := i+diagonal[0][0]
	startJ := j+diagonal[0][1]
	if (isOutbound(input, startI, startJ)) {
		return false
	}

	endI := i+diagonal[1][0]
	endJ := j+diagonal[1][1]
	if (isOutbound(input, endI, endJ)) {
		return false
	}

	if input[startI][startJ] == "M" && input[endI][endJ] == "S" {
		return true
	}
	
	return input[startI][startJ] == "S" && input[endI][endJ] == "M"
}

func isOutbound(input [][]string, i int, j int) bool {
	return i < 0 || j < 0 || i >= len(input) || j >= len(input[i])
}

func wordSearch(lineIterator *utils.LineIterator, searchFunc func([][]string, int, int) int) int {
	input := [][]string{}

	for lineIterator.Next() {
		line := lineIterator.Value()
		input = append(input, strings.Split(line, ""))
	}

	result := 0
	for i, row := range input {
		for j, _ := range row {
			result += searchFunc(input, i, j)
		}
	}
	return result
}

func init() {
	solutions.RegisterSolution(4, Solution04{})
}
