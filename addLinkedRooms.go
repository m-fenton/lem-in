package main

// creates an array of structs (rooms) that can later be added to the current room structs
func linkedRoomStructs(currentRoomName string, Rooms []*Room, edges []string) []*Room {

	roomLinks := linkedRooms(currentRoomName, edges)
	var LinkedRoomsStructs []*Room

	for roomLinkNumber := 0; roomLinkNumber < len(roomLinks); roomLinkNumber++ {
		for room := 0; room < len(Rooms); room++ {
			if Rooms[room].name == roomLinks[roomLinkNumber] {
				LinkedRoomsStructs = append(LinkedRoomsStructs, Rooms[room])
			}

		}
	}
	return LinkedRoomsStructs
}

// Adds the array of linked rooms (structs) to each of the room structs
func addLinkedRoomStruct(structRooms []*Room, edges []string) []*Room {

	for i := 0; i < len(structRooms); i++ {
		structRooms[i].linkedRooms = linkedRoomStructs(structRooms[i].name, structRooms, edges)
	}

	return structRooms
}
