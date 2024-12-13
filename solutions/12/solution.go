package solution12

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"strings"
)

type Solution12 struct{}

func (s Solution12) PartA(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, computeRegionPerimeterAndArea)
}

func (s Solution12) PartB(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, computeRegionSidesAndArea)
}

func runSolution(lineIterator *utils.LineIterator, computeRegionMeasures func([][]string, int, int, string, [][]bool) (int, int)) int {
	input := [][]string{}
	visited := [][]bool{}

	for lineIterator.Next() {
		line := lineIterator.Value()
		input = append(input, strings.Split(line, ""))
		visited = append(visited, make([]bool, len(line)))
	}

	// Starting from each unvisited square, compute perimeter/sides and area of its region
	result := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if visited[i][j] {
				continue
			}
			perimeterMeasure, area := computeRegionMeasures(input, i, j, input[i][j], visited)
			result += perimeterMeasure * area
		}
	}

	return result
}

func computeRegionPerimeterAndArea(input [][]string, i int, j int, current string, visited [][]bool) (int, int) {
	visited[i][j] = true
	perimeter := 0
	// Each cell increase area by 1
	area := 1
	// For each adjacent cell
	for _, direction := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		nextI := i + direction[0]
		nextJ := j + direction[1]
		// If next cell is outbound or different region it's a side so increase perimeter by 1
		if utils.IsOutbound(input, nextI, nextJ) || input[nextI][nextJ] != current {
			perimeter += 1
		} else if visited[nextI][nextJ] { // If already visited ignore it
			continue
		} else { // Otherwise call recursively on next cell and sum measures
			nextPerimeter, nextArea := computeRegionPerimeterAndArea(input, nextI, nextJ, current, visited)
			perimeter += nextPerimeter
			area += nextArea
		}
	}

	return perimeter, area
}

func computeRegionSidesAndArea(input [][]string, i int, j int, current string, visited [][]bool) (int, int) {
	visited[i][j] = true

	// Each cell increase area by 1
	area := 1
	// For each diagonal check if it's a corner and if so increase sides count by 1
	// Number of corners == numbers of sides
	sides := 0
	for _, diagonal := range [][]int{{1, 1}, {-1, -1}, {-1, 1}, {1, -1}} {
		side1I, side1J := i+diagonal[0], j
		side2I, side2J := i, j+diagonal[1]
		cornerI, cornerJ := i+diagonal[0], j+diagonal[1]
		if !isSameRegion(input, side1I, side1J, current) && !isSameRegion(input, side2I, side2J, current) {
			sides += 1
		} else if isSameRegion(input, side1I, side1J, current) && isSameRegion(input, side2I, side2J, current) && !isSameRegion(input, cornerI, cornerJ, current) {
			sides += 1
		}
	}

	// For each adjacent cell
	for _, direction := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		nextI := i + direction[0]
		nextJ := j + direction[1]
		// If next cell is outbound or different region or already visited ignore it
		if utils.IsOutbound(input, nextI, nextJ) || visited[nextI][nextJ] || input[nextI][nextJ] != current {
			continue
		} else { // Otherwise call recursively on next cell and sum measures
			nextSides, nextArea := computeRegionSidesAndArea(input, nextI, nextJ, current, visited)
			sides += nextSides
			area += nextArea
		}
	}

	return sides, area
}

func isSameRegion(input [][]string, i int, j int, current string) bool {
	return !utils.IsOutbound(input, i, j) && input[i][j] == current
}

func init() {
	solutions.RegisterSolution(12, Solution12{})
}
