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
	tempAntNum, tempStartingRoom, tempEndingRoom, tempTunnels, tempRooms, tempCoord := functions.Parsing()
	Ants := &utils.Ants{
		AntNum:       tempAntNum,
		Rooms:        tempRooms,
		Position:     tempCoord,
		Tunnels:      tempTunnels,
		StartingRoom: tempStartingRoom,
		EndingRoom:   tempEndingRoom,
	}
	g := &utils.Graph{
		Rooms: make(map[string]*utils.Room),
	}
	for i := range Ants.Rooms {
		g.AddRoom(Ants.Rooms[i])
	}
	g.LinkRooms(Ants.Tunnels)
}
