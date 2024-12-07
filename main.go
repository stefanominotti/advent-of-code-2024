package main

import (
	"advent-of-code/solutions"
	_ "advent-of-code/solutions/01"
	_ "advent-of-code/solutions/02"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Command-line flags
	runAll := flag.Bool("all", false, "Run all solutions")
	runSingle := flag.String("solution", "", "Run a specific solution (e.g., 01, 02)")
	flag.Parse()

	if *runAll {
		solutions.RunAll()
	} else if *runSingle != "" {
		day, err := strconv.Atoi(*runSingle)
		if err != nil || day < 0 || day > 25 {
			fmt.Printf("Error: %s is not a valid day", *runSingle)
			os.Exit(1)
		}
		solutions.RunSolution(day)
	} else {
		fmt.Println("Please specify a solution with -solution or run all with -all")
		flag.Usage()
	}
}
