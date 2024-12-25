package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func filePath(day int) string {
	return fmt.Sprintf("solutions/%02d/input.txt", day)
}

func RunSolution[T any](day int, processingFunctions ...func(lineIterator *LineIterator) T) {
	downloadInput(day)
	fmt.Printf("Day %d\n", day)
	for idx, function := range processingFunctions {
		start := time.Now() // Record the start time
		result := processFile(day, function) // Call your processFile function
		elapsed := time.Since(start) 
		fmt.Printf("Part %s: %v (%s)\n", string('A'+idx), result, elapsed)
	}
}

func downloadInput(day int) {
	// Check if the file already exists
	if _, err := os.Stat(filePath(day)); err == nil {
		// File already exists, skipping download
		return
	} else if !os.IsNotExist(err) {
		panic(err)
	}

	// Retrieve the session token from the environment variable
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		panic("AOC_SESSION environment variable is not set")
	}

	// Construct the URL
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Set the session cookie
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Error donwloading input: received status code %d", resp.StatusCode))
	}

	// Save the response body to a file or print it
	file, err := os.Create(filePath(day))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}

func processFile[T any](day int, processingFunction func(lineIterator *LineIterator) T) T {
	// Open input file
	readFile, err := os.Open(filePath(day))

	if err != nil {
		panic(err)
	}

	// Build iterator for file
	fileScanner := bufio.NewScanner(readFile)
	lineIterator := NewLineIterator(fileScanner)

	// Process file iterator
	result := processingFunction(lineIterator)

	readFile.Close()

	return result
}
