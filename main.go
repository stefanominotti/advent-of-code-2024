package main

import (
	"advent-of-code/solutions"
	_ "advent-of-code/solutions/01"
	_ "advent-of-code/solutions/02"
	_ "advent-of-code/solutions/03"
	_ "advent-of-code/solutions/04"
	_ "advent-of-code/solutions/05"
	_ "advent-of-code/solutions/06"
	_ "advent-of-code/solutions/07"
	_ "advent-of-code/solutions/08"
	_ "advent-of-code/solutions/09"
	_ "advent-of-code/solutions/10"
	_ "advent-of-code/solutions/11"
	_ "advent-of-code/solutions/12"
	_ "advent-of-code/solutions/13"
	_ "advent-of-code/solutions/14"
	_ "advent-of-code/solutions/15"
	_ "advent-of-code/solutions/16"
	_ "advent-of-code/solutions/17"
	_ "advent-of-code/solutions/18"
	_ "advent-of-code/solutions/19"
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
