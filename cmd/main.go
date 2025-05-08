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
		tempTunnels, tempRooms, tempCoord := functions.Parsing()

	Colony := &utils.Colony{
		AntNum:       tempAntNum,
		Rooms:        tempRooms,
		Position:     tempCoord,
		Tunnels:      tempTunnels,
		StartingRoom: tempStartingRoom,
		EndingRoom:   tempEndingRoom,
	}

	// functions.ValidCoords(Ants.Position)
	// functions.ValidRooms(Ants.Rooms, Ants.Tunnels)
	
	g := &utils.Graph{
		Rooms: make(map[string]*utils.Room),
	}

	for i := range Colony.Rooms {
		g.AddRoom(Colony.Rooms[i])
	}

	g.LinkRooms(Colony.Tunnels)
	ants := utils.CreateAnts(Colony.AntNum, g.Rooms[Colony.StartingRoom])
	
	// next loop condition is for the fact that we need to find all the paths possible
	for {
		path := g.BFS(Colony.StartingRoom, Colony.EndingRoom)

		if path == nil {
			bestCombo := g.Combinations(Colony.StartingRoom, Colony.EndingRoom)
			g.Paths = [][]string{}
			g.Paths = append(g.Paths, bestCombo...)
			break
		}

		g.Paths = append(g.Paths, path)
	}

	if len(g.Paths) == 0 {
		fmt.Println("ERROR: No path was found")
		return
	}
	
	fmt.Println(g.Paths)
	g.Simulation(ants, Colony.StartingRoom, Colony.EndingRoom)
	fmt.Println("LenPaths", len(g.Paths))
}
