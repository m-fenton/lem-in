package main

// if the nextRoom is the same as any of the visitedRooms then return false (will use this to only proceed to next room if true)
func roomBacktrackPreventer(visitedRooms []*Room, nextRoom *Room) bool {
	for _, a := range visitedRooms {
		if a == nextRoom {
			return true
		}
	}
	return false
}
