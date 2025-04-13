package utils

import (
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

type Ants struct {
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
}

type Room struct {
	Name      string
	Occupied  bool
	Neighbors []*Room
}

type Graph struct {
	Rooms map[string]*Room
}

func (g *Graph) AddRoom(roomName string) {
	if _, exists := g.Rooms[roomName]; !exists {
		g.Rooms[roomName] = &Room{
			Name:      roomName,
			Neighbors: make([]*Room, 0),
		}
	}
}

func (g *Graph) LinkRooms(tunnels map[string][]string) {
	for roomName, neighbors := range tunnels {
		for _, neighborName := range neighbors {
			if g.Rooms[roomName] != nil && g.Rooms[neighborName] != nil {
				g.Rooms[roomName].Neighbors = append(g.Rooms[roomName].Neighbors, g.Rooms[neighborName])
			}
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, room := range g.Rooms {
		fmt.Printf("Room: %s, Neighbors: %v\n", room.Name, room.Neighbors)
	}
}

func (g *Graph) BFS(start, end string, ants []*Ant) {

	queue := []string{start}
	prev := make(map[string]string)
	str := ""
	// c := 0
	allAntsAtEnd := false
	for !allAntsAtEnd {
		for _, Ant := range ants {
			for len(queue) > 0 {
				current := queue[0]
				// deleat from queue
				queue = queue[1:]

				if current == end {
					// creating a specefic path for each ant
					for at := end; at != start; at = prev[at] {
						Ant.Path = append([]string{at}, Ant.Path...)
					}
					Ant.CurrentRoom = g.Rooms[Ant.Path[0]]
					Ant.CurrentRoom.Occupied = true
					str += fmt.Sprintf("%v-%v", Ant.Id, Ant.CurrentRoom)
					Ant.Path = Ant.Path[1:]
					queue = []string{start}
					continue
				}

				for _, neighbor := range g.Rooms[current].Neighbors {
					if !neighbor.Occupied {
						prev[neighbor.Name] = current
						queue = append(queue, neighbor.Name)
					}
				}
			}

		}
		for i, Ant := range ants {
			if Ant.CurrentRoom.Name != end {
				Ant.CurrentRoom.Occupied = false
				break
			}
			if i == len(ants)-1 {
				allAntsAtEnd = true
			}
		}
	}
}
