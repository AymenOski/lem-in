package functions

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/utils"
)

func Parsing() (int, string, string, map[string][]string, []string, []utils.Coordinates) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please create a file in the root of the project and put input to it")
		os.Exit(0)
	}
	Lines := strings.Split(string(file), "\n")
	AntNum := 0
	Rooms := make([]string, 0)
	Coord := make([]utils.Coordinates, 0)
	Tunnels := make(map[string][]string)
	// FC denoted as "format checker"
	FC_StartFlag := 0
	FC_EndFlag := 0
	for i := range Lines {
		switch {
		case strings.HasPrefix(Lines[i], "L"):
			fmt.Println("ERROR: Invalid input format, Your room name starts with an L")
			os.Exit(0)
		case Lines[i] == "##start":
			FC_StartFlag = i
		case Lines[i] == "##end":
			FC_EndFlag = i
		}
	}

	FormatChekcer(FC_StartFlag, FC_EndFlag, Rooms)
	i := 0
	for i < len(Lines) {
		val := Lines[i]
		switch {
		case i == 0:
			AntNum, err = strconv.Atoi(val)
			if err != nil {
				fmt.Println("ERROR: Invalid input format, Please ensure the number of ants is provided as an integer on the first line.")
				os.Exit(0)
			}
			if AntNum <= 0 {
				fmt.Println("ERROR: invalid data format, Number of ants should be greater than 0")
				os.Exit(0)
			}
		case val == "##start":
			for j := i + 1; !strings.HasPrefix(Lines[j], "##end"); j++ {
				// if its a comment skip it
				if strings.HasPrefix(Lines[j], "#") {
					continue
				}
				// Storing staring room ,rooms and coordinates
				data := strings.Fields(Lines[j])
				if len(data) != 3 {
					fmt.Println("ERROR: Invalid input format.")
					os.Exit(0)
				} else {
					Rooms = append(Rooms, data[0])
					Coord = append(Coord, utils.Coordinates{X: utils.Atoi(data[1]), Y: utils.Atoi(data[2])})
				}
				// its necessary to update i so we can reduce the number of unnecessary iterations
				i = j
			}
		case val == "##end":
			cmp := 0
			for j := i + 1; j < len(Lines); j++ {
				// if its a comment skip it
				if strings.HasPrefix(Lines[j], "#") {
					continue
				}
				if !strings.Contains(Lines[j], "-") {
					cmp++
					temp := strings.Fields(Lines[j])
					if len(temp) != 3 {
						fmt.Println("ERROR: Invalid input format.")
						os.Exit(0)
					} else {
						Rooms = append(Rooms, temp[0])
						Coord = append(Coord, utils.Coordinates{X: utils.Atoi(temp[1]), Y: utils.Atoi(temp[2])})
					}
				}
				if cmp != 1 {
					fmt.Println("ERROR: Invalid input format.")
					os.Exit(0)
				}
				if strings.Contains(Lines[j], "-") {
					// Storing end room and tunnels
					tunnel := strings.Split(Lines[j], "-")
					if len(tunnel) == 2 {
						Tunnels[tunnel[0]] = append(Tunnels[tunnel[0]], tunnel[1])
					}
				}
				i = j
			}
		}
		i++
	}
	return AntNum, Rooms[0], Rooms[len(Rooms)-1], Tunnels, Rooms, Coord
}
