package main

import (
	"fmt"
	"os"
	"strings"
)

// EndRoom returns a string that is the end room
func EndRoom(data []string) (string, string) {
	var endRoomString string
	var endRoomWithCoords string

	// checks that there is an endroom, or else writes an error
	sliceContains(data, "##end")

	for i := 0; i < len(data); i++ {
		if data[i] == "##end" {
			if !strings.Contains(data[i+1], " ") {
				fmt.Println("ERROR: invalid data format, end room does not exist.")
				os.Exit(0)
			}
			endRoomWithCoords = data[i+1]
			endRoomString = strings.Split(string(data[i+1]), " ")[0]
		}
	}
	return endRoomString, endRoomWithCoords
}

//Finds and returns the room (struct) that is the endRoom; useful in a variety of functions
func endRoomStructFinder(endRoom string, updatedRooms []*Room) *Room {
	var endStructRoom *Room
	for i := 0; i < len(updatedRooms); i++ {
		if updatedRooms[i].name == endRoom {

			endStructRoom = updatedRooms[i]
		}
	}
	return endStructRoom
}
