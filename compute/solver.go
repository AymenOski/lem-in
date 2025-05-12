package compute

import (
	"fmt"
	"os"

	"lem-in/utils"
)

func Solver(g *utils.Graph, ants []*utils.Ant) {
	// this for loop is to stop unecessary processing
	for {

		path := g.BFS(g.Col.StartingRoom, g.Col.EndingRoom)
		if path == nil {
			if len(g.Paths) >= g.Col.AntNum {
				if g.Paths == nil {
					fmt.Println("ERROR: No path was found")
					os.Exit(0)
				}
				g.Simulation(ants, g.Col.StartingRoom, g.Col.EndingRoom)
				os.Exit(0)
			} else {
				break
			}
		}
		g.Paths = append(g.Paths, path)

	}

	// finding all paths and choosing the best combination possible
	for {

		path := g.BFS(g.Col.StartingRoom, g.Col.EndingRoom)
		if path == nil {
			bestCombo := g.Combinations(g.Col.StartingRoom, g.Col.EndingRoom)
			g.Paths = [][]string{}
			g.Paths = append(g.Paths, bestCombo...)
			break
		}
		g.Paths = append(g.Paths, path)

	}

	if g.Paths == nil {
		fmt.Println("ERROR: No path was found")
		os.Exit(0)
	}
}
