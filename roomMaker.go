package main

// Makes a struct for each room
func roomMaker(rooms []string) []*Room {
	var allRooms []*Room
	for i := 0; i < len(rooms); i++ {
		allRooms = append(allRooms, &Room{
			name: rooms[i],
			//	linkedRooms: addLinkedRooms(name),
		})
	}
	return allRooms
}
