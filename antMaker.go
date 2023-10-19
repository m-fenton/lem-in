package main

import (
	"strconv"
)

// Make an  array of struct Ant, called swarm; then exports it
func antMaker(numAnts int, startRoom *Room, structRooms []*Room, paths [][]*Room) []*Ant {

	// pathNumber dictates what path we'll send the ants down; we start with the first path, 0
	pathNumber := 0
	// We declare what swarm is equal to (an array of Ant), so that we can append onto it
	swarm := []*Ant{}

	for antNumber := 0; antNumber < numAnts; antNumber++ {

		swarm = append(swarm, &Ant{
			name:        "L" + strconv.Itoa(antNumber+1),
			currentRoom: startRoom,
			path:        paths[pathNumber],
			antPathstep: 1,
		})
		// increases the path number for the next loop
		pathNumber++
	}
	return swarm
}

//cycle through all of the paths until an option runs out, then cycle through the rest
