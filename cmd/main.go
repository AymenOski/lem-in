package main

import (
	"fmt"
	"os"

	"lem-in/compute"
	"lem-in/parse"
	"lem-in/utils"
)

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) < 5 {
		fmt.Println("Usage: go run ./cmd/ <Input_file_name.txt> ")
		return
	}

	g, err := parse.FileToGraph(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	ants := utils.CreateAnts(g.Col.AntNum, g.Rooms[g.Col.StartingRoom]) // create ants
	compute.Solver(g, ants)

	// fmt.Println("len(g.Paths) :", len(g.Paths))
	// fmt.Print(g.Paths)
	if len(g.Paths) == 0 {
		fmt.Println("No path is found")
		return
	}
	g.Simulation(ants, g.Col.StartingRoom, g.Col.EndingRoom)
}
