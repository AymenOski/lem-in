package functions

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"lem-in/utils"
)

func Parsing() (int, string, string, map[string][]string, []string, []utils.Coordinates) {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please create a file in the root of the project and put input to it")
		os.Exit(0)
	}

	delimiter := "\n"
	if strings.Contains(string(file), "\r\n") {
		delimiter = "\r\n"
	}

	Lines := strings.Split(string(file), delimiter)

	AntNum := 0
	c := 0
	TempEndRoom := ""
	TempStartRoom := ""
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

	before := false
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

			before = true

			// start processing after the ##start flag
			for j := i + 1; !strings.HasPrefix(Lines[j], "##end"); j++ {

				if c == 0 {
					TempStartRoom = strings.Fields(Lines[j])[0]
					c = 1
				}

				// if its a comment skip it
				if strings.HasPrefix(Lines[j], "#") || Lines[j] == "" {
					continue
				}

				// Storing staring room ,rooms and coordinates
				if len(strings.Fields(Lines[j])) != 3 {
					fmt.Println("ERROR: Invalid input format.")
					os.Exit(0)

				} else {

					// checking if the rooms are not duplicated
					if !slices.Contains(Rooms, strings.Fields(Lines[j])[0]) {
						Rooms = append(Rooms, strings.Fields(Lines[j])[0])
					} else {
						fmt.Println("ERROR: Invalid input format, Room name already exists")
						os.Exit(0)
					}

					Coord = append(Coord, utils.Coordinates{X: Atoi(strings.Fields(Lines[j])[1]), Y: Atoi(strings.Fields(Lines[j])[2])})
				}

				// its necessary to update i so we can reduce the number of unnecessary iterations
				i = j

			}
			// this case is to store data before ##start flag if there is any
		case !before:
			// checking if the rooms are not duplicated
			if !slices.Contains(Rooms, strings.Fields(val)[0]) {
				Rooms = append(Rooms, strings.Fields(val)[0])
			} else {
				fmt.Println("ERROR: Invalid input format, Room name already exists")
				os.Exit(0)
			}

			Coord = append(Coord, utils.Coordinates{X: Atoi(strings.Fields(val)[1]), Y: Atoi(strings.Fields(val)[2])})

		case val == "##end":

			for j := i + 1; j < len(Lines); j++ {

				// if its a comment or there is an empty line skip it
				if strings.HasPrefix(Lines[j], "#") || Lines[j] == "" {
					continue
				}

				if !strings.Contains(Lines[j], "-") {

					if c == 1 {
						TempEndRoom = strings.Fields(Lines[j])[0]
						c = -1
					}

					if len(strings.Fields(Lines[j])) != 3 {
						fmt.Println("ERROR: Invalid input format, somthing is wrong with the rooms")
						os.Exit(0)

					} else {
						Rooms = append(Rooms, strings.Fields(Lines[j])[0])
						Coord = append(Coord, utils.Coordinates{X: Atoi(strings.Fields(Lines[j])[1]), Y: Atoi(strings.Fields(Lines[j])[2])})
						continue
					}

				}

				if strings.Contains(Lines[j], "-") {

					// Storing end room and tunnels
					tunnel := strings.Split(Lines[j], "-")

					if len(tunnel) == 2 && tunnel[0] != "" && tunnel[1] != "" {

						TunnelsMaker(Tunnels, tunnel[0], tunnel[1])
						TunnelsMaker(Tunnels, tunnel[1], tunnel[0])

					} else {
						fmt.Println("ERROR: Invalid input format, Tunnels should be in the format of room1-room2")
						os.Exit(0)
					}

				} else {
					fmt.Println("ERROR: Invalid input format, Tunnels should be in the format of room1-room2")
					os.Exit(0)
				}

				i = j
			}
		}

		i++
	}

	return AntNum, TempStartRoom, TempEndRoom, Tunnels, Rooms, Coord
}
