package main

import (
	"fmt"
	"os"
	"strings"

	"lem-in/functions"
	"lem-in/utils"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".txt") || len(os.Args[1]) < 5 {
		fmt.Println("Usage: go run ./cmd/ <Input_file_name.txt> ")
		return
	}

	tempAntNum, tempStartingRoom, tempEndingRoom,
		tempTunnels, tempCoord := functions.Parsing()

	Colony := &utils.Colony{
		AntNum:       tempAntNum,
		Position:     tempCoord,
		Tunnels:      tempTunnels,
		StartingRoom: tempStartingRoom,
		EndingRoom:   tempEndingRoom,
	}

	//-- ⚠️ we need also to check if the tunnels edges does exists as rooms

	// functions.ValidCoords(Ants.Position)
	// functions.ValidRooms(Ants.Rooms, Ants.Tunnels)

	g := &utils.Graph{
		Rooms: make(map[string]*utils.Room),
	}
	for room := range Colony.Tunnels {
		g.AddRoom(room)
	}
	g.LinkRooms(Colony.Tunnels)
	ants := utils.CreateAnts(Colony.AntNum, g.Rooms[Colony.StartingRoom])

	// finding all paths and choosing the best combination possible
	for {

		path := g.BFS(Colony.StartingRoom, Colony.EndingRoom)
		if path == nil {

			// this condition is to stop unecessary processing
			if len(g.Paths) >= Colony.AntNum {
				if g.Paths == nil {
					fmt.Println("ERROR: No path was found")
					os.Exit(0)
				}
				g.Simulation(ants, Colony.StartingRoom, Colony.EndingRoom)
				os.Exit(0)
			}
			bestCombo := g.Combinations(Colony.StartingRoom, Colony.EndingRoom)
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
	g.Simulation(ants, Colony.StartingRoom, Colony.EndingRoom)
}
