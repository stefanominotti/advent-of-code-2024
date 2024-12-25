package solution05

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type Solution05 struct{}

func (s Solution05) PartA(lineIterator *utils.LineIterator) any {
	processUpdateFunc := func(rules map[string]bool, update []string) int {
		if isValidUpdate(rules, update) {
			updateInts := utils.StringsToIntegers(update)
			return updateInts[len(updateInts)/2]
		}
		return 0
	}
	return processUpdates(lineIterator, processUpdateFunc)
}

func (s Solution05) PartB(lineIterator *utils.LineIterator) any {
	processUpdateFunc := func(rules map[string]bool, update []string) int {
		if !isValidUpdate(rules, update) {
			update := fixUpdate(rules, update)
			updateInts := utils.StringsToIntegers(update)
			return updateInts[len(updateInts)/2]
		}
		return 0
	}
	return processUpdates(lineIterator, processUpdateFunc)
}

func fixUpdate(rules map[string]bool, update []string) []string {
	for !isValidUpdate(rules, update) {
		for i := 1; i < len(update); i++ {
			current := update[i]
			for j := 0; j < i; j++ {
				rule := fmt.Sprintf("%s|%s", current, update[j])
				_, contained := rules[rule]
				if contained {
					update[i] = update[j]
					update[j] = current
					break
				}
			}
		}
	}
	return update
}

func isValidUpdate(rules map[string]bool, update []string) bool {
	for idx, num := range update {
		if idx == 0 {
			continue
		}
		for i := 0; i < idx; i++ {
			rule := fmt.Sprintf("%s|%s", num, update[i])
			_, contained := rules[rule]
			if contained {
				return false
			}
		}
	}
	return true
}

func processUpdates(lineIterator *utils.LineIterator, processUpdateFunc func(map[string]bool, []string) int) int {
	rules := map[string]bool{}
	isUpdatesSection := false
	result := 0
	for lineIterator.Next() {
		line := lineIterator.Value()
		if line == "" {
			isUpdatesSection = true
			continue
		}
		if !isUpdatesSection {
			rules[line] = true
			continue
		}

		update := strings.Split(line, ",")
		result += processUpdateFunc(rules, update)
	}
	return result
}

func init() {
	solutions.RegisterSolution(5, Solution05{})
}
