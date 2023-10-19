package main

func subsequentPaths(previousPath []*Room, updatedRooms []*Room, startRoom *Room, endRoom *Room, edges []string) ([]*Room, []*Room) {

	var allVistedRooms []*Room
	// Create array of rooms that is all previously visited rooms?  Could feed this into the function
	// Feed it into the function instead of previous path then repeatedly add to itd
	// Make it... recursive? until hits an error

	// Creating a slice of pointers to Room structs.  The length of the slice is 1.
	thisPath := make([]*Room, 1)
	thisPath[0] = startRoom
	counter := 0
	//Turn counter

	for roomsAdded := 0; roomsAdded < len(updatedRooms); roomsAdded++ {
		//if there's only one way out of the startRoom then skip all this
		//we know it can't produce more than one valid route
		if len(startRoom.linkedRooms) < 2 {
			break
		}
		//if there's only one way to the endRoom then skip all this
		//we know it can't produce more than one valid route
		if len(endRoom.linkedRooms) < 2 {
			break
		}
		//I think this was to return the path (which will be shown to be invalid)
		//If the route reached a dead end
		if roomsAdded > len(thisPath)-1 {
			return thisPath, allVistedRooms
		}
		// if the previous path's length is 2 (so it goes straight to the end)
		//if len(previousPath) == 2 {

		for i := 0; i < len(thisPath[roomsAdded].linkedRooms); i++ {
			counter++
			if roomsAdded > 0 {

				//if the next room options is the end, append the end on and return Path
				if endChecker(linkedRooms(thisPath[roomsAdded].name, edges), endRoom.name) {

					thisPath = append(thisPath, endRoom)

					allVistedRooms = previousPath
					allVistedRooms = append(allVistedRooms, thisPath...)

					return thisPath, allVistedRooms
				}
			}
			nextRoom := thisPath[roomsAdded].linkedRooms[i]

			// A bodge solution to the pathfinding... Need to revamp
			if nextRoom.name != "G0" && nextRoom.name != "D1" {
				if !roomBacktrackPreventer(previousPath, nextRoom) {

					if !roomBacktrackPreventer(thisPath, nextRoom) {

						thisPath = append(thisPath, nextRoom)

						roomsAdded++
						i = -1
					}
				}
			}
		}

	}
	allVistedRooms = previousPath
	allVistedRooms = append(allVistedRooms, thisPath...)

	return thisPath, allVistedRooms
}
