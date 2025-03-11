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
	tempAntNum, tempStartingRoom, tempEndingRoom, tempTunnels, tempRooms := functions.ExtractingDate()
	ants := &utils.Ants{
		AntNum:       tempAntNum,
		Rooms:        tempRooms,
		Tunnels:      tempTunnels,
		StartingRoom: tempStartingRoom,
		EndingRoom:   tempEndingRoom,
	}
	fmt.Println(ants.AntNum)
	fmt.Println(ants.StartingRoom)
	fmt.Println(ants.EndingRoom)
	fmt.Println(ants.Rooms)
	fmt.Println(ants.Tunnels)
}
