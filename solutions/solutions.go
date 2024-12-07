package solutions

import (
	"advent-of-code/utils"
	"fmt"
)

// Solution interface with PartA and PartB
type Solution interface {
	PartA(lineIterator *utils.LineIterator) int
	PartB(lineIterator *utils.LineIterator) int
}

// RunAll executes both PartA and PartB for all registered solutions.
func RunAll() {
	for i := 1; i <= 25; i++ {
		solution, exists := solutionRegistry[i]
		if !exists {
			continue
		}
		utils.RunSolution(i, solution.PartA, solution.PartB)
		fmt.Println()
	}
}

// RunSolution executes PartA and PartB for a specific solution.
func RunSolution(id int) {
	solution, exists := solutionRegistry[id]
	if !exists {
		panic(fmt.Sprintf("Solution %d not found", id))
	}

	utils.RunSolution(id, solution.PartA, solution.PartB)
}

// Registry of solutions
var solutionRegistry = map[int]Solution{}

// RegisterSolution adds a solution to the registry.
func RegisterSolution(id int, solution Solution) {
	solutionRegistry[id] = solution
}
