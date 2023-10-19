package main

import (
	"fmt"
	"os"
)

// Finds the first path
func pathFinder(updatedRooms []*Room, startRoom *Room, endRoom *Room, edges []string) []*Room {

	thisPath := make([]*Room, 1)
	thisPath[0] = startRoom

	//Turn counter
	for turnCounter := 0; turnCounter < len(updatedRooms); turnCounter++ {

		//Room Option Counter
		for j := 0; thisPath[turnCounter] != endRoom; j++ {

			//if the next room options is the end, append the end on and return Path
			if endChecker(linkedRooms(thisPath[turnCounter].name, edges), endRoom.name) {

				thisPath = append(thisPath, endRoom)

				return thisPath
			}
			nextRoom := thisPath[turnCounter].linkedRooms[j]

			//append next room on only if it's not backtracking
			if !roomBacktrackPreventer(thisPath, nextRoom) {

				thisPath = append(thisPath, nextRoom)

				turnCounter++
				j = -1
			}

		}
	}

	if thisPath[len(thisPath)-1] != endRoom {
		fmt.Println("ERROR: invalid data format.")
		os.Exit(0)
	}

	return thisPath
}
