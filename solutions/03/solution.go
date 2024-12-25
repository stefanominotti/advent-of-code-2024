package solution03

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"regexp"
	"strings"
)

type Solution03 struct{}

func (s Solution03) PartA(lineIterator *utils.LineIterator) any {
	result := 0
	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)

	for lineIterator.Next() {
		line := lineIterator.Value()
		for _, match := range regex.FindAllString(line, -1) {
			result += performMultiplication(match)
		}
	}
	return result
}

type DoOperation string

const (
	Do   DoOperation = "do()"
	Dont DoOperation = "don't()"
)

func (s Solution03) PartB(lineIterator *utils.LineIterator) any {
	result := 0
	regex := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	action := Do
	for lineIterator.Next() {
		line := lineIterator.Value()
		for _, match := range regex.FindAllString(line, -1) {
			if match == string(Do) {
				action = Do
			} else if match == string(Dont) {
				action = Dont
			} else if action == Do {
				result += performMultiplication(match)
			}
		}
	}
	return result
}

func performMultiplication(mul string) int {
	mul = strings.ReplaceAll(mul, "mul(", "")
	mul = strings.ReplaceAll(mul, ")", "")
	numbers := utils.StringsToIntegers(strings.Split(mul, ","))
	return numbers[0] * numbers[1]
}

func init() {
	solutions.RegisterSolution(3, Solution03{})
}
