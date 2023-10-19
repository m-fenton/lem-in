package main

import (
	"fmt"
)

type Ant struct {
	name        string
	currentRoom *Room
	path        []*Room
	antPathstep int
}

type Room struct {
	name        string
	linkedRooms []*Room
}

func main() {

	// filenameChecker checks for the filename, else gives an error message
	filename := filenameChecker()
	// dataReader reads the filename and returns it as a slice of string 'data'
	data := dataReader(filename)
	//numAnts reads data and returns the int number of ants
	numAnts := NumAnts(data)
	//roomFinder reads data and finds the room strings, then splits them up and returns just their names
	rooms := roomFinder(data)
	//checks for duplicate rooms, returns an error if one is found
	duplicateChecker(rooms)
	//edgemaker finds the edges (links) between rooms
	edges := edgeMaker(data)

	//Creates an array of structs based on the room names
	structRooms := roomMaker(rooms)
	// finds and adds linked rooms (structs) to rooms structs
	updatedRooms := addLinkedRoomStruct(structRooms, edges)
	//StartRoom returns both just the room name and the name with coords
	startRoomString, _ := StartRoom(data)
	//checks links to start room
	startEndLink(startRoomString, edges)
	//Finds the start room string (startRoom) and uses that to find and return a copy of the start room struct
	startRoom := startRoomStructFinder(startRoomString, updatedRooms)
	endRoomString, _ := EndRoom(data)
	//checks links to end room
	startEndLink(endRoomString, edges)
	//Finds the end room string (endRoom) and uses that to find and return a copy of the end room struct
	endRoom := endRoomStructFinder(endRoomString, updatedRooms)

	//Finds the first path, we'll use this to test all subsequent paths
	firstPath := pathFinder(updatedRooms, startRoom, endRoom, edges)
	// if the second path found is shorter than the first then we'll use this one as the first path instead
	// but only if it's a valid path (ends at the endRoom)
	if len(startRoom.linkedRooms) > 1 {
		secondPath, _ := subsequentPaths(firstPath, updatedRooms, startRoom, endRoom, edges)
		if len(firstPath) > len(secondPath) && secondPath[len(secondPath)-1] == endRoom {
			firstPath = secondPath
		}
	}
	allOfThePaths := allPaths(firstPath, updatedRooms, startRoom, endRoom, edges)
	allOfTheAntsPaths := whichPath(allOfThePaths, numAnts)
	for i := 0; i < len(allOfTheAntsPaths); i++ {

	}
	//swarm is our completed array of ants, including the paths that each will take
	swarm := antMaker(numAnts, startRoom, updatedRooms, allOfTheAntsPaths)

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
	fmt.Println()
	antMover(allOfThePaths, swarm, endRoom)
	fmt.Println()

}
