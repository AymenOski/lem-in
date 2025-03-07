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
	tempAntNum, tempStartingRoom, tempEndingRoom, tempTunnels := functions.ExtractingDate()
	ants := utils.Ants{
		AntNum:       tempAntNum,
		StartingRoom: tempStartingRoom,
		EndingRoom:   tempEndingRoom,
		Tunnels:      tempTunnels,
	}
	fmt.Println(ants.AntNum)
	fmt.Println(ants.StartingRoom)
	fmt.Println(ants.EndingRoom)
	fmt.Println(ants.Tunnels)
}
