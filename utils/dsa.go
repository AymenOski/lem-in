package utils

import (
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

type Colony struct {
	AntNum       int
	Rooms        []string
	Tunnels      map[string][]string
	StartingRoom string
	EndingRoom   string
}

type Ant struct {
	Id          string
	CurrentRoom *Room
	Path        []string
	HasMoved    bool
	Step        int
}

type Room struct {
	Name      string
	Occupied  bool
	Neighbors []*Room
}

type Graph struct {
	Rooms map[string]*Room
	Paths [][]string
	Data  ParseInfo
	Col   *Colony
}

type ParseInfo struct {
	Coords               map[[2]int]bool
	Phase                int
	StartFound, EndFound bool
}

func CreateAnts(AntNum int, StartingRoom *Room) []*Ant {
	ants := []*Ant{}
	for i := 1; i <= AntNum; i++ {
		ant := &Ant{
			Id:          fmt.Sprintf("L%d", i),
			CurrentRoom: StartingRoom,
		}
		ants = append(ants, ant)
	}
	return ants
}

// Check if we have start room attache to end room.
// if so ,skip to another neighbor and go see other paths.
func hasLengthString(paths [][]string) bool {
	for i := range paths {
		if len(paths[i]) == 2 {
			return true
		}
	}
	return false
}
