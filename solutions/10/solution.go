package solution10

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type Solution10 struct{}

func (s Solution10) PartA(lineIterator *utils.LineIterator) int {
	calculateScore := func(input [][]int, i int, j int) int {
		heightsReachable := map[string]bool{}
		return calculateScore(input, i, j, heightsReachable)
	}
	return runSolution(lineIterator, calculateScore)
}

func (s Solution10) PartB(lineIterator *utils.LineIterator) int {
	calculateScore := func(input [][]int, i int, j int) int {
		return calculateScore(input, i, j, nil)
	}
	return runSolution(lineIterator, calculateScore)
}

func runSolution(lineIterator *utils.LineIterator, calculateScore func([][]int, int, int) int) int {
	input := [][]int{}

	for lineIterator.Next() {
		line := lineIterator.Value()
		input = append(input, utils.StringsToIntegers(strings.Split(line, "")))
	}
	result := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			// From each 0 calculate score 
			if input[i][j] == 0 {
				result += calculateScore(input, i, j)
			}
		}
	}
	return result
}

func calculateScore(input [][]int, i int, j int, heightsReachable map[string]bool) int {
	currentValue := input[i][j]
	// Destination reached
	if currentValue == 9 {
		posString := fmt.Sprintf("%d%d", i, j)
		// If should count the same destination twice
		if heightsReachable == nil {
			return 1
		}
		// If should not count the same destination twice, and it's already visited
		if _, isVisited := heightsReachable[posString]; isVisited {
			return 0
		}
		// If should not count the same destination twice, and it's not already visited
		heightsReachable[posString] = true
		return 1
	}

	result := 0
	// Calculate score for each direction
	for _, direction := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		nextI := i + direction[0]
		nextJ := j + direction[1]
		if !utils.IsOutbound(input, nextI, nextJ) && input[nextI][nextJ] == currentValue+1 {
			result += calculateScore(input, nextI, nextJ, heightsReachable)
		}
	}
	return result
}

func init() {
	solutions.RegisterSolution(10, Solution10{})
}
