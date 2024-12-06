package solution01

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

type Solution01 struct{}

func (s Solution01) PartA(lineIterator *utils.LineIterator) int {
	leftList, rightList := buildLists(lineIterator)

	result := 0
	for i := 0; i < len(leftList); i++ {
		result += int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
	}
	return result
}

func (s Solution01) PartB(lineIterator *utils.LineIterator) int {
	leftList, rightList := buildLists(lineIterator)

	scoresMap := map[int]int{}
	for i := 0; i < len(rightList); i++ {
		currentNumber := rightList[i]
		_, ok := scoresMap[currentNumber]
		if !ok {
			scoresMap[currentNumber] = currentNumber
		} else {
			scoresMap[currentNumber] += currentNumber
		}
	}

	result := 0
	for i := 0; i < len(leftList); i++ {
		currentNumber := leftList[i]
		score, ok := scoresMap[currentNumber]
		if ok {
			result += score
		}
	}
	return result
}

func buildLists(lineIterator *utils.LineIterator) ([]int, []int) {
	leftList := []int{}
	rightList := []int{}

	for lineIterator.Next() {
		line := lineIterator.Value()
		numbersString := strings.Split(line, "   ")
		leftNumber, err := strconv.Atoi(numbersString[0])
		if err != nil {
			panic(err)
		}
		rightNumber, err := strconv.Atoi(numbersString[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}
	return leftList, rightList
}

func init() {
	solutions.RegisterSolution(1, Solution01{})
}
