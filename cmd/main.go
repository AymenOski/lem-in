package main

import (
	"fmt"
	"os"
	"strings"

	"lem-in/parse"
	"lem-in/utils"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".txt") || len(os.Args[1]) < 5 {
		fmt.Println("Usage: go run ./cmd/ <Input_file_name.txt> ")
		return
	}

	g, _ := parse.FileToGraph(os.Args[1])


	//-- ⚠️ we need also to check if the tunnels edges does exists as rooms

	for _, room := range g.Col.Rooms { // create room struct
		g.AddRoom(room)
	}
	g.LinkRooms(g.Col.Tunnels)                                          // links
	ants := utils.CreateAnts(g.Col.AntNum, g.Rooms[g.Col.StartingRoom]) // create ants

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

	fmt.Println("len(g.Paths) :", len(g.Paths))
	g.Simulation(ants, g.Col.StartingRoom, g.Col.EndingRoom)
}
