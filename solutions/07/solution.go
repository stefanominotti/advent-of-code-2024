package solution07

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

type Solution07 struct{}

type Operator string

const (
	Add      Operator = "+"
	Multiply Operator = "*"
	Concat   Operator = "||"
	None     Operator = ""
)

func (s Solution07) PartA(lineIterator *utils.LineIterator) any {
	return runSolution(lineIterator, getNextOperatorWithoutConcat)
}

func (s Solution07) PartB(lineIterator *utils.LineIterator) any {
	return runSolution(lineIterator, getNextOperatorWithConcat)
}

func runSolution(lineIterator *utils.LineIterator, getNextOperator func(Operator, int, int) Operator) int {
	result := 0
	for lineIterator.Next() {
		line := lineIterator.Value()
		parts := strings.Split(line, ": ")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		numbers := utils.StringsToIntegers(strings.Split(parts[1], " "))
		if _, isValid := calculateOperators(target, numbers, getNextOperator); isValid {
			result += target
		}
	}
	return result
}

func calculateOperators(target int, numbers []int, getNextOperator func(Operator, int, int) Operator) ([]Operator, bool) {
	operators := make([]Operator, len(numbers)-1)
	// Set starting operators with the ones which gives the lower result
	for i := 0; i < len(operators); i++ {
		// If b = 1, a * b < a + b
		// || operator always give an higer result than + and *
		if numbers[i+1] == 1 {
			operators[i] = Multiply
		} else {
			operators[i] = Add
		}
	}
	partials := make([]int, len(operators))
	partials[0] = numbers[0]
	return calculateOperatorsRec(target, numbers, operators, getNextOperator, partials, 0)
}

func calculateOperatorsRec(target int, numbers []int, operators []Operator, getNextOperator func(Operator, int, int) Operator, partials []int, partialsValid int) ([]Operator, bool) {
	// Calculate result by sequentially applying operators starting after the
	// partial which are not valid (i.e. the partials after the operator change)
	// At the same time build an array of the next operators to try
	// i.e. operators which gives the next higher result than the current ones
	// And also keep tracks of the partial results
	result := partials[partialsValid]
	nextOperators := make([]Operator, len(operators))
	for idx, operator := range operators {
		if partialsValid <= idx {
			nextOperators[idx] = getNextOperator(operator, result, numbers[idx+1])
			result = applyOperator(result, numbers[idx+1], operator)
			if idx+1 < len(partials) {
				partials[idx+1] = result
			}
		} else {
			result = partials[idx+1]
		}
	}

	// If target is matched return
	if result == target {
		return operators, true
	}

	// If current result is greater than target then the current branch should be stopped
	// since the next iterations will only lead to higer results
	if result > target {
		return nil, false
	}

	// For each operator after valid index, branch and try the next one
	for idx := partialsValid; idx < len(operators); idx++ {
		// If we already tried the "highest rank" operator skip this branch
		nextOperator := nextOperators[idx]
		if nextOperator == None {
			continue
		}

		// Update operators list with next operator
		operatorsCopy := make([]Operator, len(operators))
		copy(operatorsCopy, operators)
		operatorsCopy[idx] = nextOperator

		partialsCopy := make([]int, len(partials))
		copy(partialsCopy, partials)

		// Try the branch by recursively call the function and if the branch is valid return
		// Keep valid the partial results before the operator that is changed
		validOperators, isValid := calculateOperatorsRec(target, numbers, operatorsCopy, getNextOperator, partialsCopy, idx)
		if isValid {
			return validOperators, true
		}
	}
	return nil, false
}

func concatOperators(operators []Operator) string {
	result := ""
	for _, operator := range operators {
		result += string(operator)
	}
	return result
}

func getNextOperatorWithoutConcat(current Operator, _ int, b int) Operator {
	switch current {
	case Add:
		if b != 1 {
			return Multiply
		}
	case Multiply:
		if b == 1 {
			return Add
		}
	}
	return None
}

func getNextOperatorWithConcat(current Operator, a int, b int) Operator {
	concatResult := applyOperator(a, b, Concat)
	multiplyResult := applyOperator(a, b, Multiply)
	switch current {
	case Add:
		if b != 1 && concatResult >= multiplyResult {
			return Multiply
		}
		return Concat
	case Multiply:
		if b == 1 {
			return Add
		}
		if concatResult >= multiplyResult {
			return Concat
		}
	case Concat:
		if concatResult < multiplyResult {
			return Multiply
		}
	}
	return None
}

func applyOperator(a int, b int, operator Operator) int {
	switch operator {
	case Add:
		return a + b
	case Multiply:
		return a * b
	case Concat:
		return a*int(math.Pow10(len(strconv.Itoa(b)))) + b
	}
	panic("Invalid operator")
}

func init() {
	solutions.RegisterSolution(7, Solution07{})
}
