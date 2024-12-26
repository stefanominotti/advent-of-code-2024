package solution19

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"strings"
)

type Solution19 struct{}

func (s Solution19) PartA(lineIterator *utils.LineIterator) any {
	patterns, targets := parseInput(lineIterator)
	result := 0
	for _, target := range targets {
		if isTargetPossible(patterns, target, 0) {
			result += 1
		}
	}
	return result
}

func (s Solution19) PartB(lineIterator *utils.LineIterator) any {
	patterns, targets := parseInput(lineIterator)
	result := 0
	visitedCombinations := map[string]int{}
	for _, target := range targets {
		result += countPossibleCombinations(patterns, target, 0, visitedCombinations)
	}
	return result
}

func isTargetPossible(patterns []string, target string, index int) bool {
	possible := false
	for _, pattern := range patterns {
		nextIndex := index + len(pattern)
		// If current index plus length of pattern is outbound the pattern is not valid
		if nextIndex > len(target) {
			continue
		}
		// If target slice between current index and length of pattern is equal to the pattern
		// and the next index is outbound, than the we found a possible combination for
		// the given target
		// Otherwise just go on with the next index to check if the target is feasible
		if target[index:nextIndex] == pattern {
			if nextIndex == len(target) {
				return true
			}
			possible = possible || isTargetPossible(patterns, target, nextIndex)
		}
	}
	return possible
}

func countPossibleCombinations(patterns []string, target string, index int, visitedCombinations map[string]int) int {
	// If remaining part of the target is already visited, just return the
	// number of combination known
	if visitedCombination, ok := visitedCombinations[target[index:]]; ok {
		return visitedCombination
	}
	combinations := 0
	for _, pattern := range patterns {
		nextIndex := index + len(pattern)
		// If current index plus length of pattern is outbound the pattern is not valid
		if nextIndex > len(target) {
			continue
		}
		// If target slice between current index and length of pattern is equal to the pattern
		// and the next index is outbound, than the we found a possible combination for
		// the given target
		// Otherwise just go on with the next index to keep counting the combinations
		if target[index:nextIndex] == pattern {
			if nextIndex == len(target) {
				combinations += 1
			} else {
				combinations += countPossibleCombinations(patterns, target, nextIndex, visitedCombinations)
			}
		}
	}
	// Store combinations for the current slice of target
	visitedCombinations[target[index:]] = combinations
	return combinations
}

func parseInput(lineIterator *utils.LineIterator) ([]string, []string) {
	lineIterator.Next()
	patterns := strings.Split(lineIterator.Value(), ", ")

	lineIterator.Next()

	targets := []string{}
	for lineIterator.Next() {
		targets = append(targets, lineIterator.Value())
	}

	return patterns, targets
}

func init() {
	solutions.RegisterSolution(19, Solution19{})
}
