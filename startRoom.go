package main

import (
	"fmt"
	"os"
	"strings"
)

// StartRoom returns an int that is the starting room
func StartRoom(data []string) (string, string) {
	var startRoomString string
	var startRoomWithCoords string
	sliceContains(data, "##start")

	for i := 0; i < len(data); i++ {
		if data[i] == "##start" {
			if !strings.Contains(data[i+1], " ") {
				fmt.Println("ERROR: invalid data format, start room does not exist.")
				os.Exit(0)
			}
			startRoomWithCoords = string(data[i+1])
			startRoomString = strings.Split(string(data[i+1]), " ")[0]
		}
	}
	return startRoomString, startRoomWithCoords
}

//Finds and returns the room (struct) that is the startRoom; useful in a variety of functions
func startRoomStructFinder(startRoom string, structRooms []*Room) *Room {
	var startStructRoom *Room
	for i := 0; i < len(structRooms); i++ {
		if structRooms[i].name == startRoom {

			startStructRoom = structRooms[i]
		}
	}
	return startStructRoom
}
