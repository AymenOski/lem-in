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
	fmt.Println(ants.Tunnels)
	fmt.Println(ants.Rooms)
	fmt.Println(ants.Position)
}
