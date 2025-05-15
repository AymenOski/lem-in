package utils

func (g *Graph) Combinations(start, end string) [][]string {
	AllCombination := [][][]string{}
	ComboElements := [][]string{}

	for _, neighbor := range g.Rooms[start].Neighbors {
		neighbor.Occupied = false
	}

	// calculating all the probabilities
	for _, ThePathToTest := range g.Paths {
		for _, room := range ThePathToTest {
			g.Rooms[room].Occupied = true
		}

		for {

			path := g.BFS(start, end)

			if path == nil {

				// the path test with
				ComboElements = append(ComboElements, ThePathToTest)

				for _, rooms := range ComboElements {
					for _, v := range rooms {
						g.Rooms[v].Occupied = false
					}
				}

				break
			}

			for _, room := range path {
				g.Rooms[room].Occupied = true
			}

			// the path we found at this iteration
			ComboElements = append(ComboElements, path)
		}
		AllCombination = append(AllCombination, ComboElements)
		ComboElements = [][]string{}
	}

	// -----------for debuging purposes :-------- \\
	// last line after 🔵 Group will be the path to test with \\

	// for i, group := range AllCombination {
	// 	fmt.Printf("🔵 Group %d:\n", i+1)
	// 	for j, path := range group {
	// 		fmt.Printf("  🔸 Path %d: %v\n", j+1, path)
	// 	}
	// 	fmt.Println()
	// }

	return Filtring(AllCombination)
}

func Filtring(allCombos [][][]string) [][]string {
	var bestCombo [][]string
	maxPaths := -1
	minRooms := int(^uint(0) >> 1)
	for _, combo := range allCombos {
		numPaths := len(combo)
		numRooms := 0

		for _, path := range combo {
			numRooms += len(path)
		}

		if numPaths > maxPaths || (numPaths == maxPaths && numRooms < minRooms) {
			bestCombo = combo
			maxPaths = numPaths
			minRooms = numRooms
		}
	}

	return bestCombo
}
