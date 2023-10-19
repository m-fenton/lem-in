package main

func allPaths(firstPath []*Room, updatedRooms []*Room, startRoom *Room, endRoom *Room, edges []string) [][]*Room {
	var thisPath []*Room
	//var allPaths [2][]*Room
	var allOfThePaths [][]*Room
	var allVistedPaths []*Room
	var experimentalPath []*Room

	//Fills variable allOfThePaths with the first path that we found, we'll be adding to this [][] later on
	allOfThePaths = append(allOfThePaths, firstPath)
	//finds the second path, this path; also returns a list of every room that has been thus far visited
	thisPath, allVistedPaths = subsequentPaths(firstPath, updatedRooms, startRoom, endRoom, edges)

	//if the first iteration of subsequent Paths finds a successful path
	if thisPath[len(thisPath)-1] == endRoom {
		//append the path to allOfThePaths
		allOfThePaths = append(allOfThePaths, thisPath)
		for newPathCount := 0; newPathCount < 1; newPathCount++ {

			experimentalPath, allVistedPaths = subsequentPaths(allVistedPaths, updatedRooms, startRoom, endRoom, edges)
			// if it finds a successful path then add it to the list of all paths
			if experimentalPath[len(experimentalPath)-1] == endRoom {
				allOfThePaths = append(allOfThePaths, experimentalPath)

			} else {
				break
			}
		}
	}

	return allOfThePaths
}
