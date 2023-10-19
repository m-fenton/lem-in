package main

import (
	"strings"
)

// Finds the input lines that correlate to paths, separates them into ints
func edgeMaker(data []string) []string {
	var edgeString []string
	var edges []string

	//Goes through each line in turn
	for lineNumber := 0; lineNumber < len(data); lineNumber++ {
		//if the line doesn't contain a "#" and does contain a "-" then...
		if !strings.Contains(data[lineNumber], "#") && strings.Contains(data[lineNumber], "-") {
			// split it by the "-""
			edgeString = strings.Split(data[lineNumber], "-")

			//	pathInts, _ := strconv.Atoi(edgeString[j])
			edges = append(edges, edgeString...)

		}
	}

	return edges
}
