package main

import "strings"

// the rooms are 3 digits separated by spaces, finds the strings with a ' ' (and without #) and splits those string by ' ', then converts them into numbers
func roomFinder(data []string) []string {
	var roomNames []string
	var roomData []string

	for i := 0; i < len(data); i++ {
		// For every slice of string, ignore if it contains a # (# lines are for comments, etc)
		//...looking for " " (only rooms will have it)
		if !strings.Contains(data[i], "#") && strings.Contains(data[i], " ") {

			// splits the lines by " "
			roomData = strings.Split(data[i], " ")
			for j := 0; j < len(roomData); j += 3 {

				roomNames = append(roomNames, roomData[j])
			}
		}
	}
	return roomNames
}
