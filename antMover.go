package main

import "fmt"

func antMover(allPaths [][]*Room, swarm []*Ant, endRoom *Room) {

	numberOfPaths := len(allPaths)
	var sentAnts []*Ant

	// turnCounter keeps the programme ticking until len(swarm) = 0
	// Ants are removed when they reach the end, so len(swarm) will = 0 when all ants have left through the end
	for turnCounter := 0; len(swarm) != 0; turnCounter++ {

		//Each turn it adds as many ants as there are paths, until there are no more ants to add
		for antsEntering := 0; antsEntering < numberOfPaths; antsEntering++ {
			if len(sentAnts) < len(swarm) {
				sentAnts = append(sentAnts, swarm[antsEntering])
			}
		}

		antsPresent := len(sentAnts)
		// Go through each ant in turn, up to the number currently in the Paths.  As ants are removed we decrease the number of antsPresent
		for antsInTransitNumber := 0; antsInTransitNumber < antsPresent; antsInTransitNumber++ {
			//An inelegant way of moving the ants around whilst leaving the original rooms in tact... I think
			swarm[antsInTransitNumber].currentRoom = swarm[antsInTransitNumber].path[swarm[antsInTransitNumber].antPathstep]
			//Prints ants name and location
			fmt.Print(swarm[antsInTransitNumber].name, "-", swarm[antsInTransitNumber].currentRoom.name, " ")
			// I added a counter to each ant, so that I could increase their path without risking later ants trying to start at a later node
			swarm[antsInTransitNumber].antPathstep++
			// if the current ant is in the endRoom then remove it and decrease the number of ants in transit and the number of ants present
			// I feel like maybe there's a better way of writing this? Dunno, I'm tired
			if swarm[antsInTransitNumber].currentRoom == endRoom {

				antNumber := antsInTransitNumber
				sentAnts = removeAnt(sentAnts, antNumber)
				swarm = removeAnt(swarm, antNumber)

				antsInTransitNumber--
				antsPresent--

			}
		}
		//not sure what this does
		fmt.Println()
	}

}

// Removes the ants from the swarm when they reach the end
func removeAnt(swarm []*Ant, antNumber int) []*Ant {
	return append(swarm[:antNumber], swarm[antNumber+1:]...)
}
