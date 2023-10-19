package main

import (
	"fmt"
	"os"
)

// Makes sure that there's the correct number of arguments (2)
func filenameChecker() string {

	var filename string
	if len(os.Args) == 2 {
		// specify the filename in the arg, e.g. go run main.go antfarm.txt
		filename = os.Args[1]
	} else {
		// Otherwise returns an error
		err := fmt.Errorf("ERROR: Invalid number of arguments.  Please specify the input file. ")
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return filename
}

// checks if there are repeat rooms
func duplicateChecker(rooms []string) {

	for i := 0; i < len(rooms); i++ {
		counter := 0
		for j := 0; j < len(rooms); j++ {
			if rooms[i] == rooms[j] {
				counter++
				if counter == 2 {
					fmt.Println("ERROR: there are duplicate rooms.")
					os.Exit(0)
				}
			}
		}
	}

}

// checks to see if array of slices contains string 'e'
// Normally ##start or ##end
func sliceContains(data []string, e string) {

	for a := 0; a < len(data); a++ {
		if data[a] == e {
			return
		}
	}
	fmt.Println("ERROR: invalid data format,", e, "room does not exist.")
	os.Exit(0)
}

// Checks to see that there's links for the input room (start or end)
func startEndLink(se string, edges []string) {

	if len(linkedRooms(se, edges)) == 0 {
		fmt.Println("ERROR: invalid data format, no path found.")
		os.Exit(0)
	}
}
