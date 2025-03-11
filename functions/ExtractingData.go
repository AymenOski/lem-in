package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExtractingDate() (int, string, string, map[string][]string, []string) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please create a file in the root of the project and put input to it")
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	c := 0
	antNum := 0
	startingRoom := ""
	endingRoom := ""
	i := 0
	// FC denoted as "format checker"
	FC_StartFlag := 0
	FC_EndFlag := 0
	FC_AntNumber := -1
	Tunnels := make(map[string][]string)
	Rooms := make([]string, 0)
	check := false
	for scanner.Scan() {
		// adding the ability to make comments
		if !(scanner.Text() == "##start") && !(scanner.Text() == "##end") {
			if strings.HasPrefix(scanner.Text(), "#") || strings.HasPrefix(scanner.Text(), "L") {
				continue
			}
		}
		if c == 0 {
			antNum, err = strconv.Atoi(scanner.Text())
			if antNum <= 0 {
				fmt.Println("ERROR: invalid data format, Number of ants should be greater than 0")
				os.Exit(0)
			}
			FC_AntNumber = 0
			if err != nil {
				fmt.Println(err)
				fmt.Println("ERROR: invalid data format, Ants number should be an integer.")
				os.Exit(0)
			}
			c++
		}
		if c == 3 {
			// storing data after the ##end flag
			tunnel := strings.Split(scanner.Text(), "-")
			if len(tunnel) == 2 {
				Tunnels[tunnel[0]] = append(Tunnels[tunnel[0]], tunnel[1])
			}
		}
		// storing data after the ##start flag
		if check {
			if len(scanner.Text()) > 0 {
				if strings.HasPrefix(scanner.Text(), "#") {
					// continue and dont store this data
				} else {
					// store this data
					Rooms = append(Rooms, scanner.Text()[:1])
				}
			} else {
				fmt.Println("ERROR: invalid data format, Missing ants starting room")
				os.Exit(0)
			}
		} else if !check && c == 2 {
			if len(scanner.Text()) > 0 {
				if strings.HasPrefix(scanner.Text(), "#") {
					// continue and dont store this data
				} else {
					endingRoom = scanner.Text()[:1]
					c = 3
				}
			} else {
				fmt.Println("ERROR: invalid data format, Missing ants ending room")
				os.Exit(0)
			}
		}
		if scanner.Text() == "##start" {
			FC_StartFlag = i
			check = true
		} else if scanner.Text() == "##end" {
			FC_EndFlag = i
			check = false
			c = 2
		}
		i++
	}
	startingRoom = Rooms[0]
	Rooms = append(Rooms, endingRoom)
	FormatChekcer(FC_AntNumber, FC_StartFlag, FC_EndFlag, i)
	return antNum, startingRoom, endingRoom, Tunnels, Rooms
}
