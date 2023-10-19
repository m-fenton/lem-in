package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the number of ants from the text file
func NumAnts(data []string) int {
	antNum := data[0]

	if _, err := strconv.Atoi(antNum); err != nil {
		fmt.Println("ERROR: invalid data format, ant number needs to be an int.")
		os.Exit(0)
	}

	if antNum == "" || antNum == "0" {
		fmt.Println("ERROR: invalid data format, invalid number of Ants.")
		os.Exit(0)
	}
	//I added this in case having no entry for ants ans line 1 beginning with a '#' then give error.
	//This may never arise so may be useless. We can take it out.
	if strings.HasPrefix(antNum, "#") {
		fmt.Println("ERROR: invalid data format.")
		os.Exit(0)
	}

	antNumInt, _ := strconv.Atoi(antNum)
	return antNumInt
}
