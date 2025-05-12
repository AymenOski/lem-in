package utils

import "fmt"

func GraphConstructor() *Graph {
	return &Graph{Rooms: make(map[string]*Room)}
}

func (g *Graph) LinkRooms(tunnels map[string][]string) {
	// bidirectional links between rooms
	for roomName, neighbors := range tunnels {
		for _, neighborName := range neighbors {
			if g.Rooms[roomName] != nil && g.Rooms[neighborName] != nil {
				g.Rooms[roomName].Neighbors = append(g.Rooms[roomName].Neighbors, g.Rooms[neighborName])
			}
		}
	}
}

func (g *Graph) AddRoom(roomName string) {
	if _, exists := g.Rooms[roomName]; !exists {
		g.Rooms[roomName] = &Room{
			Name:      roomName,
			Neighbors: make([]*Room, 0),
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, room := range g.Rooms {
		fmt.Printf("Room: %s, Neighbors: %v\n", room.Name, room.Neighbors)
	}
}

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

	for i, group := range AllCombination {
		fmt.Printf("🔵 Group %d:\n", i+1)
		for j, path := range group {
			fmt.Printf("  🔸 Path %d: %v\n", j+1, path)
		}
		fmt.Println()
	}

	return Filtring(AllCombination)
}

func (g *Graph) BFS(start, end string) []string {
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	prev := make(map[string]string)
	skipStartEndPath := false

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range g.Rooms[current].Neighbors {
			if !skipStartEndPath {
				if hasLengthString(g.Paths) && neighbor.Name == end {
					skipStartEndPath = true
					continue
				}
			}

			if !visited[neighbor.Name] && (!neighbor.Occupied || neighbor.Name == end) {
				visited[neighbor.Name] = true
				prev[neighbor.Name] = current
				queue = append(queue, neighbor.Name)
			}
		}
	}
	if !visited[end] {
		return nil
	}

	var path []string

	for at := end; at != ""; at = prev[at] {
		if prev[at] == start {
			g.Rooms[at].Occupied = true
		}
		path = append([]string{at}, path...)
	}

	return path
}

func (g *Graph) Simulation(ants []*Ant, Start string, End string) {
	pathLens := make([]int, len(g.Paths))

	for _, room := range g.Rooms {
		room.Occupied = false
	}

	for i := range pathLens {
		pathLens[i] = len(g.Paths[i])
	}

	// Assign the right path to the ants
	for _, ant := range ants {
		// Assign paths using a non-greedy strategy to balance load across paths
		// The goal is to optimize total movement time for the entire colony
		bestIdx := 0

		// Selecting the shortest available path to reduce overall steps needed
		for i := 1; i < len(pathLens); i++ {
			if pathLens[bestIdx] > pathLens[i] {
				bestIdx = i
			}
		}

		ant.Path = g.Paths[bestIdx]
		pathLens[bestIdx]++
	}

	c := 0

	allReachedEnd := false

	for !allReachedEnd {
		tunnelCrowding := false

		// start one round at a time
		for _, ant := range ants {
			if ant.CurrentRoom != g.Rooms[End] && ant.Step < len(ant.Path)-1 {
				// Each tunnel can only be used once per turn.
				// Skip ants on direct Start-End path if already used this round (tunnelCrowding = true)
				if len(ant.Path) == 2 && tunnelCrowding {
					continue
				}

				nextRoomName := ant.Path[ant.Step+1]
				nextRoom := g.Rooms[nextRoomName]

				if !tunnelCrowding {
					if nextRoom.Name == End && ant.CurrentRoom.Name == Start {
						tunnelCrowding = true
						ant.Step++
						ant.CurrentRoom = g.Rooms[End]
						ant.HasMoved = true
					}
				}

				if !nextRoom.Occupied || nextRoom.Name == End {
					ant.CurrentRoom = nextRoom
					ant.CurrentRoom.Occupied = true

					ant.Step++
					ant.HasMoved = true
				}
			}
		}

		// Only print ants that actually moved this round (using HasMoved flag)
		for _, ant := range ants {
			if ant.HasMoved {
				fmt.Printf("%s-%s ", ant.Id, ant.CurrentRoom.Name)
				ant.HasMoved = false
				ant.CurrentRoom.Occupied = false
			}
		}
		c++
		fmt.Println()

		allReachedEnd = true
		for _, ant := range ants {
			if ant.CurrentRoom != g.Rooms[End] {
				allReachedEnd = false
			}
		}

	}

	fmt.Println("Steps : ", c)
}
