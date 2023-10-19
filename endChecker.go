package main

// Checks if any of the surrounding rooms (linkedRooms) are the endRoom
func endChecker(linkedRooms []string, endRoom string) bool {
	for _, a := range linkedRooms {
		if a == endRoom {
			return true
		}
	}
	return false
}
