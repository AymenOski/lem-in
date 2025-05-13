package utils

import "fmt"

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
