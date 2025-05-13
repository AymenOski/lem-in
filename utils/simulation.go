package utils

import (
	"fmt"
	"os"
)

// Simulation assigns ants to paths in a balanced way and simulates their step-by-step movement.
// Each turn ensures that room occupancy and tunnel rules are respected until all ants reach the end.
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
	// create the out for the standard output
	out, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	output := string(out)
	output += "\n"
	fmt.Println(output)

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
	// for debbuging puposes
	// fmt.Println("Steps : ", c)
}
