package main

import (
	"fmt"
	"os"
	"sort"
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
	ants := functions.CreateAnts(Colony.AntNum, g.Rooms[Colony.StartingRoom])
	for {
		path := g.BFS(Colony.StartingRoom, Colony.EndingRoom)
		if path == nil {
			break
		}
		g.Paths = append(g.Paths, path)
	}
	
	fmt.Println(g.Paths)
	sort.Slice(g.Paths, func(i, j int) bool {
		return len(g.Paths[i]) < len(g.Paths[j])
	})

	g.Simulation(ants, Colony.StartingRoom, Colony.EndingRoom)
}
