package main

// Returns all rooms linked to room 'inputRoom'
func linkedRooms(inputRoom string, edges []string) []string {
	var linkedRooms []string

	for i := 0; i < len(edges); i++ {
		if edges[i] == inputRoom {
			if i%2 == 0 {
				ifOdd := edges[i+1]
				linkedRooms = append(linkedRooms, ifOdd)

			} else {
				ifEven := edges[i-1]
				linkedRooms = append(linkedRooms, ifEven)

			}
		}
	}

	return linkedRooms
}
