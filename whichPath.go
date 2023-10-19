package main

import (
	"sort"
)

func whichPath(allOfThePaths [][]*Room, numAnts int) [][]*Room {
	var allOfTheAntsPaths [][]*Room
	antNumber := 0

	if len(allOfThePaths) > 1 {
		//sorts the paths by length
		sort.Slice(allOfThePaths, func(i, j int) bool {
			return len(allOfThePaths[i]) > len(allOfThePaths[j])
		})

		//antsInPaths := make([]int, len(allOfThePaths))
		antsPathCombo := make([]int, len(allOfThePaths))
		//Creates an array of the length of each path
		for pathNumber := 0; pathNumber < len(allOfThePaths); pathNumber++ {
			antsPathCombo[pathNumber] = len(allOfThePaths[pathNumber])

		}

		//Trying to get the order of paths that needs to be applied to ants
		for turnCounter := 0; antNumber < numAnts; turnCounter++ {
			for pathNumber := 0; pathNumber < len(allOfThePaths); pathNumber++ {

				if pathNumber < len(allOfThePaths)-1 {
					if antsPathCombo[pathNumber] <= antsPathCombo[pathNumber+1] {
						antsPathCombo[pathNumber]++
						antNumber++
						allOfTheAntsPaths = append(allOfTheAntsPaths, allOfThePaths[pathNumber])

						if antNumber == numAnts {
							for i, j := 0, len(allOfTheAntsPaths)-1; i < j; i, j = i+1, j-1 {
								allOfTheAntsPaths[i], allOfTheAntsPaths[j] = allOfTheAntsPaths[j], allOfTheAntsPaths[i]
							}
							return allOfTheAntsPaths
						}
					}
				} else if antsPathCombo[len(allOfThePaths)-1] < antsPathCombo[len(allOfThePaths)-2] {
					antsPathCombo[len(allOfThePaths)-1]++
					antNumber++
					allOfTheAntsPaths = append(allOfTheAntsPaths, allOfThePaths[len(allOfThePaths)-1])

					if antNumber == numAnts {
						for i, j := 0, len(allOfTheAntsPaths)-1; i < j; i, j = i+1, j-1 {
							allOfTheAntsPaths[i], allOfTheAntsPaths[j] = allOfTheAntsPaths[j], allOfTheAntsPaths[i]
						}
						return allOfTheAntsPaths
					}
				}
			}

		}
	} else if len(allOfThePaths) == 1 {
		for i := 0; i < numAnts; i++ {
			allOfTheAntsPaths = append(allOfTheAntsPaths, allOfThePaths[0])
		}

	}

	return allOfTheAntsPaths
}
