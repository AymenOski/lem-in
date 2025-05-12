package utils

import "fmt"

func GraphConstructor() *Graph {
	return &Graph{
		Rooms: make(map[string]*Room),
		Data:  ParseInfo{Coords: make(map[[2]int]bool)},
	}
}

// func (g *Graph) LinkRooms(tunnels map[string][]string) {
// 	// bidirectional links between rooms
// 	for roomName, neighbors := range tunnels {
// 		for _, neighborName := range neighbors {
// 			if g.Rooms[roomName] != nil && g.Rooms[neighborName] != nil {
// 				g.Rooms[roomName].Neighbors = append(g.Rooms[roomName].Neighbors, g.Rooms[neighborName])
// 			}
// 		}
// 	}
// }

func (g *Graph) AddRoom(roomName string) {
	if _, exists := g.Rooms[roomName]; !exists {
		g.Rooms[roomName] = &Room{
			Name:      roomName,
			Neighbors: make([]*Room, 0),
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, room := range g.Rooms {
		fmt.Printf("Room: %s, Neighbors: %v\n", room.Name, room.Neighbors)
	}
}





