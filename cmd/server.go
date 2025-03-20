package main

import (
	"fmt"
	"os"

	"lem-in/functions"
	"lem-in/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run ./cmd/ <input.txt> ")
		return
	}
	tempAntNum, tempStartingRoom, tempEndingRoom, tempTunnels, tempRooms, tempCoord := functions.Parsing()
	ants := &utils.Ants{
		AntNum:       tempAntNum,
		Rooms:        tempRooms,
		Position:     tempCoord,
		Tunnels:      tempTunnels,
		StartingRoom: tempStartingRoom,
		EndingRoom:   tempEndingRoom,
	}
	fmt.Println(ants.AntNum)
	fmt.Println(ants.StartingRoom)
	fmt.Println(ants.EndingRoom)
	fmt.Println(ants.Position)
	fmt.Println(ants.Rooms)
	fmt.Println(ants.Tunnels)
}
