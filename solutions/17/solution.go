package solution17

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

type Solution17 struct{}

func (s Solution17) PartA(lineIterator *utils.LineIterator) any {
	registers, program := parseInput(lineIterator)
	output := runProgram(registers, program)
	result := ""
	for _, num := range output {
		result += strconv.Itoa(num)
		result += ","
	}
	return result[:len(result)-1]
}

func (s Solution17) PartB(lineIterator *utils.LineIterator) any {
	_, program := parseInput(lineIterator)
	// Init counter to the min number for which the lenght of the output is equal
	// to the length of the program, i.e. 8^(program_length-1)
	counter := int(math.Pow(8, float64(len(program) - 1)))
	registers := [3]int{counter, 0, 0}
	output := runProgram(registers, program)

	// For each number of the output (starting from the last one)
	// increase the counter by 8^number_index until the nuber match
	// the program number for that index
	for i := len(program) - 1; i >= 0; i-- {
		for program[i] - output[i] != 0 {
			counter += int(math.Pow(8, float64(i)))
			registers[0] = counter
			output = runProgram(registers, program)
		}
	}
	return counter
}

func runProgram (registers [3]int, program []int) []int {
	instructionPointer := 0
	output := []int{}
	for instructionPointer < len(program) {
		updatedRegisters, nextPointer, isOutput, instructionOutput := runNextInstruction(program, registers, instructionPointer)
		registers = updatedRegisters
		instructionPointer = nextPointer
		if isOutput {
			output = append(output, instructionOutput)
		}
	}
	return output
}

func runNextInstruction(program []int, registers [3]int, instructionPointer int) ([3]int, int, bool, int) {
	instruction := program[instructionPointer]
	operand := program[instructionPointer + 1]
	return applyInstruction(instruction, operand, registers, instructionPointer)
}

func parseInput(lineIterator *utils.LineIterator) ([3]int, []int) {
	registers := [3]int{}
	for idx := range 3 {
		lineIterator.Next()
		registers[idx] = parseRegisterLine(lineIterator.Value())
	}
	lineIterator.Next()
	lineIterator.Next()
	programString := strings.Split(lineIterator.Value(), ": ")[1]
	program := utils.StringsToIntegers(strings.Split(programString, ","))
	return registers, program
}

func parseRegisterLine(line string) int {
	numberString := strings.Split(line, ": ")[1]
	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}
	return number
}

func getComboOperandValue(operand int, registers [3]int) int {
	if operand <= 3 {
		return operand
	}
	return registers[operand-4]
}

func comboOperandMod8(operand int, registers [3]int) int {
	return getComboOperandValue(operand, registers) % 8
}

func comboOperandDivision(operand int, registers [3]int) int {
	numerator := registers[0]
	comboOperand := getComboOperandValue(operand, registers)
	denominator := int(math.Pow(2, float64(comboOperand)))
	return numerator/denominator
}

func applyInstruction(instruction int, operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	var instructonFunc func(operand int, registers [3]int, pointer int) ([3]int, int, bool, int)
	switch instruction {
	case 0:
		instructonFunc = adv
	case 1:
		instructonFunc = bxl
	case 2:
		instructonFunc = bst
	case 3:
		instructonFunc = jnz
	case 4:
		instructonFunc = bxc
	case 5:
		instructonFunc = out
	case 6:
		instructonFunc = bdv
	case 7:
		instructonFunc = cdv
	}
	return instructonFunc(operand, registers, pointer)
}

func adv(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	registers[0] = comboOperandDivision(operand, registers)
	return registers, pointer + 2, false, 0
}

func bxl(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	registers[1] = registers[1] ^ operand
	return registers, pointer + 2, false, 0
}

func bst(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	registers[1] = comboOperandMod8(operand, registers)
	return registers, pointer + 2, false, 0
}

func jnz(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	if registers[0] == 0 {
		return registers, pointer + 2, false, 0
	}
	return registers, operand, false, 0
}

func bxc(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	return bxl(registers[2], registers, pointer)
}

func out(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	return registers, pointer + 2, true, comboOperandMod8(operand, registers)
}

func bdv(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	registers[1] = comboOperandDivision(operand, registers)
	return registers, pointer + 2, false, 0
}

func cdv(operand int, registers [3]int, pointer int) ([3]int, int, bool, int) {
	registers[2] = comboOperandDivision(operand, registers)
	return registers, pointer + 2, false, 0
}

func init() {
	solutions.RegisterSolution(17, Solution17{})
}
