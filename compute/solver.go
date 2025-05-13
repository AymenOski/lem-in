package compute

import (
	"fmt"
	"os"

	"lem-in/utils"
)

// Finding all paths and choosing the best combination possible
func Solver(g *utils.Graph, ants []*utils.Ant) {
	for {

		path := g.BFS(g.Col.StartingRoom, g.Col.EndingRoom)
		if path == nil {
			// this condition is to stop unecessary processing
			if len(g.Paths) >= g.Col.AntNum {
				if len(g.Paths) == 0 {
					fmt.Println("ERROR: No path was found")
					os.Exit(0)
				}
				g.Simulation(ants, g.Col.StartingRoom, g.Col.EndingRoom)
				os.Exit(0)
			}

			bestCombo := g.Combinations(g.Col.StartingRoom, g.Col.EndingRoom)
			g.Paths = [][]string{}
			g.Paths = append(g.Paths, bestCombo...)
			break
		}
		g.Paths = append(g.Paths, path)

	}

	if len(g.Paths) == 0 {
		fmt.Println("ERROR: No path was found")
		os.Exit(0)
	}
}
