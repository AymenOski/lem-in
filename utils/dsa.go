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
	Position     []Coordinates
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

func hasLengthString(paths [][]string) bool {
	for i := range paths {
		if len(paths[i]) == 2 {
			return true
		}
	}
	return false
}

func Filtring(allCombos [][][]string) [][]string {
	var bestCombo [][]string
	maxPaths := -1
	minRooms := int(^uint(0) >> 1)

	for _, combo := range allCombos {
		numPaths := len(combo)
		roomSet := make(map[string]bool)

		for _, path := range combo {
			for _, room := range path {
				roomSet[room] = true
			}
		}
		numRooms := len(roomSet)

		if numPaths > maxPaths || (numPaths == maxPaths && numRooms < minRooms) {
			bestCombo = combo
			maxPaths = numPaths
			minRooms = numRooms
		}
	}

	return bestCombo
}
