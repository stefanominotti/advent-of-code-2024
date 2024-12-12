package solution08

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"strings"
)

type Solution08 struct{}

func (s Solution08) PartA(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, addAntinodesPartA)
}

func (s Solution08) PartB(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, addAntinodesPartB)
}

func runSolution(lineIterator *utils.LineIterator, addAntinodes func([][]string, int, int, int, int, [][]bool)) int {
	input := [][]string{}
	antinodes := [][]bool{}

	for lineIterator.Next() {
		line := lineIterator.Value()
		input = append(input, strings.Split(line, ""))
		antinodes = append(antinodes, make([]bool, len(line)))
	}

	for i, row := range input {
		for j, col := range row {
			if col == "." {
				continue
			}
			for _, antenna := range searchSameFrequencyAntennas(input, i, j, col) {
				addAntinodes(input, i, j, antenna[0], antenna[1], antinodes)
			}
		}
	}

	result := 0
	for _, row := range antinodes {
		for _, col := range row {
			if col {
				result += 1
			}
		}
	}

	return result
}

func addAntinodesPartA(input [][]string, i1 int, j1 int, i2 int, j2 int, antinodes [][]bool) {
	iDelta := i1 - i2
	jDelta := j1 - j2

	topAntinodeI := i2 - iDelta
	topAntinodeJ := j2 - jDelta
	bottomAntinodeI := i1 + iDelta
	bottomAntinodeJ := j1 + jDelta

	if !utils.IsOutbound(input, topAntinodeI, topAntinodeJ) {
		antinodes[topAntinodeI][topAntinodeJ] = true
	}
	if !utils.IsOutbound(input, bottomAntinodeI, bottomAntinodeJ) {
		antinodes[bottomAntinodeI][bottomAntinodeJ] = true
	}
}

func addAntinodesPartB(input [][]string, i1 int, j1 int, i2 int, j2 int, antinodes [][]bool) {
	iDelta := i1 - i2
	jDelta := j1 - j2

	topAntinodeI := i2
	topAntinodeJ := j2
	for !utils.IsOutbound(input, topAntinodeI, topAntinodeJ) {
		antinodes[topAntinodeI][topAntinodeJ] = true
		topAntinodeI = topAntinodeI - iDelta
		topAntinodeJ = topAntinodeJ - jDelta
	}

	bottomAntinodeI := i1
	bottomAntinodeJ := j1
	for !utils.IsOutbound(input, bottomAntinodeI, bottomAntinodeJ) {
		antinodes[bottomAntinodeI][bottomAntinodeJ] = true
		bottomAntinodeI = topAntinodeI - iDelta
		bottomAntinodeJ = topAntinodeJ - jDelta
	}
}

func searchSameFrequencyAntennas(input [][]string, currI int, currJ int, frequency string) [][]int {
	result := [][]int{}
	for i, row := range input {
		for j, col := range row {
			if i == currI && j == currJ {
				continue
			}
			if col == frequency {
				coord := []int{i, j}
				result = append(result, coord)
			}
		}
	}
	return result
}

func init() {
	solutions.RegisterSolution(8, Solution08{})
}
