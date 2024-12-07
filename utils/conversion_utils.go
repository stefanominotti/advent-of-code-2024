package utils

import "strconv"

func StringsToIntegers(strings []string) []int {
	var integers []int
	for _, line := range strings {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		integers = append(integers, n)
	}
  
	return integers
}