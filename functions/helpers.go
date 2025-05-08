package functions

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func TunnelsMaker(Tunnels map[string][]string, a string, b string) {
	// if b is not already linked to a, create a new tunnel
	check_1 := slices.Contains(Tunnels[a], b)
	if !check_1 {
		Tunnels[a] = append(Tunnels[a], b)
	}
}

func Atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("ERROR: invalid data format, Make sure the Coordinates are numbers")
		os.Exit(0)
	}
	return val
}

func CleanStr(str string) string {
	s := ""
	for _, val := range str {
		if val == ' ' {
			continue
		}
		s += string(val)
	}
	return s
}

// func ValidCoords(Coords []utils.Coordinates) {
// 	//=========//====//====//====//====//====//====//====//====//====//====//====//====//
// 	// logic not valid --> O(n^2) : we need a map to store the coordinates instead
// 	//====//====//====//====//====//====//====//====//====//====//====//====//====//====//

// 	// for i := range Coords {
// 	// 	for j := range Coords {
// 	// 		if i != j {
// 	// 			if Coords[i].X == Coords[j].X && Coords[i].Y == Coords[j].Y {
// 	// 				fmt.Println("ERROR: invalid data format, duplicate coordinates")
// 	// 				os.Exit(0)
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }
// }

// func ValidRooms(Rooms []string, Tunnels map[string][]string) {
// 	// intoSlice := []string{}
// 	// // converting the map keys to a slice
// 	// for val := range Tunnels {
// 	// 	intoSlice = append(intoSlice, val)
// 	// }
// 	// for i := range Rooms {
// 	// 	if slices.Contains(intoSlice, Rooms[i]) {
// 	// 		fmt.Println("ERROR: invalid data format, Rooms issue detected")
// 	// 		os.Exit(0)
// 	// 	}
// 	// }
// }
