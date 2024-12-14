package solution13

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

type Solution13 struct{}

type ClawMachine struct {
	button1 []int
	button2 []int
	prize   []int
}

func (s Solution13) PartA(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, 0)
}

func (s Solution13) PartB(lineIterator *utils.LineIterator) int {
	return runSolution(lineIterator, 10000000000000)
}

func runSolution(lineIterator *utils.LineIterator, targetModifier int) int {
	clawMachines := parseInput(lineIterator, targetModifier)
	cost := 0
	for _, clawMachine := range clawMachines {
		cost += calcCost(clawMachine)
	}
	return cost
}

func parseInput(lineIterator *utils.LineIterator, targetModifier int) []ClawMachine {
	clawMachines := []ClawMachine{}
	currentClawMachine := ClawMachine{}
	currentClawMachineLineIdx := 0
	for lineIterator.Next() {
		line := lineIterator.Value()
		if len(line) == 0 {
			clawMachines = append(clawMachines, currentClawMachine)
			currentClawMachine = ClawMachine{}
			currentClawMachineLineIdx = 0
			continue
		}
		numbersString := strings.Split(line, ": ")[1]
		xString := strings.Split(numbersString, ", ")[0]
		yString := strings.Split(numbersString, ", ")[1]
		x, err := strconv.Atoi(xString[2:])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(yString[2:])
		if err != nil {
			panic(err)
		}
		if currentClawMachineLineIdx == 0 {
			currentClawMachine.button1 = []int{x, y}
		} else if currentClawMachineLineIdx == 1 {
			currentClawMachine.button2 = []int{x, y}
		} else {
			currentClawMachine.prize = []int{x + targetModifier, y + targetModifier}
		}
		currentClawMachineLineIdx++
	}
	clawMachines = append(clawMachines, currentClawMachine)
	return clawMachines
}

func calcCost(clawMachine ClawMachine) int {
	// Number of iterations for each button is the solution of the system
	// equation 1: iterButton1 * b1X + iterButton2 * b2X = tX
	// equation 2: iterButton1 * b1Y + iterButton2 * b2Y = tY
	// Where:
	// iterButton1 = number of times button 1 is pressed
	// iterButton2 = number of times button 2 is pressed
	// b1X = X increment for button 1
	// b1Y = Y increment for button 1
	// b2X = X increment for button 2
	// b2Y = Y increment for button 2
	// tX = X coord of the prize
	// tY = Y coord of the prize
	// The result is considered valid only if the resulting numbers are integers
	tX := float64(clawMachine.prize[0])
	tY := float64(clawMachine.prize[1])
	b1X := float64(clawMachine.button1[0])
	b1Y := float64(clawMachine.button1[1])
	b2X := float64(clawMachine.button2[0])
	b2Y := float64(clawMachine.button2[1])
	iterButton1 := (b2Y*tX-b2X*tY)/(b1X*b2Y-b1Y*b2X)
	iterButton2 := (b1Y*tX-b1X*tY)/(b1Y*b2X-b1X*b2Y)
	if (math.Mod(iterButton1, 1) == 0 && math.Mod(iterButton2, 1) == 0) {
		return int(3*iterButton1 + iterButton2)	
	}
	return 0
}

func init() {
	solutions.RegisterSolution(13, Solution13{})
}
