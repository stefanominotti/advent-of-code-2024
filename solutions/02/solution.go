package solution02

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"math"
	"strings"
)

type Solution02 struct{}

type Sorting int

const (
	Undefined Sorting = 0
	Asc       Sorting = 1
	Desc      Sorting = 2
)

func (s Solution02) PartA(lineIterator *utils.LineIterator) int {
	return countValidReports(lineIterator, isReportOk)
}

func (s Solution02) PartB(lineIterator *utils.LineIterator) int {
	isReportOk := func(report []int) bool {
		if isReportOk(report) {
			return true
		} else {
			for i := 0; i < len(report); i++ {
				reportFixed := make([]int, len(report))
				copy(reportFixed, report)
				reportFixed = append(reportFixed[:i], reportFixed[i+1:]...)
				if isReportOk(reportFixed) {
					return true
				}
			}
		}
		return false
	}

	return countValidReports(lineIterator, isReportOk)
}

func countValidReports(lineIterator *utils.LineIterator, isReportValidFunc func([]int) bool) int {
	result := 0
	for lineIterator.Next() {
		line := lineIterator.Value()
		report := utils.StringsToIntegers(strings.Split(line, " "))

		if isReportValidFunc(report) {
			result += 1
		}
	}
	return result
}

func isReportOk(report []int) bool {
	previousNumber := -1
	currentSorting := Undefined
	for _, number := range report {
		if previousNumber == -1 {
			previousNumber = number
			continue
		}

		delta := math.Abs(float64(number) - float64(previousNumber))
		if delta < 1 || delta > 3 {
			return false
		}

		isAsc := number > previousNumber
		if currentSorting == Undefined {
			if isAsc {
				currentSorting = Asc
			} else {
				currentSorting = Desc
			}
		} else if isAsc && currentSorting == Desc || !isAsc && currentSorting == Asc {
			return false
		}
		previousNumber = number
	}
	return true
}

func init() {
	solutions.RegisterSolution(2, Solution02{})
}
