package main

import (
	"fmt"
	"os"
	"strings"

	"lem-in/compute"
	"lem-in/parse"
	"lem-in/utils"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Usage: go run ./cmd/ <Input_file_path>/<example.txt> ")
		return
	}

	g, err := parse.FileToGraph(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	ants := utils.CreateAnts(g.Col.AntNum, g.Rooms[g.Col.StartingRoom]) // Create the nat object
	compute.Solver(g, ants)

	g.Simulation(ants, g.Col.StartingRoom, g.Col.EndingRoom)
}
