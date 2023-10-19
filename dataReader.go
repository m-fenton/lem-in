package main

import (
	"fmt"
	"os"
	"strings"
)

// Reads antfarm.txt and returns each line of text as a separate string
func dataReader(filename string) []string {
	// reads input from antfarm.txt
	input, err := os.ReadFile("examples/" + filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format, unable to find/read file.")
		os.Exit(0)
	}
	// splits input by newline
	data := strings.Split(string(input), "\n")
	return data
}
